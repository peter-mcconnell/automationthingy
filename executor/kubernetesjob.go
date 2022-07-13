package executor

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
}

func (k *KubernetesjobExecutor) Execute() error {
	configFile, err := os.ReadFile("/home/pete/.kube/config.yaml")
	if err != nil {
		return err
	}
	config, err := clientcmd.RESTConfigFromKubeConfig(configFile)
	namespace := "default"
	if k.Script.Kubernetesjob.Namespace != "" {
		namespace = k.Script.Kubernetesjob.Namespace
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	jobs := clientset.BatchV1().Jobs(namespace)
	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s", k.Script.ID),
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    fmt.Sprintf("%s", k.Script.ID),
							Image:   k.Script.Kubernetesjob.Image,
							Command: k.Script.Command,
						},
					},
					RestartPolicy: v1.RestartPolicy(k.Script.Kubernetesjob.Restartpolicy),
				},
			},
			BackoffLimit: &k.Script.Kubernetesjob.Backofflimit,
		},
	}

	_, err = jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	return err
}
