package api

import (
	"net/http"

	"github.com/peter-mcconnell/automationthingy/auth"
)

func (s *Server) apiV1GithubLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		GithubConfig:    &s.Config.General.Api.Auth.Github,
		SecretmgrConfig: &s.Config.Secretmgr,
	}
	github.LoginHandler(w, r)
}

func (s *Server) apiV1GithubCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		GithubConfig:    &s.Config.General.Api.Auth.Github,
		SecretmgrConfig: &s.Config.Secretmgr,
	}
	github.CallbackHandler(w, r)
}
