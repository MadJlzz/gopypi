package registry

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/storage"
	"context"
	"github.com/MadJlzz/gopypi/configs"
	"go.uber.org/zap"
	"log"
)

type Factory interface {
	CreateRegistry() Registry
}

func GetFactory(logger *zap.SugaredLogger, configuration configs.StorageConfiguration) Factory {
	var f Factory
	switch configuration.GetType() {
	case configs.GCS:
		c, _ := configuration.(*configs.GCPConfiguration)
		f = &GCSFactory{
			logger:        logger,
			configuration: c,
		}
	case configs.S3:
		log.Fatalln("gopypi doesn't support S3 storage backend for the moment")
	}
	return f
}

type GCSFactory struct {
	logger        *zap.SugaredLogger
	configuration *configs.GCPConfiguration
}

func (f *GCSFactory) CreateRegistry() Registry {
	ctx := context.TODO()
	storageCli, err := storage.NewClient(ctx)
	if err != nil {
		f.logger.Fatalf("impossible to initialize GCS storage. got: %v", err)
	}
	secretCli, err := secretmanager.NewClient(ctx)
	if err != nil {
		f.logger.Fatalf("impossible to initialize SecretManager storage. got: %v", err)
	}
	return NewGCStorage(f.logger, storageCli, secretCli, f.configuration.GCS)
}
