package api

import (
	"bufio"
	"io"
	"net/http"
	"os/exec"

	executor "github.com/peter-mcconnell/automationthingy/executor"
)

type CommandRequest struct {
	Command string
	Args    string
}

func (s *Server) apiV1ExecutorLocal(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("failed to set flusher")
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	apiRequest := s.parseApiRequest(r.URL.Path)
	exectr := executor.LocalExecutor{
		ID: apiRequest.id,
	}
	exectr.Execute()
	cmd := exec.Command("ping", "-c4", "8.8.8.8")
	out, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		io.WriteString(w, scanner.Text()+"\n")
		flusher.Flush()
	}
}
