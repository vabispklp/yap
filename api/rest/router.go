package rest

import (
	"github.com/go-chi/chi/v5"

	"github.com/vabispklp/yap/api/rest/handlers"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"github.com/vabispklp/yap/internal/app/storage/ondisk"
	"github.com/vabispklp/yap/internal/config"
)

func initRouter(cfg config.ConfigExpected) (*chi.Mux, error) {
	router := chi.NewRouter()

	storageOnDisk, err := ondisk.New(cfg.GetFileStoragePath())
	if err != nil {
		return nil, err
	}

	shortenerService, err := shortener.NewShortener(storageOnDisk, cfg.GetBaseURL())
	if err != nil {
		return nil, err
	}

	h, err := handlers.NewHandler(shortenerService)
	if err != nil {
		return nil, err
	}

	router.Get("/{id}", h.GetHandleGetURL())
	router.Post("/", h.GetHandlerAddURL())
	router.Post("/api/shorten", h.GetHandlerAddShorten())

	return router, nil
}
