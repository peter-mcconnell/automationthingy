package api

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func (s *Server) apiV1ExecutorDemo(w http.ResponseWriter, r *http.Request) {
	// apiRequest := s.parseApiRequest(r.URL.Path)
	// exec := executor.DemoExecutor{}
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("failed to set flusher")
	}
	w.Header().Set("Connection", "Keep-Alive")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	// for t := range exec.Execute(apiRequest.id) {
	ticker := time.NewTicker(time.Millisecond * 500)
	limit := 20
	current := 0
	for t := range ticker.C {
		io.WriteString(w, "Chunk\n")
		fmt.Println("Tick at", t)
		flusher.Flush()
		current += 1
		if current >= limit {
			ticker.Stop()
			break
		}
	}
}
