package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/vabispklp/yap/api/rest"
	"github.com/vabispklp/yap/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Init cinfig error: %s", err)
	}
	ctx := context.Background()

	db, err := mustDB(ctx, cfg.GetDatabaseDSN())
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	server, err := rest.NewServer(cfg, db)
	if err != nil {
		log.Fatal(err)
	}

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

func mustDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
