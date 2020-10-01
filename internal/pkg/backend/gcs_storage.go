package backend

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/MadJlzz/gopypi/internal/pkg/model"
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"net/url"
	"path/filepath"
)

type GoogleCloudStorage struct {
	localDir string
	bucket   string
	client   *storage.Client
}

func NewGoogleCloudStorage(localDir, bucket string, opts ...option.ClientOption) *GoogleCloudStorage {
	c, err := storage.NewClient(context.Background(), opts...)
	if err != nil {
		logrus.Fatalf("could not create google cloud storage client.\ngot: [%v]", err)
	}
	return &GoogleCloudStorage{
		localDir: localDir,
		bucket: bucket,
		client: c,
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
			utils.Logger.Warnf("could not read objects from bucket [%s].\ngot: [%v]", gcs.bucket, err)
		}

		u, err := url.Parse(attrs.MediaLink)
		if err != nil {
			utils.Logger.Warnf("could not parse object [%s] download url [%s].\ngot: [%v]", attrs.Name, attrs.MediaLink, err)
		}

		pf := &model.PackageFile{
			Name: filepath.Base(attrs.Name),
			URL:  u,
		}

		key := filepath.Dir(attrs.Name)
		if _, found := pkgs[key]; found {
			pkgs[key].AppendPackageFile(pf)
		} else {
			pkgs[key] = model.New(key, pf)
		}
	}
	return pkgs
}

//func (gcs *GoogleCloudStorage) Open(name string) (http.File, error) {
	//f := &os.File{}
	//for _, name := range names {
	//	rc, err := bh.Object(name).NewReader(context.TODO())
	//	if err != nil {
	//		logrus.Warnln(err)
	//	}
	//
	//	data, err := ioutil.ReadAll(rc)
	//	if err != nil {
	//		logrus.Warnln(err)
	//	}
	//	logrus.Println(data)
	//	_ = rc.Close()
	//}
	//resp, _ := http.Get("https://google.com")
	//defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)

	// To open a file, you have first to be authenticated to Google Cloud
	// Also we retrieve a file with http.Get
	// If there is an error, we treat it
	// Otherwise we return the file
	//return f, nil
//}

func (gcs *GoogleCloudStorage) Close() error {
	return gcs.client.Close()
}
