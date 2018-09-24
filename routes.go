package main

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc APIHandle
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
	Route{"TodoIndex", "GET", "/todos", app.handler.TodoIndex},
	Route{"TodoShow", "GET", "/todos/{todoId}", app.handler.TodoShow},
	Route{"TodoCreate", "POST", "/todos", app.handler.TodoCreate},
	Route{"TodoDelete", "DELETE", "/todos/{todoId}", app.handler.TodoDelete},
}
