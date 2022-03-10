package rest

import (
	"context"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"log"
	"net/http"
	"time"

	"github.com/vabispklp/yap/internal/config"
)

// Server implements HTTP server and keeps its dependencies.
// Contract: not thread-safe
type Server struct {
	server *http.Server

	started bool
}

// NewServer создает структуру для запуска http сервера
func NewServer(cfg config.Сonfig, shortener *shortener.Shortener) (*Server, error) {
	server := http.Server{Addr: cfg.ServerAddr}

	router, err := initRouter(shortener)
	if err != nil {
		return nil, err
	}

	server.Handler = router

	return &Server{server: &server}, nil
}

// Start запускает http сервер
func (s *Server) Start() error {
	go func() {
		log.Print(s.server.ListenAndServe())
	}()

	s.started = true

	return nil
}

// Close останавливает http сервер
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
