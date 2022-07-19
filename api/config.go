package api

import (
	"net/http"
)

func (s *Server) apiV1Config(w http.ResponseWriter, r *http.Request) {
	cfgJ, err := s.Config.GetConfigAsJson()
	if err != nil {
		s.errInternalError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(cfgJ))
}
