package rest

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/vabispklp/yap/api/rest/handlers"
	"github.com/vabispklp/yap/api/rest/middleware"
	"github.com/vabispklp/yap/internal/app/service/shortener"
)

func initRouter(shortener *shortener.Shortener) (*chi.Mux, error) {
	router := chi.NewRouter()
	router.Use(middleware.GzipHandle, middleware.AuthHandle)
	router.Mount("/debug", chiMiddleware.Profiler())

	h, err := handlers.NewHandler(shortener)
	if err != nil {
		return nil, err
	}

	router.Get("/{id}", h.GetHandleGetURL())
	router.Post("/", h.GetHandlerAddURL())
	router.Post("/api/shorten", h.GetHandlerAddShorten())
	router.Get("/api/user/urls", h.GetHandleGetUserURLs())
	router.Get("/ping", h.GetHandlerPing())
	router.Post("/api/shorten/batch", h.GetHandlerAddBatch())
	router.Delete("/api/user/urls", h.GetHandlerDelete())

	return router, nil
}
