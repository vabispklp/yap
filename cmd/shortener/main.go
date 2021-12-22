package main

import (
	"github.com/vabispklp/yap/internal/app/handler"
	"log"
	"net/http"
)

func main() {
	server := http.Server{Addr: "localhost:8080"}
	mux := http.NewServeMux()

	mux.Handle("/", handler.New())

	server.Handler = mux

	log.Fatal(server.ListenAndServe())
}
