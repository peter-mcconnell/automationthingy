package executor

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type KubernetesjobExecutor struct {
	ID             uuid.UUID
	Config         config.Config
	Script         config.Script
	ResponseWriter http.ResponseWriter
	Logger         config.Logger
	clientset      *kubernetes.Clientset
}

func (k *KubernetesjobExecutor) Execute() error {
	flusher, ok := k.ResponseWriter.(http.Flusher)
	if !ok {
		return errors.New("failed to set flusher")
	}
	configFile, err := os.ReadFile("/home/pete/.kube/config.yaml")
	if err != nil {
		return err
	}
	config, err := clientcmd.RESTConfigFromKubeConfig(configFile)
	namespace := "default"
	if k.Script.Kubernetesjob.Namespace != "" {
		namespace = k.Script.Kubernetesjob.Namespace
	}
	k.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	ctx := context.TODO()
	name := fmt.Sprintf("%s", k.Script.ID)
	jobs := k.clientset.BatchV1().Jobs(namespace)
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    name,
							Image:   k.Script.Kubernetesjob.Image,
							Command: k.Script.Command,
						},
					},
					RestartPolicy: v1.RestartPolicy(k.Script.Kubernetesjob.Restartpolicy),
				},
			},
			BackoffLimit:            &k.Script.Kubernetesjob.Backofflimit,
			TTLSecondsAfterFinished: &k.Script.Kubernetesjob.Ttlsecondsafterfinished,
		},
	}

	job, err := jobs.Create(ctx, jobSpec, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	pods := k.clientset.CoreV1().Pods(namespace)
	jobPodsMap := map[string]bool{}
	for job.Status.Succeeded == 0 && job.Status.Failed == 0 {
		job, err = jobs.Get(ctx, job.GetName(), metav1.GetOptions{})
		if err != nil {
			return err
		}
		if job.Status.Active == 0 {
			k.Logger.Debug("waiting for job to start")
			time.Sleep(2 * time.Second)
			continue
		}
		jobPods, err := pods.List(ctx, metav1.ListOptions{
			LabelSelector: fmt.Sprintf("job-name=%s", job.GetName()),
		})
		if err != nil {
			return err
		}
		for _, jp := range jobPods.Items {
			if _, ok := jobPodsMap[jp.GetName()]; !ok {
				jobPodsMap[jp.GetName()] = true
				go k.getpodlogs(jp.GetName(), namespace)
			}
		}
		flusher.Flush()
		time.Sleep(1 * time.Second)
	}
	if job.Status.Failed != 0 {
		k.ResponseWriter.Write([]byte("job failed\n"))
		k.Logger.Warn("job had failures")
	}
	return err
}

func (k *KubernetesjobExecutor) getpodlogs(podname, namespace string) {
	k.ResponseWriter.Write([]byte(fmt.Sprintf("getting logs for pod %s\n", podname)))
	pods := k.clientset.CoreV1().Pods(namespace)
	for {
		time.Sleep(1 * time.Second)
		k.ResponseWriter.Write([]byte(fmt.Sprintf("waiting on logs from pod %s\n", podname)))
		podLogs := pods.GetLogs(podname, &v1.PodLogOptions{
			Follow: true,
		})
		podLogsStream, err := podLogs.Stream(context.TODO())
		if err != nil {
			k.ResponseWriter.Write([]byte(err.Error() + "\n"))
			continue
		}
		defer podLogsStream.Close()
		for {
			time.Sleep(1 * time.Second)
			k.ResponseWriter.Write([]byte("\n"))
			buf := make([]byte, 100)
			numBytes, err := podLogsStream.Read(buf)
			if numBytes == 0 {
				continue
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				k.ResponseWriter.Write([]byte(err.Error() + "\n"))
			}
			k.ResponseWriter.Write(buf[:numBytes])
		}
		break
	}
}
