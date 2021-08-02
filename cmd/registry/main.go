package main

import (
	"github.com/MadJlzz/gopypi/configs"
	"github.com/MadJlzz/gopypi/internal/http/rest"
	"github.com/MadJlzz/gopypi/internal/registry"
	"github.com/MadJlzz/gopypi/internal/view"
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

// configurationFromEnv retrieve the correct configuration given the storage type to use.
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
	rg := factory.CreateRegistry()
	logger.Infof("new connection with storage backend [%v]", rg)

	// set up HTTP server...
	handler := rest.Handler(logger, view.NewSimpleRepositoryTemplate(), rg)

	logger.Info("The PyPi server is live: http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.Fatal(err)
	}
}
