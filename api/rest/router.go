package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/vabispklp/yap/api/rest/handlers"
	"github.com/vabispklp/yap/api/rest/middleware"
	"github.com/vabispklp/yap/internal/app/service/shortener"
)

func initRouter(shortener *shortener.Shortener) (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Use(middleware.GzipHandle, middleware.AuthHandle)

	h, err := handlers.NewHandler(shortener)
	if err != nil {
		return nil, err
	}

	router.Get("/{id}", h.GetHandleGetURL())
	router.Post("/", h.GetHandlerAddURL())
	router.Post("/api/shorten", h.GetHandlerAddShorten())
	router.Get("/user/urls", h.GetHandleGetUserURLs())
	router.Get("/ping", h.GetHandlerPing())

	return router, nil
}
