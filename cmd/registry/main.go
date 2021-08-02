package main

import (
	"github.com/MadJlzz/gopypi/configs"
	"github.com/MadJlzz/gopypi/internal/http/rest"
	"github.com/MadJlzz/gopypi/internal/registry"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
)

// initializeLogger generates a new SugaredLogger that includes both printf-style APIs.
func initializeLogger() *zap.SugaredLogger {
	l, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap l: %v", err)
	}
	return l.Sugar()
}

func configurationFromEnv(storageType configs.StorageType) configs.StorageConfiguration {
	var c configs.StorageConfiguration
	switch storageType {
	case configs.GCS:
		c = &configs.GCPConfiguration{}
	case configs.S3:
		log.Fatalln("gopypi doesn't support S3 storage backend for the moment")
	default:
		log.Fatalf("could not load configuration for storage type: '%s'", storageType)
	}
	return c
}

func main() {
	logger := initializeLogger()
	defer logger.Sync()

	configuration := configurationFromEnv(configs.StorageType(os.Getenv("STORAGE_BACKEND")))
	configuration.LoadConfiguration()

	factory := registry.Factory{Logger: logger, Configuration: configuration}
	storage := factory.CreateRegistry()
	logger.Infof("new connection with storage backend [%v]", storage)

	// set up HTTP server...
	ph := rest.NewRepositoryHandler(logger, storage)
	router := ph.Router()

	logger.Info("The PyPi server is live: http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Fatal(err)
	}
}
