package main

import (
	"context"
	"github.com/MadJlzz/gopypi/internal/http/rest"
	"github.com/MadJlzz/gopypi/internal/storage/gcs"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap l: %v", err)
	}
	defer l.Sync()

	// SugaredLogger includes both printf-style APIs.
	logger := l.Sugar()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		logger.Infof("Defaulting to port %s", port)
	}

	ctx := context.Background()

	// TODO: use a factory to retrieve the correct storage and be more flexible.
	storage := gcs.NewStorage(ctx, logger, "gopypi")
	logger.Infof("new connection with storage backend [%v]", storage)

	// set up HTTP server
	ph := rest.NewRepositoryHandler(logger, storage)
	router := ph.Router(ctx)

	logger.Info("The PyPi server is live: http://localhost:8080")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Fatal(err)
	}
}
