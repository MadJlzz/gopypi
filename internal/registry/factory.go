package registry

import (
	"github.com/MadJlzz/gopypi/configs"
	"github.com/MadJlzz/gopypi/internal/backend/gcs"
	"go.uber.org/zap"
	"log"
)

type Factory struct {
	Logger        *zap.SugaredLogger
	Configuration configs.StorageConfiguration
}

func (f *Factory) CreateRegistry() Registry {
	var r Registry
	switch f.Configuration.GetType() {
	case configs.GCS:
		r = gcs.NewStorage(f.Logger, f.Configuration)
	case configs.S3:
		log.Fatalln("gopypi doesn't support S3 storage backend for the moment")
	}
	return r
}