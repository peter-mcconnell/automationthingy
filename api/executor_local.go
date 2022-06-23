package api

import (
	"net/http"

	"github.com/peter-mcconnell/automationthingy/executor"
	"github.com/peter-mcconnell/automationthingy/model"
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
	scriptModel := model.Script{
		Config: *s.Config,
	}
	script, err := scriptModel.GetOne(apiRequest.id)
	if err != nil {
		// TODO: add proper error handling
		panic(err)
	}
	exectr := executor.LocalExecutor{
		ID:             apiRequest.id,
		Config:         *s.Config,
		Script:         script,
		Flusher:        flusher,
		ResponseWriter: w,
	}
	exectr.Execute()
}
