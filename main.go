package main

import (
	"flag"
	"github.com/MadJlzz/gopypi/internal/pkg/backend"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"time"
)

var credentials = flag.String("credentials", "credentials/service-account-dev.json", "GCP JSON credentials file.")

var packageLocation = flag.String("package-location", "C:/DefaultStorage", "Location from which we should load packages.")

func main() {
	//tmpl := template.New()
	//ls := backend.NewLocalStorage(*packageLocation)
	gcs := backend.NewGoogleCloudStorage(*packageLocation,"gopypi-nextgatetech-dev", option.WithCredentialsFile(*credentials))
	//ctrl := web.New(ls, tmpl)

	r := mux.NewRouter()
	//r.HandleFunc("/", ctrl.Index)
	//r.HandleFunc("/simple/{name}/", ctrl.Package)

	r.PathPrefix("/simple/").Handler(http.StripPrefix("/simple/", http.FileServer(http.Dir(*packageLocation))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		_= gcs.Close()
		log.Fatal(err)
	}
}


