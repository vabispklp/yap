package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Print(s.server.ListenAndServe())
	}()

	s.started = true
	log.Print("Server Started")

	<-done
	log.Print("Graceful shutdown Started")

	if err := s.Close(ctx); err != nil {
		log.Fatalf("Server Close error: %s", err)
		return err
	}

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
