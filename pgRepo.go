package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type TodoRepo struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) Repository {
	return TodoRepo{db: db}
}

func (repo TodoRepo) GetAll() (todos Todos, err error) {
	err = repo.db.Find(&todos).Error
	if gorm.IsRecordNotFoundError(err) {
		err = ErrNotFound{}
	}
	return
}

func (repo TodoRepo) GetByID(id int) (todo Todo, err error) {
	todo = Todo{}
	err = repo.db.First(&todo, id).Error
	if gorm.IsRecordNotFoundError(err) {
		err = ErrNotFound{}
	}
	return
}

func (repo TodoRepo) Create(todo Todo) (int, error) {
	err := repo.db.Create(&todo).Error
	return todo.ID, err
}

func (repo TodoRepo) Update(todo Todo) (Todo, error) {
	t := Todo{}
	err := repo.db.Model(&t).Updates(todo).Error
	if gorm.IsRecordNotFoundError(err) {
		return t, ErrNotFound{}
	}
	return t, err
}

func (repo TodoRepo) Remove(id int) (int, error) {
	t := Todo{ID: id}
	i := repo.db.Delete(&t).RowsAffected
	if i == 0 {
		return id, ErrNotFound{}
	}
	return id, nil
}
