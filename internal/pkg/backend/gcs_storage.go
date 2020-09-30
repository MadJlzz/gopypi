package backend

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type googleCloudStorage struct {
	localDir string
	bucket   string
	client   *storage.Client
}

func NewGoogleCloudStorage(localDir, bucket string, opts ...option.ClientOption) *googleCloudStorage {
	c, err := storage.NewClient(context.Background(), opts...)
	if err != nil {
		logrus.Fatalf("could not create google cloud storage client.\ngot: [%v]", err)
	}
	return &googleCloudStorage{
		localDir: localDir,
		bucket: bucket,
		client: c,
	}
}

func (gcs *googleCloudStorage) Open(name string) (http.File, error) {
	bh := gcs.client.Bucket(gcs.bucket)
	q := &storage.Query{Prefix: ""}

	var names []string
	it := bh.Objects(context.TODO(), q)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}
	f := &os.File{}
	for _, name := range names {
		rc, err := bh.Object(name).NewReader(context.TODO())
		if err != nil {
			logrus.Warnln(err)
		}

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			logrus.Warnln(err)
		}
		logrus.Println(data)
		_ = rc.Close()
	}
	//resp, _ := http.Get("https://google.com")
	//defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)

	// To open a file, you have first to be authenticated to Google Cloud
	// Also we retrieve a file with http.Get
	// If there is an error, we treat it
	// Otherwise we return the file
	return f, nil
}

func (gcs *googleCloudStorage) Close() error {
	return gcs.client.Close()
}
