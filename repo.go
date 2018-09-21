package main

import "fmt"

var currentID int

var todos Todos

func init() {

}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) int {
	currentID += 1
	t.ID = currentID
	todos = append(todos, t)
	return currentID
}

func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Todo with if of %d to delete", id)
}