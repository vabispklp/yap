package main

import (
	"context"
	"log"

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

	log.Print(server.Start(ctx))
}
