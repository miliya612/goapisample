package repo

import "github.com/miliya612/goapisample/domain/model"

type TodoRepo interface {
	GetAll() (model.Todos, error)
	GetByID(int) (model.Todo, error)
	Create(todo model.Todo) (int, error)
	Update(todo model.Todo) (model.Todo, error)
	Remove(id int) (int, error)
}
