package main

import (
	"flag"
	"github.com/MadJlzz/gopypi/internal/pkg/backend"
	"github.com/MadJlzz/gopypi/internal/pkg/template"
	"github.com/MadJlzz/gopypi/internal/pkg/web"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var credentials = flag.String("credentials", "credentials/service-account-dev.json", "GCP JSON credentials file.")

var packageLocation = flag.String("package-location", "C:/DefaultStorage", "Location from which we should load packages.")

func main() {
	tmpl := template.New()
	gcs := backend.NewGoogleCloudStorage(*packageLocation,"gopypi-nextgatetech-dev", *credentials)
	ctrl := web.New(gcs, tmpl)

	r := mux.NewRouter()
	r.HandleFunc("/simple/", ctrl.Index)
	r.HandleFunc("/simple/{name}/", ctrl.Package)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		_= gcs.Close()
		log.Fatal(err)
	}
}
