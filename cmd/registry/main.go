package main

import (
	backend "cloud.google.com/go/storage"
	"context"
	"github.com/MadJlzz/gopypi/internal/http/rest"
	"github.com/MadJlzz/gopypi/internal/storage/gcs"
	"go.uber.org/zap"
	"log"
	"net/http"
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
	defer client.Close()

	// TODO: use a factory to retrieve the correct storage and be more flexible.
	storage := gcs.NewStorage(logger, client, "gopypi")
	logger.Infof("new connection with storage backend [%v]", storage)

	// set up HTTP server
	ph := rest.NewRepositoryHandler(logger, storage)
	router := ph.Router(ctx)

	logger.Info("The PyPi server is live: http://localhost:8080")
	logger.Fatal(http.ListenAndServe(":8080", router))
}
