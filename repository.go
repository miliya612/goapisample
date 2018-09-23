package main

type Repository interface {
	GetAll() (Todos, error)
	GetByID(int) (Todo, error)
	Create(todo Todo) (int, error)
	Update(todo Todo) (Todo, error)
	Remove(id int) (int, error)
}
