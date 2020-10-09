package web

import (
	"github.com/MadJlzz/gopypi/internal/pkg/backend"
	"github.com/MadJlzz/gopypi/internal/pkg/template"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Controller exposes endpoints for the web server.
type Controller struct {
	storage  *backend.GoogleCloudStorage
	template *template.SimpleRepositoryTemplate
}

// New is the simplest way to get started with a Controller.
func New(gcs *backend.GoogleCloudStorage, tmpl *template.SimpleRepositoryTemplate) *Controller {
	return &Controller{
		storage:  gcs,
		template: tmpl,
	}
}

// Index generates the auto-index page from the loaded packages.
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	pkgs := c.storage.Load()
	if err := c.template.Execute(w, "index", pkgs); err != nil {
		log.Errorf("could not execute template [index]. [%v]\n", err)
		//Some fancy HTTP error code that is user friendly
	}
}

// Package generates an HTML page listing URLs for downloading
// a Package's files.
//
// TODO: put in place a cache instead of loading everytime from the source.
func (c *Controller) Package(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, found := vars["name"]; !found {
		log.Errorln("routing variable [name] wasn't provided")
		//Some fancy HTTP error code
	}
	pkgs := c.storage.Load()
	pkg, found := pkgs[vars["name"]]
	if !found {
		log.Errorf("package [%s] is not available anymore...\n", vars["name"])
		//Some fancy HTTP error code
	}
	if err := c.template.Execute(w, "package", pkg); err != nil {
		log.Errorf("could not execute template [package]. [%v]\n", err)
		//Some fancy HTTP error code that is user friendly
	}
}
