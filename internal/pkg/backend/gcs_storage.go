package backend

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/MadJlzz/gopypi/internal/pkg/model"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"path/filepath"
	"time"
)

type GoogleCloudStorage struct {
	bucket           string
	client           *storage.Client
	signedUrlOptions *storage.SignedURLOptions
}

func NewGoogleCloudStorage(bucket, credentials string) *GoogleCloudStorage {
	c, err := storage.NewClient(context.Background(), option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatalf("could not create google cloud storage client.\ngot: [%v]", err)
	}
	jsonConfig, err := ioutil.ReadFile(credentials)
	if err != nil {
		log.Fatalf("could not read service account file [%s].\ngot: [%v]", credentials, err)
	}
	conf, err := google.JWTConfigFromJSON(jsonConfig)
	if err != nil {
		log.Fatalf("could not parse service account file [%s].\ngot: [%v]", credentials, err)
	}
	return &GoogleCloudStorage{
		bucket: bucket,
		client: c,
		signedUrlOptions: &storage.SignedURLOptions{
			Scheme:         storage.SigningSchemeV4,
			Method:         "GET",
			GoogleAccessID: conf.Email,
			PrivateKey:     conf.PrivateKey,
			Expires:        time.Now().Add(1 * time.Minute),
		},
	}
}

func (gcs *GoogleCloudStorage) Load() map[string]*model.Package {
	pkgs := make(map[string]*model.Package)

	bh := gcs.client.Bucket(gcs.bucket)
	it := bh.Objects(context.TODO(), &storage.Query{Prefix: ""})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Warnf("could not read objects from bucket [%s].\ngot: [%v]", gcs.bucket, err)
		}

		// Generate the signed URL for authorizing download by whoever has the link.
		u, err := storage.SignedURL(gcs.bucket, attrs.Name, gcs.signedUrlOptions)
		if err != nil {
			log.Warnf("impossible to generate signed url for object [%s]. Skipping...\ngot:[%v]", attrs.Name, err)
		}

		pf := model.NewPackageFile(filepath.Base(attrs.Name), u)

		key := filepath.Dir(attrs.Name)
		if _, found := pkgs[key]; found {
			pkgs[key].AppendPackageFile(pf)
		} else {
			pkgs[key] = model.NewPackage(key, pf)
		}
	}
	return pkgs
}

func (gcs *GoogleCloudStorage) Close() error {
	return gcs.client.Close()
}
