package gcs

import (
	backend "cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

type Storage struct {
	logger *zap.SugaredLogger
	client *backend.Client
	bucket string
}

func NewStorage(logger *zap.SugaredLogger, client *backend.Client, bucket string) *Storage {
	return &Storage{
		logger: logger,
		client: client,
		bucket: bucket,
	}
}

func (s Storage) GetAllPackages(ctx context.Context) []listing.PackageReference {
	bkt := s.client.Bucket(s.bucket)
	query := &backend.Query{Prefix: ""}

	var pkgsRef []listing.PackageReference
	it := bkt.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("an error occured while retrieving files from Google Cloud Storage. got: %v", err)
		}
		pkgsRef = append(pkgsRef, listing.PackageReference(attrs.Name))
	}
	return pkgsRef
}

func (s Storage) String() string {
	return fmt.Sprintf("GoogleCloudStorage[bucket=%q]", s.bucket)
}
