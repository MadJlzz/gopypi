package registry

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	backend "cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/configs"
	"go.uber.org/zap"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"path"
	"strings"
	"time"
)

type GCStorage struct {
	logger *zap.SugaredLogger
	client *backend.Client
	secret *secretmanager.Client
	bucket string
}

func NewGCStorage(logger *zap.SugaredLogger, configuration configs.StorageConfiguration) *GCStorage {
	ctx := context.TODO()
	t, _ := configuration.(*configs.GCPConfiguration)

	client, err := backend.NewClient(ctx)
	if err != nil {
		logger.Fatalf("impossible to initialize GCS client. got: %v", err)
	}
	secret, err := secretmanager.NewClient(ctx)
	if err != nil {
		logger.Fatalf("impossible to initialize SecretManager client. got: %v", err)
	}
	return &GCStorage{
		logger: logger,
		client: client,
		secret: secret,
		bucket: t.GCS.BucketName,
	}
}

func (s GCStorage) GetAllProjects() []Project {
	bkt := s.client.Bucket(s.bucket)
	q := &backend.Query{
		Prefix:    "",
		Delimiter: "/",
	}
	err := q.SetAttrSelection([]string{"Name"})
	if err != nil {
		s.logger.Errorf("query attr selection is invalid. got: %v", err)
	}

	it := bkt.Objects(context.TODO(), q)
	var projects []Project
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("an error occured while retrieving files from Google Cloud GCStorage. got: %v", err)
		}
		project := strings.Trim(attrs.Prefix, "/")
		projects = append(projects, Project(project))
	}
	return projects
}

func (s GCStorage) GetAllProjectPackages(project string) []Package {
	ctx := context.TODO()

	bkt := s.client.Bucket(s.bucket)
	q := &backend.Query{
		Prefix: project,
	}
	err := q.SetAttrSelection([]string{"Name"})
	if err != nil {
		s.logger.Errorf("query attr selection is invalid. got: %v", err)
	}

	it := bkt.Objects(ctx, q)
	var pkgs []Package
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			s.logger.Errorf("an error occured while retrieving files from Google Cloud GCStorage. got: %v", err)
		}
		if attrs.Name == project+"/" {
			continue
		}
		pkgs = append(pkgs, Package{
			Filename: path.Base(attrs.Name),
			URI:      s.generateSignedURL(ctx, attrs.Name),
			//URI:      fmt.Sprintf("https://storage.cloud.google.com/%s/%s", s.bucket, attrs.Name),
		})
	}
	return pkgs
}

func (s GCStorage) String() string {
	return fmt.Sprintf("GoogleCloudStorage[bucket=%q]", s.bucket)
}

func (s GCStorage) generateSignedURL(ctx context.Context, name string) string {
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: "projects/561924096032/secrets/gopypi-sa-private-key/versions/latest",
	}
	resp, err := s.secret.AccessSecretVersion(ctx, req)
	if err != nil {
		s.logger.Errorf("could not retrieve secret to sign URL. got: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(resp.Payload.GetData())
	if err != nil {
		s.logger.Errorf("could not prepare JWT config file. got: %v", err)
	}

	opts := &backend.SignedURLOptions{
		Scheme:         backend.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(15 * time.Minute),
	}
	u, err := backend.SignedURL(s.bucket, name, opts)
	if err != nil {
		s.logger.Errorf("could not sign URL. got: %v", err)
	}
	return u
}
