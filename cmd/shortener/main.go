package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := http.Server{Addr: "localhost:8080"}
	router, err := initRouter()
	if err != nil {
		log.Fatalf("create router error: %s", err)
	}

	server.Handler = router

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Print(server.ListenAndServe())
	}()
	log.Print("Server Started")

	<-done
	log.Print("Graceful shutdown Started")
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Print("Graceful shutdown Finished")
}
