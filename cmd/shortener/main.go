package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vabispklp/yap/api/rest"
)

func main() {
	server, err := rest.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	defer func() {
		if err = server.Close(ctx); err != nil {
			log.Fatalf("Server Close error: %s", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err = server.Start(ctx); err != nil {
		log.Fatalf("Server Close error: %s", err)
	}
	log.Print("Server Started")

	<-done
	log.Print("Graceful shutdown Started")
}
