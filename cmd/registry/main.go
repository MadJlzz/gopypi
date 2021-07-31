package main

import (
	backend "cloud.google.com/go/storage"
	"context"
	"github.com/MadJlzz/gopypi/internal/storage/gcs"
	"go.uber.org/zap"
	"log"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap l: %v", err)
	}
	defer l.Sync()

	// SugaredLogger includes both printf-style APIs.
	logger := l.Sugar()

	ctx := context.Background()
	client, err := backend.NewClient(ctx)
	if err != nil {
		logger.Fatalf("impossible to initialize GCS client. got: %v", err)
	}

	storage := gcs.NewStorage(logger, client, "gopypi")
	logger.Info(storage)

	pkgsRef := storage.GetAllPackages(ctx)
	logger.Info(pkgsRef)
}
