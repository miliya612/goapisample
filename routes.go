package main

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc MyHandle
}

type Routes []Route

type Handler struct {
	TodoHandler
}

func NewHandler(todo TodoHandler) Handler {
	return Handler{todo}
}

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"TodoIndex", "GET", "/todos", injection.handler.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", injection.handler.TodoShow},
	Route{"TodoCreate", "POST", "/todos", injection.handler.TodoCreate},
	Route{"TodoDelete", "DELETE", "/todos/{todoId}", injection.handler.TodoDelete},
}
