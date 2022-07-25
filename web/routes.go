package web

import (
	"net/http"

	"github.com/peter-mcconnell/automationthingy/model"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type route struct {
	pattern string
	handler http.Handler
}

func (s *Server) HandleFunc(pattern string, handler http.Handler) {
	oltphandler := otelhttp.NewHandler(
		otelhttp.WithRouteTag(
			pattern,
			handler,
		),
		pattern,
		otelhttp.WithPublicEndpoint(),
	)
	s.Mux.Handle(pattern, oltphandler)
}

func (s *Server) staticAsset(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func (s *Server) addRoutes() error {
	projects := model.Project{}
	commonViewData := commonViewData{
		Title:    "AutomationThingy",
		Projects: projects.GetAll(),
		BaseHref: s.Config.General.Web.Host,
	}
	var headers map[string]string
	var routes = []route{
		newRoute("/login/github/callback", s.githubCallback),
		newRoute("/login/github", s.githubLogin),
		newRoute("/job/", func(w http.ResponseWriter, r *http.Request) {
			s.job(w, r, commonViewData)
		}),
		newRoute("/project/", func(w http.ResponseWriter, r *http.Request) {
			s.project(w, r, commonViewData)
		}),
		newRoute("/static/", s.staticAsset),
		newRoute("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				s.page("home", w, r, headers, commonViewData)
			} else {
				w.WriteHeader(404)
				s.page("404", w, r, headers, commonViewData)
			}
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
