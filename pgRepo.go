package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) Repository {
	return TodoRepo{db: db}
}

func (repo TodoRepo) GetAll() (todos Todos, err error) {
	rows, err := repo.db.Query("select id, name, completed, due from todos")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due)
		if err != nil {
			return
		}
		todos = append(todos, todo)
	}
	return
}

func (repo TodoRepo) GetByID(id int) (todo Todo, err error) {
	todo = Todo{}
	err = repo.db.QueryRow("select id, name, completed, due from todos where id = $1", id).Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.Due)
	return
}

func (repo TodoRepo) Create(todo Todo) (int, error) {
	stmt, err := repo.db.Prepare("insert into todos (name, due) VALUES ($1, $2) returning id")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(todo.Name, todo.Due).Scan(&todo.ID)
	id := todo.ID
	return id, err
}

func (repo TodoRepo) Update(todo Todo) (Todo, error) {
	_, err := repo.db.Exec("update todos set name = $2, completed = $3, due = $4 where id = $1", todo.ID, todo.Name, todo.Completed, todo.Due)
	return todo, err
}

func (repo TodoRepo) Remove(id int) (int, error) {
	_, err := repo.db.Exec("delete from todos where id = $1", id)
	return id, err
}
