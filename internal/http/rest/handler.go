package rest

import (
	"context"
	"fmt"
	"github.com/MadJlzz/gopypi/internal/listing"
	"github.com/gorilla/mux"
	"net/http"
)

func Handler(ctx context.Context, repository listing.Repository) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/simple/", index(ctx, repository))

	return r
}

func index(ctx context.Context, repository listing.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pkgsRef := repository.GetAllPackages(ctx)
		_, _ = fmt.Fprintln(w, pkgsRef)
	}
}
