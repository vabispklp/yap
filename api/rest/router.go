package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/vabispklp/yap/api/rest/handlers"

	"github.com/vabispklp/yap/internal/app/service/shortener"
	"github.com/vabispklp/yap/internal/app/storage/inmem"
)

func initRouter() (*chi.Mux, error) {
	router := chi.NewRouter()

	storage := inmem.New()
	shortenerService, err := shortener.NewShortener(storage)
	if err != nil {
		return nil, err
	}

	h, err := handlers.NewHandler(shortenerService)
	if err != nil {
		return nil, err
	}

	router.Get("/{id}", h.GetHandleGetURL())
	router.Post("/", h.GetHandlerAddURL())

	return router, nil
}
