package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *AppServer) publicRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/").Subrouter()

	subRouter.Path("/").Handler(a.Service).Methods("GET")
	subRouter.Path("/healthcheck").HandlerFunc(healthcheck).Methods("GET")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `product service is good`)
}
