package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/vabispklp/yap/internal/app/handler"
	"github.com/vabispklp/yap/internal/app/storage"
)

func main() {
	server := http.Server{Addr: "localhost:8080"}

	r := chi.NewRouter()

	h := handler.New(storage.New())

	r.Get("/{path}", h.ServeHTTP)
	r.Post("/", h.ServeHTTP)

	server.Handler = r

	log.Fatal(server.ListenAndServe())
}
