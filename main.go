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

var packageLocation = flag.String("package-location", "C:/DefaultStorage", "Location from which we should load packages.")

func main() {
	tmpl := template.New()
	ls := backend.NewLocalStorage(*packageLocation)
	ctrl := web.New(ls, tmpl)

	r := mux.NewRouter()
	r.HandleFunc("/", ctrl.Index)
	r.HandleFunc("/simple/{name}/", ctrl.Package)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}


