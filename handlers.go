package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var KB int64 = 1024

type TodoHandler struct {
	repo Repository
}

func NewTodoHandler(repository Repository) TodoHandler {
	return TodoHandler{repo: repository}
}

func Index(_ *http.Request) Responder {
	return Respond(http.StatusOK, "welcome")
}

func (h *TodoHandler) TodoIndex(_ *http.Request) Responder {
	todos, err := h.repo.GetAll()
	if err != nil {
		return Error(http.StatusInternalServerError, "something went wrong", err)
	}
	return Ok(todos)
}

func (h *TodoHandler) TodoShow(r *http.Request) Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	t, err := h.repo.GetByID(id)
	if err != nil {
		return Error(http.StatusInternalServerError, "something went wrong", err)
	}
	if t.ID == 0 && t.Name == "" {
		return Empty(http.StatusNotFound)
	}

	return Ok(t)
}

func (h *TodoHandler) TodoCreate(r *http.Request) Responder {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*KB)) // 1MB
	if err != nil {
		return Error(http.StatusInternalServerError, "request body is too large", err)
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &todo); err != nil {
		return Error(http.StatusInternalServerError, "failed marshalling json", err)
	}

	id, err := h.repo.Create(todo)
	if err != nil {
		return Error(http.StatusInternalServerError, "something went wrong", err)
	}
	todo.ID = id
	location := fmt.Sprintf("http://%s/%d", r.Host, id)
	return Created(todo, location)
}

func (h *TodoHandler) TodoDelete(r *http.Request) Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	if _, err = h.repo.Remove(id); err != nil {
		return Empty(http.StatusNotFound)
	}

	return Empty(http.StatusNoContent)
}

func parseTodoId(r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["todoId"])
	if err != nil {
		return -1, errors.New("todoId should be number.")
	}
	return id, nil
}
