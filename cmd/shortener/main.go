package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"github.com/vabispklp/yap/internal/app/storage/inmem"
	"github.com/vabispklp/yap/internal/app/storage/ondisk"
	"github.com/vabispklp/yap/internal/app/storage/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vabispklp/yap/api/rest"
	"github.com/vabispklp/yap/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Init cinfig error: %s", err)
	}
	ctx := context.Background()

	storageService, err := postgres.NewStorage(cfg.GetDatabaseDSN())
	if err != nil {
		log.Print(err)
	}

	if storageService == nil {
		storageService, err = ondisk.NewStorage(cfg.GetFileStoragePath())
		if err != nil {
			log.Print(err)
		}
	}

	if storageService == nil {
		storageService = inmem.NewStorage()
		if err != nil {
			log.Print(err)
		}
	}

	defer storageService.Close()

	shortenerService, err := shortener.NewShortener(storageService, cfg.GetBaseURL())
	if err != nil {
		log.Print(err)
	}

	server, err := rest.NewServer(cfg, shortenerService)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = server.Close(ctx); err != nil {
			log.Println("Server Close error: " + err.Error())
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if err = server.Start(ctx); err != nil {
		log.Println("ыerver Close error: " + err.Error())
	}
	log.Print("Server Started")

	<-done
	log.Print("Graceful shutdown Started")
}
