package main

import (
	"log"
	"net/http"

	"github.com/vabispklp/yap/internal/app/handler"
	"github.com/vabispklp/yap/internal/app/storage"
)

func main() {
	server := http.Server{Addr: "localhost:8080"}
	mux := http.NewServeMux()

	mux.Handle("/", handler.New(storage.New()))

	server.Handler = mux

	log.Fatal(server.ListenAndServe())
}
