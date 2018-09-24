package main

import "database/sql"

type Injection struct {
	handler  Handler
	todoRepo Repository
}

func Inject(db *sql.DB) Injection {
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

func injectTodoRepo(db *sql.DB) Repository {
	return NewTodoRepo(db)
}
