package main

import (
	"github.com/jinzhu/gorm"
)

type Injection struct {
	handler  Handler
	todoRepo Repository
}

func Inject(db *gorm.DB) Injection {
	todoRepo := injectTodoRepo(db)
	todoHandler := injectTodoHandler(todoRepo)
	handler := injectHandler(todoHandler)
	return Injection{
		handler:  handler,
		todoRepo: todoRepo,
	}
}

func injectHandler(todo TodoHandler) Handler {
	return NewHandler(todo)
}

func injectTodoHandler(repository Repository) TodoHandler {
	return NewTodoHandler(repository)
}

func injectTodoRepo(db *gorm.DB) Repository {
	return NewTodoRepo(db)
}
