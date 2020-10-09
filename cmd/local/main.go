package main

import (
	"flag"
	_ "github.com/MadJlzz/gopypi/internal/pkg/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func main() {
	const PypiBaseUrl = "/simple/"
	var (
		port            = flag.String("port", "3000", "Port of the app")
		packageLocation = flag.String("package-location", "C:/DefaultStorage", "Location from which we should load packages.")
	)
	flag.Parse()

	if _, err := os.Stat(*packageLocation); os.IsNotExist(err) {
		log.Fatalf("directory [%s] doesn't exist.\ngot: [%v]", *packageLocation, err)
	}

	r := mux.NewRouter()
	r.PathPrefix(PypiBaseUrl).
		Handler(
			http.StripPrefix(PypiBaseUrl,
				http.FileServer(
					http.Dir(*packageLocation),
				),
			),
		)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Infof("Static file server scanning directory [\"%s\"] started on port [%s]...\n", *packageLocation, *port)
	log.Fatal(srv.ListenAndServe())
}
