package main

import (
	"github.com/MadJlzz/gopypi/internal/storage/disk"
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

	storage := disk.NewStorage(logger, disk.WithPath("K:\\test-zone\\gopypi"))
	logger.Info(storage)

	dist := storage.GetAllPackages()
	logger.Info(dist)
}
