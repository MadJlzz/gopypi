package gcs

import (
	backend "cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"path"
	"strings"
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

func (s Storage) GetAllProjects(ctx context.Context) []listing.Project {
	bkt := s.client.Bucket(s.bucket)
	q := &backend.Query{
		Prefix:    "",
		Delimiter: "/",
	}
	err := q.SetAttrSelection([]string{"Name"})
	if err != nil {
		s.logger.Errorf("query attr selection is invalid. got: %v", err)
	}

	it := bkt.Objects(ctx, q)
	var projects []listing.Project
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("an error occured while retrieving files from Google Cloud Storage. got: %v", err)
		}
		project := strings.Trim(attrs.Prefix, "/")
		projects = append(projects, listing.Project(project))
	}
	return projects
}

func (s Storage) GetAllProjectPackages(ctx context.Context, project string) []listing.Package {
	bkt := s.client.Bucket(s.bucket)
	q := &backend.Query{
		Prefix: project,
	}
	err := q.SetAttrSelection([]string{"Name", "MediaLink"})
	if err != nil {
		s.logger.Errorf("query attr selection is invalid. got: %v", err)
	}

	it := bkt.Objects(ctx, q)
	var pkgs []listing.Package
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("an error occured while retrieving files from Google Cloud Storage. got: %v", err)
		}
		pkgs = append(pkgs, listing.Package{
			Filename: path.Base(attrs.Name),
			URI:      attrs.MediaLink,
		})
	}
	return pkgs
}

func (s Storage) String() string {
	return fmt.Sprintf("GoogleCloudStorage[bucket=%q]", s.bucket)
}
