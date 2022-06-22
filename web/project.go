package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/model"
)

func (s *Server) project(w http.ResponseWriter, r *http.Request, data commonViewData) {
	project := model.Project{}
	project_id := uuid.New()
	data.View = struct{ Project model.ProjectData }{Project: project.GetOne(project_id)}
	err := s.templates.ExecuteTemplate(w, "project", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
