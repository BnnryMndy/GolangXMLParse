package servers

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpserver *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpserver = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    100 + time.Second,
		WriteTimeout:   100 + time.Second,
	}

	return s.httpserver.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpserver.Shutdown(ctx)
}
