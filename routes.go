package main

import "github.com/julienschmidt/httprouter"

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoShow", "GET", "/todos/:todoId", TodoShow},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
	Route{"TodoDelete", "DELETE", "/todos/:todoId", TodoDelete},
}
