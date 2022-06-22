package web

import (
	"html/template"
	"net/http"
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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Printf("%s %s", r.Method, r.URL.Path)
	s.mux.ServeHTTP(w, r)
}

func NewServer(logger Logger, mux *http.ServeMux) (*Server, error) {
	server := &Server{
		logger: logger,
		mux:    mux,
	}
	if err := server.addRoutes(); err != nil {
		return server, err
	}
	if err := server.addTemplates(); err != nil {
		return server, err
	}

	return server, nil
}
