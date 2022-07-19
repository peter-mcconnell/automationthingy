package api

import (
	"net/http"

	"github.com/peter-mcconnell/automationthingy/auth"
)

func (s *Server) apiV1GithubLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		ClientID:    "",
		RedirectUri: "",
	}
	github.LoginHandler(w, r)
}

func (s *Server) apiV1GithubCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	github := auth.Github{
		ClientID:    "",
		RedirectUri: "",
	}
	github.CallbackHandler(w, r)
}
