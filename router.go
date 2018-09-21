package main

import "github.com/julienschmidt/httprouter"

func NewRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle
		handle = route.HandlerFunc
		handle = Logger(handle, route.Name)
		router.Handle(route.Method, route.Path, handle)
	}
	return router
}
