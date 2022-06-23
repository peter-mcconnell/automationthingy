/*This executor will execute the script local to the API invoking it*/
package executor

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	"github.com/google/uuid"

	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/scm"
)

type LocalExecutor struct {
	ID             uuid.UUID
	Config         config.Config
	Script         config.ScriptData
	Flusher        http.Flusher
	ResponseWriter http.ResponseWriter
}

func (e *LocalExecutor) Execute() {
	fmt.Printf("running %s", e.Script.Name)
	// TODO: abstract to support other sources (e.g. local disk)
	if e.Script.Source.Git.Repo != "" {
		if err := scm.CloneScriptRepos([]config.ScriptData{e.Script}); err != nil {
			panic(err)
		}
	}
	dir := fmt.Sprintf("scripts/%s", e.Script.ID)
	path := fmt.Sprintf("%s/.automationthingy.yaml", dir)
	sourceConfig, err := e.Config.LoadSourceConfig(path)
	if err != nil {
		// TODO: improve error handling
		panic(err)
	}
	var targetScript config.SourceScriptData
	for _, script := range sourceConfig.Scripts {
		if script.ID == e.ID {
			targetScript = script
			break
		}
	}
	if targetScript.Name == "" {
		// TODO: improve error handling
		panic("script not found")
	}
	fmt.Printf("running %s\n", targetScript.Command)
	args := strings.Fields(strings.TrimSpace(targetScript.Command))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = fmt.Sprintf("%s/%s", dir, targetScript.Workdir)
	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	// TODO: distinguish stderr as errors
	cmd.Stderr = cmd.Stdout
	cmd.Start()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		io.WriteString(e.ResponseWriter, scanner.Text()+"\n")
		e.Flusher.Flush()
	}
}
