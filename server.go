package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	Client *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.Client = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.Client.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Client.Shutdown(ctx)
}
