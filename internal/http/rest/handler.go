package rest

import (
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
	"github.com/MadJlzz/gopypi/internal/view"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type RepositoryHandler struct {
	logger     *zap.SugaredLogger
	template   *view.SimpleRepositoryTemplate
	repository listing.Repository
}

func NewRepositoryHandler(logger *zap.SugaredLogger, repository listing.Repository) *RepositoryHandler {
	return &RepositoryHandler{
		logger:     logger,
		template:   view.NewSimpleRepositoryTemplate(),
		repository: repository,
	}
}

func (rh RepositoryHandler) Router(ctx context.Context) http.Handler {
	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/simple/", rh.index(ctx))
	r.HandleFunc("/simple/{project}/", rh.project(ctx))
	return r
}

func (rh RepositoryHandler) index(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := rh.repository.GetAllProjects(ctx)
		if err := rh.template.Execute(w, "index", projects); err != nil {
			_ = fmt.Errorf("could not execute template [index]. [%v]\n", err)
			//Some fancy HTTP error code that is user friendly
		}
	}
}

func (rh RepositoryHandler) project(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projects := rh.repository.GetAllProjectPackages(ctx, vars["project"])
		if err := rh.template.Execute(w, "project", projects); err != nil {
			_ = fmt.Errorf("could not execute template [project]. [%v]\n", err)
			//Some fancy HTTP error code that is user friendly
		}
	}
}
