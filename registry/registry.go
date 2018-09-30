package registry

import (
	"database/sql"
	"github.com/miliya612/goapisample/domain/repo"
	"github.com/miliya612/goapisample/infra/persistance/datastore"
	"github.com/miliya612/goapisample/presentation/handler"
)

type Registration struct {}

type Registerer interface {
	InjectDBCon() *sql.DB
	InjectTodo() repo.Repository
	InjectTodoHandler() handler.TodoHandler
}

func (r *Registration) RegisterDBCon() *sql.DB {
	db, err := sql.Open("postgres", "user=todoapp dbname=todoapp password=todopass sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

func (r *Registration) RegisterTodoRepo() repo.Repository {
	return datastore.NewTodoRepo(r.RegisterDBCon())
}

func (r *Registration) RegisterTodoHandler() handler.AppHandler {
	return handler.NewTodoHandler(r.RegisterTodoRepo())
}