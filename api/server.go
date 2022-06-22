package api

import (
	"html/template"
	"net/http"

	"github.com/peter-mcconnell/automationthingy/scm"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Server struct {
	logger    Logger
	mux       *http.ServeMux
	templates *template.Template
	routes    []*route
}

type ApiRequest struct {
	apiVersion    string
	resource      string
	sub_resources []string
	id            string
	writer        http.ResponseWriter
	request       *http.Request
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.mux.ServeHTTP(w, r)
}

func NewServer(logger Logger, mux *http.ServeMux) (*Server, error) {
	server := &Server{
		logger: logger,
		mux:    mux,
	}
	go scm.CloneJobRepos()
	if err := server.addRoutes(); err != nil {
		return server, err
	}
	return server, nil
}
