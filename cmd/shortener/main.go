package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vabispklp/yap/api/rest"
	"github.com/vabispklp/yap/internal/config"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Init cinfig error: %s", err)
	}

	server, err := rest.NewServer(cfg)
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
