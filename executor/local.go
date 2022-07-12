/*This executor will execute the script local to the API invoking it*/
package executor

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"

	"github.com/google/uuid"

	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/types"
)

type LocalExecutor struct {
	ID             uuid.UUID
	Config         config.Config
	Script         config.Script
	ResponseWriter http.ResponseWriter
	Logger         types.Logger
}

func (e *LocalExecutor) Execute() error {
	e.Logger.Infof("running script: %s [%s]", e.Script.Name, e.Script.ID)
	flusher, ok := e.ResponseWriter.(http.Flusher)
	if !ok {
		return errors.New("failed to set flusher")
	}
	dir := fmt.Sprintf("scripts/%s", e.Script.ID)
	e.Logger.Debugf("have command: %s", e.Script.Command)
	args := []string{}
	if len(e.Script.Command) > 1 {
		args = e.Script.Command[1:]
	}
	cmd := exec.Command(e.Script.Command[0], args...)
	cmd.Dir = fmt.Sprintf("%s/%s", dir, e.Script.Workdir)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// TODO: distinguish stderr as errors
	cmd.Stderr = cmd.Stdout
	cmd.Start()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		io.WriteString(e.ResponseWriter, scanner.Text()+"\n")
		flusher.Flush()
	}
	return nil
}
