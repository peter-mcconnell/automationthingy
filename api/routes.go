package api

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type route struct {
	pattern string
	handler http.Handler
}

func (s *Server) parseApiRequest(url string) ApiRequest {
	uriParts := strings.Split(url, "/")[2:]
	var (
		id  uuid.UUID
		err error
	)
	if len(uriParts) > 1 {
		id, err = uuid.Parse(uriParts[2])
		if err != nil {
			panic(err)
		}
	}
	return ApiRequest{
		apiVersion:    uriParts[0],
		resource:      uriParts[1],
		sub_resources: strings.Split(uriParts[1], "_")[1:],
		id:            id,
	}
}

func (s *Server) HandleFunc(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *Server) addRoutes() error {
	var routes = []route{
		newRoute("/api/v1/executor_demo/", s.apiV1ExecutorDemo),
		newRoute("/api/v1/executor_local/", s.apiV1ExecutorLocal),
		newRoute("/", func(w http.ResponseWriter, r *http.Request) {
			http.NotFound(w, r)
		}),
	}
	for _, route := range routes {
		s.HandleFunc(route.pattern, route.handler)
	}
	return nil
}

func newRoute(pattern string, handler http.HandlerFunc) route {
	return route{pattern, handler}
}
