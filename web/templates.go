package web

import (
	"html/template"
	"net/http"

	"github.com/peter-mcconnell/automationthingy/model"
)

type commonViewData struct {
	Title    string
	Projects []model.ProjectData
	View     any
	BaseHref string
}

func (s *Server) addTemplates() error {
	allPaths := []string{
		"./web/templates/home.tmpl",
		"./web/templates/header.tmpl",
		"./web/templates/footer.tmpl",
		"./web/templates/navbar.tmpl",
		"./web/templates/sidebar.tmpl",
		"./web/templates/default.tmpl",
		"./web/templates/project.tmpl",
		"./web/templates/job.tmpl",
		"./web/templates/404.tmpl",
	}
	templates, err := template.ParseFiles(allPaths...)
	if err != nil {
		return err
	}
	s.templates = templates
	return nil
}

func (s *Server) page(template string, w http.ResponseWriter, r *http.Request, headers map[string]string, data any) {
	err := s.templates.ExecuteTemplate(w, template, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
