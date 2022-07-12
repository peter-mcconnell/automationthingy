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
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	apiRequest := s.parseApiRequest(r.URL.Path)
	scriptModel := model.Script{
		Config: *s.Config,
	}
	script, err := scriptModel.GetOne(apiRequest.id)
	if err != nil {
		s.errNotFound(w, err)
		return
	}
	exectr := executor.LocalExecutor{
		Logger:         s.logger,
		ID:             apiRequest.id,
		Config:         *s.Config,
		Script:         script,
		ResponseWriter: w,
	}
	if err = exectr.Execute(); err != nil {
		s.logger.Error(err.Error())
	}
}
