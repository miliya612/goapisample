package routes

import (
	"github.com/miliya612/goapisample/presentation/decorator"
	"github.com/miliya612/goapisample/presentation/handler"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc decorator.APIHandleFunc
}

type Routes []Route

func getRoutes(app handler.AppHandler) []Route {
	return Routes{
		Route{"TodoIndex", "GET", "/todos", app.TodoIndex},
		Route{"TodoShow", "GET", "/todos/{todoId}", app.TodoShow},
		Route{"TodoCreate", "POST", "/todos", app.TodoCreate},
		Route{"TodoDelete", "DELETE", "/todos/{todoId}", app.TodoDelete},
	}
}