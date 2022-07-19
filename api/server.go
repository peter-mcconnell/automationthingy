package api

import (
	"context"
	"html/template"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/config"
)

type Server struct {
	ctx       context.Context
	port      int
	logger    config.Logger
	templates *template.Template
	routes    []*route
	Mux       *http.ServeMux
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

func (s *Server) RunBackground() {
	sPort := ":" + strconv.Itoa(s.port)
	s.logger.Debugf("running server in background on port %s", sPort)
	http.ListenAndServe(sPort, s.Mux)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Debugf("%s %s", r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	s.Mux.ServeHTTP(w, r)
}

func NewServer(port int, logger config.Logger, mux *http.ServeMux) (*Server, error) {
	server := &Server{
		port:   port,
		ctx:    context.Background(),
		logger: logger,
		Mux:    mux,
	}
	server.OltpInitialize()
	automationthingyConfig, err := config.LoadConfig(&logger)
	if err != nil {
		return server, err
	}
	server.Config = &automationthingyConfig
	if err := server.addRoutes(); err != nil {
		return server, err
	}
	// dw
	// server.RunBackground()
	return server, nil
}
