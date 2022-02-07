package rest

import (
	"context"
	"database/sql"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"github.com/vabispklp/yap/internal/app/storage/ondisk"
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

func NewServer(cfg config.ConfigExpected, db *sql.DB) (*Server, error) {
	server := http.Server{Addr: cfg.GetServerAddr()}

	storageOnDisk, err := ondisk.New(cfg.GetFileStoragePath())
	if err != nil {
		return nil, err
	}

	shortenerService, err := shortener.NewShortener(storageOnDisk, db, cfg.GetBaseURL())
	if err != nil {
		return nil, err
	}

	router, err := initRouter(shortenerService)
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
