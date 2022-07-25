package web

import (
	"net/http"

	"github.com/peter-mcconnell/automationthingy/auth"
)

func (s *Server) githubLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		GithubConfig:    &s.Config.General.Web.Auth.Github,
		SecretmgrConfig: &s.Config.Secretmgr,
	}
	github.LoginHandler(w, r)
}

func (s *Server) githubCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		GithubConfig:    &s.Config.General.Web.Auth.Github,
		SecretmgrConfig: &s.Config.Secretmgr,
	}
	github.CallbackHandler(w, r)
}
