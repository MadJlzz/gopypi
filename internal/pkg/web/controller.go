package web

import (
	"github.com/MadJlzz/gopypi/internal/pkg/backend"
	"github.com/MadJlzz/gopypi/internal/pkg/template"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	storage  *backend.GoogleCloudStorage
	template *template.SimpleRepositoryTemplate
}

func New(gcs *backend.GoogleCloudStorage, tmpl *template.SimpleRepositoryTemplate) *Controller {
	return &Controller{
		storage:  gcs,
		template: tmpl,
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	pkgs := c.storage.Load()
	if err := c.template.Execute(w, "index", pkgs); err != nil {
		log.Errorf("could not execute template [index]. [%v]\n", err)
		//Some fancy HTTP error code that is user friendly
	}
}

func (c *Controller) Package(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _, found := vars["name"]; !found {
		log.Errorln("routing variable [name] wasn't provided")
		//Some fancy HTTP error code
	}
	// TODO: put in place a cache instead of loading everytime from the source.
	pkgs := c.storage.Load()
	pkg, found := pkgs[vars["name"]]
	if !found {
		log.Errorf("package [%s] is not available anymore...\n", vars["name"])
		//Some fancy HTTP error code
	}
	//pkg.Files[0] = "C:/DefaultStorage/example-pkg/example-pkg-0.0.1.tar.gz"
	if err := c.template.Execute(w, "package", pkg); err != nil {
		log.Errorf("could not execute template [package]. [%v]\n", err)
		//Some fancy HTTP error code that is user friendly
	}
}
