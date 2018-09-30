package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/miliya612/goapisample/domain/model"
	"github.com/miliya612/goapisample/domain/repo"
	"github.com/miliya612/goapisample/presentation/httputil"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var KB int64 = 1024

type TodoHandler interface {
	TodoIndex(r *http.Request) httputil.Responder
	TodoShow(r *http.Request) httputil.Responder
	TodoCreate(r *http.Request) httputil.Responder
	TodoDelete(r *http.Request) httputil.Responder
}

type todoHandler struct {
	repo repo.Repository
}

func NewTodoHandler(repository repo.Repository) TodoHandler {
	return &todoHandler{repo: repository}
}

func (h *todoHandler) TodoIndex(_ *http.Request) httputil.Responder {
	todos, err := h.repo.GetAll()
	if err != nil {
		return httputil.Error(http.StatusInternalServerError, "something went wrong", err)
	}
	return httputil.Ok(todos)
}

func (h *todoHandler) TodoShow(r *http.Request) httputil.Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return httputil.Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	t, err := h.repo.GetByID(id)
	if err != nil {
		return httputil.Error(http.StatusInternalServerError, "something went wrong", err)
	}
	if t.ID == 0 && t.Name == "" {
		return httputil.Empty(http.StatusNotFound)
	}

	return httputil.Ok(t)
}

func (h *todoHandler) TodoCreate(r *http.Request) httputil.Responder {
	var todo model.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*KB)) // 1MB
	if err != nil {
		return httputil.Error(http.StatusInternalServerError, "request body is too large", err)
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &todo); err != nil {
		return httputil.Error(http.StatusInternalServerError, "failed marshalling json", err)
	}

	id, err := h.repo.Create(todo)
	if err != nil {
		return httputil.Error(http.StatusInternalServerError, "something went wrong", err)
	}
	todo.ID = id
	location := fmt.Sprintf("http://%s/%d", r.Host, id)
	return httputil.Created(todo, location)
}

func (h *todoHandler) TodoDelete(r *http.Request) httputil.Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return httputil.Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	if _, err = h.repo.Remove(id); err != nil {
		return httputil.Empty(http.StatusNotFound)
	}

	return httputil.Empty(http.StatusNoContent)
}

func parseTodoId(r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["todoId"])
	if err != nil {
		return -1, errors.New("todoId should be number.")
	}
	return id, nil
}
