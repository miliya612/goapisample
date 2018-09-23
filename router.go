package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func decorator(h func(r *http.Request) Responder) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := h(r)
		result.Write(w)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			HandlerFunc(decorator(Logging(route.HandlerFunc, route.Name)))
	}
	return router
}
