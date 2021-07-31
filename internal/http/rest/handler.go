package rest

import (
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
	"github.com/MadJlzz/gopypi/internal/view"
	"github.com/gorilla/mux"
	"net/http"
)

func Handler(ctx context.Context, repository listing.Repository) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/simple/", index(ctx, repository))

	return r
}

func index(ctx context.Context, repository listing.Repository) func(w http.ResponseWriter, r *http.Request) {
	tpl := view.NewSimpleRepositoryTemplate()
	return func(w http.ResponseWriter, r *http.Request) {
		pkgsRef := repository.GetAllPackages(ctx)
		if err := tpl.Execute(w, "index", pkgsRef); err != nil {
			_ = fmt.Errorf("could not execute template [index]. [%v]\n", err)
			//Some fancy HTTP error code that is user friendly
		}
	}
}
