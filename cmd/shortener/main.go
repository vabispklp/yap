package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/vabispklp/yap/api/rest"
	"github.com/vabispklp/yap/internal/app/service/shortener"
	"github.com/vabispklp/yap/internal/app/storage"
	"github.com/vabispklp/yap/internal/app/storage/inmem"
	"github.com/vabispklp/yap/internal/app/storage/ondisk"
	"github.com/vabispklp/yap/internal/app/storage/postgres"
	"github.com/vabispklp/yap/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Init cinfig error: %s", err)
	}
	ctx := context.Background()

	storageService, err := buildStorage(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer storageService.Close()

	shortenerService, err := shortener.NewShortener(storageService, cfg.BaseURL)
	if err != nil {
		log.Print(err)
	}

	server, err := rest.NewServer(*cfg, shortenerService)
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

	if err = server.Start(); err != nil {
		log.Fatalf("Server Close error: " + err.Error())
	}
	log.Print("Server Started")

	<-done
	log.Print("Graceful shutdown Started")
}

func buildStorage(cfg config.Сonfig) (storage.StorageExpected, error) {
	var (
		storageService storage.StorageExpected
		err            error
	)
	if cfg.DatabaseDSN != "" {
		storageService, err = postgres.NewStorage(cfg.DatabaseDSN)
	} else if cfg.FileStoragePath != "" {
		storageService, err = ondisk.NewStorage(cfg.FileStoragePath)
	} else {
		storageService = inmem.NewStorage()
	}

	return storageService, err
}
