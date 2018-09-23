package main

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc MyHandle
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", TodoShow},
	Route{"TodoCreate", "POST", "/todos", TodoCreate},
	Route{"TodoDelete", "DELETE", "/todos/{todoId}", TodoDelete},
}
