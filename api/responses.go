package api

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status     int         `json:"status"`
	StatusText string      `json:"status_text"`
	Result     interface{} `json:"result"`
}

func (s *Server) err(w http.ResponseWriter, sErr error, statusCode int) {
	statusText := http.StatusText(statusCode)
	respStruct := ApiResponse{
		Status:     statusCode,
		StatusText: statusText,
		Result:     sErr.Error(),
	}
	resp, err := json.Marshal(respStruct)
	if err != nil {
		s.logger.Errorf("%s", err.Error())
		statusCode = http.StatusInternalServerError
		resp = []byte(err.Error())
	}
	s.logger.Debugf("%s", resp)
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func (s *Server) errNotFound(w http.ResponseWriter, err error) {
	s.err(w, err, http.StatusNotFound)
}

func (s *Server) errInternalError(w http.ResponseWriter, err error) {
	s.err(w, err, http.StatusNotFound)
}
