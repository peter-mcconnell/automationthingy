package web

import (
	"context"
	"html/template"
	"net/http"

	"github.com/peter-mcconnell/automationthingy/config"
)

type Server struct {
	ctx       context.Context
	logger    config.Logger
	templates *template.Template
	routes    []*route
	Mux       *http.ServeMux
	Config    *config.Config
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Debugf("%s %s", r.Method, r.URL.Path)
	s.Mux.ServeHTTP(w, r)
}

func NewServer(ctx context.Context, logger config.Logger, mux *http.ServeMux) (*Server, error) {
	server := &Server{
		ctx:    ctx,
		logger: logger,
		Mux:    mux,
	}
	automationthingyConfig, err := config.LoadConfig(&logger)
	if err != nil {
		return server, err
	}
	server.Config = &automationthingyConfig
	if err := server.addRoutes(); err != nil {
		return server, err
	}
	if err := server.addTemplates(); err != nil {
		return server, err
	}

	return server, nil
}
