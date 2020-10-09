package main

import (
	"flag"
	"github.com/MadJlzz/gopypi/internal/pkg/backend"
	"github.com/MadJlzz/gopypi/internal/pkg/template"
	_ "github.com/MadJlzz/gopypi/internal/pkg/utils"
	"github.com/MadJlzz/gopypi/internal/pkg/web"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func main() {
	var (
		bucket      = flag.String("bucket", "gopypi-nextgatetech-dev", "GCS Bucket name to use.")
		credentials = flag.String("credentials", "credentials/service-account-dev.json", "GCP JSON credentials file.")
		port        = flag.String("port", "3000", "Port of the web server.")
	)
	flag.Parse()

	tmpl := template.New()
	gcs := backend.NewGoogleCloudStorage(*bucket, *credentials)
	ctrl := web.New(gcs, tmpl)

	r := mux.NewRouter()
	r.HandleFunc("/simple/", ctrl.Index)
	r.HandleFunc("/simple/{name}/", ctrl.Package)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Auto-indexing packages from [%s] on port [%s]...\n", *bucket, *port)
	log.Infoln("Press CTRL+C to kill server...")
	if err := srv.ListenAndServe(); err != nil {
		_ = gcs.Close()
		log.Fatal(err)
	}
}
