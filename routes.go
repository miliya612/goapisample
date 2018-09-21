package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

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

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", ToDoIndex},
	Route{"TodoShow", "GET", "/todos/:todoId", ToDoShow},
}
