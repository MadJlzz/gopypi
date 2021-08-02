package rest

import (
	"github.com/MadJlzz/gopypi/internal/auth"
	"github.com/MadJlzz/gopypi/internal/registry"
	"github.com/MadJlzz/gopypi/internal/view"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
)

func Handler(logger *zap.SugaredLogger, tpl *view.SimpleRepositoryTemplate, rg registry.Registry) http.Handler {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/", redirectHandler())
	router.HandleFunc("/simple/", indexHandler(logger, tpl, rg))
	router.HandleFunc("/simple/{project}/", projectPackagesHandler(logger, tpl, rg))

	if _, ok := os.LookupEnv("NODE_ENV"); ok {
		authMiddleware := auth.NewAuthenticationMiddleware(logger)
		//router.Use(authMiddleware.HandleCloudIAPAuthentication)
		router.Use(authMiddleware.HandleBasicAuthentication(&auth.ApiKey{}))
		router.Use(authMiddleware.HandleNoAuthentication)
	}

	return router
}

func redirectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/simple/", http.StatusMovedPermanently)
	}
}

func indexHandler(logger *zap.SugaredLogger, tpl *view.SimpleRepositoryTemplate, rg registry.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects := rg.GetAllProjects()
		if err := tpl.Execute(w, "index", projects); err != nil {
			logger.Errorf("could not execute template [index]. got: %v", err)
			http.Error(w, "the 'index' page could not be generated", http.StatusInternalServerError)
		}
	}
}

func projectPackagesHandler(logger *zap.SugaredLogger, tpl *view.SimpleRepositoryTemplate, rg registry.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		projects := rg.GetAllProjectPackages(vars["project"])
		if err := tpl.Execute(w, "project-packages", projects); err != nil {
			logger.Errorf("could not execute template [project-packages]. got: %v", err)
			http.Error(w, "the 'project-packages' page could not be generated", http.StatusInternalServerError)
		}
	}
}
