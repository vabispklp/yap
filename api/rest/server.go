package rest

import (
	"context"
	"log"
	"net/http"
	"time"
)

// Server implements HTTP server and keeps its dependencies.
// Contract: not thread-safe
type Server struct {
	server *http.Server

	started bool
}

func NewServer() (*Server, error) {
	server := http.Server{Addr: "localhost:8080"}
	router, err := initRouter()
	if err != nil {
		return nil, err
	}

	server.Handler = router

	return &Server{server: &server}, nil
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		log.Print(s.server.ListenAndServe())
	}()

	s.started = true

	return nil
}

func (s *Server) Close(ctx context.Context) error {
	if !s.started {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.started = false
		return err
	}

	s.started = false

	return nil
}
