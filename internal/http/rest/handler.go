package rest

import (
	"fmt"
	"github.com/MadJlzz/gopypi/internal/registry"
	"github.com/MadJlzz/gopypi/internal/view"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type RepositoryHandler struct {
	logger   *zap.SugaredLogger
	template *view.SimpleRepositoryTemplate
	registry registry.Registry
}

func NewRepositoryHandler(logger *zap.SugaredLogger, repository registry.Registry) *RepositoryHandler {
	return &RepositoryHandler{
		logger:   logger,
		template: view.NewSimpleRepositoryTemplate(),
		registry: repository,
	}
}

func (rh RepositoryHandler) Router() http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/simple/", http.StatusMovedPermanently)
	})
	r.HandleFunc("/simple/", rh.index())
	r.HandleFunc("/simple/{project}/", rh.project())
	return r
}

func (rh RepositoryHandler) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := rh.registry.GetAllProjects()
		if err := rh.template.Execute(w, "index", projects); err != nil {
			_ = fmt.Errorf("could not execute template [index]. [%v]\n", err)
			//Some fancy HTTP error code that is user friendly
		}
	}
}

func (rh RepositoryHandler) project() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projects := rh.registry.GetAllProjectPackages(vars["project"])
		if err := rh.template.Execute(w, "project", projects); err != nil {
			_ = fmt.Errorf("could not execute template [project]. [%v]\n", err)
			//Some fancy HTTP error code that is user friendly
		}
	}
}
