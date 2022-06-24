package api

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
)

type Logger interface {
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

type Server struct {
	logger    Logger
	mux       *http.ServeMux
	templates *template.Template
	routes    []*route
	Config    *config.Config
}

type ApiRequest struct {
	apiVersion    string
	resource      string
	sub_resources []string
	id            uuid.UUID
	writer        http.ResponseWriter
	request       *http.Request
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Debugf("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.mux.ServeHTTP(w, r)
}

func NewServer(logger Logger, mux *http.ServeMux) (*Server, error) {
	server := &Server{
		logger: logger,
		mux:    mux,
	}
	cfg := config.Config{}
	automationthingyConfig, err := cfg.LoadConfig()
	if err != nil {
		return server, err
	}
	server.Config = &automationthingyConfig
	if err := server.addRoutes(); err != nil {
		return server, err
	}
	return server, nil
}
