/*This executor will execute the script local to the API invoking it*/
package executor

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"os/exec"

	"github.com/google/uuid"

	"github.com/peter-mcconnell/automationthingy/config"
)

type LocalExecutor struct {
	ID             uuid.UUID
	Config         config.Config
	Script         config.Script
	ResponseWriter http.ResponseWriter
	Logger         config.Logger
}

func (e *LocalExecutor) Execute() error {
	e.Logger.Infof("running script: %s [%s] in %s", e.Script.Name, e.Script.ID, e.Script.Workdir)
	flusher, ok := e.ResponseWriter.(http.Flusher)
	if !ok {
		return errors.New("failed to set flusher")
	}
	e.Logger.Debugf("running command: %s", e.Script.Command)
	args := []string{}
	if len(e.Script.Command) > 1 {
		args = e.Script.Command[1:]
	}
	cmd := exec.Command(e.Script.Command[0], args...)
	cmd.Dir = e.Script.Workdir
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// TODO: distinguish stderr as errors
	cmd.Stderr = cmd.Stdout

	scanner := bufio.NewScanner(out)
	if err = cmd.Start(); err != nil {
		return err
	}
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		e.Logger.Debugf("text: %s", scanner.Text())
		io.WriteString(e.ResponseWriter, scanner.Text()+"\n")
		flusher.Flush()
	}
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
