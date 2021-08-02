package registry

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/storage"
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
	logger        *zap.SugaredLogger
	storage       *storage.Client
	secret        *secretmanager.Client
	configuration configs.GCSConfiguration
}

func NewGCStorage(logger *zap.SugaredLogger, storageCli *storage.Client, secretCli *secretmanager.Client, configuration configs.GCSConfiguration) *GCStorage {
	return &GCStorage{
		logger:        logger,
		storage:       storageCli,
		secret:        secretCli,
		configuration: configuration,
	}
}

func (s GCStorage) GetAllProjects() []Project {
	bkt := s.storage.Bucket(s.configuration.BucketName)
	q := &storage.Query{
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

	bkt := s.storage.Bucket(s.configuration.BucketName)
	q := &storage.Query{
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
		sUrl, err := GenerateSignedURL(s.configuration.BucketName, attrs.Name, s.secret, s.configuration.Secret.Name)
		if err != nil {
			s.logger.Errorf("an error occured while signing file from GCS. got: %v", err)
		}
		pkgs = append(pkgs, Package{
			Filename: path.Base(attrs.Name),
			URI:      sUrl,
		})
	}
	return pkgs
}

func (s GCStorage) String() string {
	return fmt.Sprintf("GoogleCloudStorage[bucket=%s, signingSecret=%s]", s.configuration.BucketName, s.configuration.Secret.Name)
}

func GenerateSignedURL(bucket string, objectReference string, secretCli *secretmanager.Client, secret string) (string, error) {
	req := &secretmanagerpb.AccessSecretVersionRequest{Name: secret}
	resp, err := secretCli.AccessSecretVersion(context.TODO(), req)
	if err != nil {
		return "", fmt.Errorf("could not retrieve secret[%s] to sign URL. got: %v", secret, err)
	}
	conf, err := google.JWTConfigFromJSON(resp.Payload.GetData())
	if err != nil {
		return "", fmt.Errorf("could not load JWT configuration file from secret. got: %v", err)
	}
	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(1 * time.Minute),
	}
	sUrl, err := storage.SignedURL(bucket, objectReference, opts)
	if err != nil {
		return "", fmt.Errorf("could not sign storage URL. got: %v", err)
	}
	return sUrl, nil
}
