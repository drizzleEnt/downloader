package main

import (
	"context"
	"downloader/internal/app"
	"log"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
