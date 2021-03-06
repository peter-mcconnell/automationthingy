package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
	"github.com/peter-mcconnell/automationthingy/model"
)

func (s *Server) job(w http.ResponseWriter, r *http.Request, data commonViewData) {
	scriptModel := model.Script{}
	script_id := uuid.New()
	script, err := scriptModel.GetOne(script_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data.View = struct{ Script config.Script }{Script: script}
	err = s.templates.ExecuteTemplate(w, "job", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
