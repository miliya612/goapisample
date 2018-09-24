package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
)

func NewRouter() *chi.Mux {
	router := chi.NewMux()
	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	for _, route := range routes {
		router.Method(
			route.Method,
			route.Path,
			decorator(route.HandlerFunc),
			)
	}
	return router
}
