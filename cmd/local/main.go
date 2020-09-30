package main

import (
	"flag"
	"github.com/MadJlzz/gopypi/internal/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

const PypiBaseUrl = "/simple/"

var port = flag.String("port", "3000", "Port of the app")

var packageLocation = flag.String("package-location", "C:/DefaultStorage", "Location from which we should load packages.")

func main() {

	if _, err := os.Stat(*packageLocation); os.IsNotExist(err) {
		utils.Logger.Fatalf("directory [%s] doesn't exist.\ngot: [%v]", *packageLocation, err)
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

	utils.Logger.Infof("Static file server scanning directory [\"%s\"] started on port [%s]...\n", *packageLocation, *port)
	utils.Logger.Fatal(srv.ListenAndServe())
}
