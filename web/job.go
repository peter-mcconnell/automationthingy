package web

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/model"
)

func (s *Server) job(w http.ResponseWriter, r *http.Request, data commonViewData) {
	job := model.Job{}
	job_id := uuid.New()
	data.View = struct{ Job model.JobData }{Job: job.GetOne(job_id)}
	err := s.templates.ExecuteTemplate(w, "job", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
