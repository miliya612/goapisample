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

func Index(_ *http.Request) Responder {
	return Respond(http.StatusOK, "welcome")
}

func TodoIndex(_ *http.Request) Responder {
	return Ok(todos)
}

func TodoShow(r *http.Request) Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		return Empty(http.StatusNotFound)
	}

	return Ok(t)
}

func TodoCreate(r *http.Request) Responder {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024*KB)) // 1MB
	if err != nil {
		return Error(http.StatusInternalServerError, "request body is too large", err)
	}
	defer r.Body.Close()

	if err = json.Unmarshal(body, &todo); err != nil {
		return Error(http.StatusInternalServerError, "failed marshalling json", err)
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	return Created(t, location)
}

func TodoDelete(r *http.Request) Responder {
	id, err := parseTodoId(r)
	if err != nil {
		return Error(http.StatusUnprocessableEntity, "invalid parameter", err)
	}

	if err = RepoDestroyTodo(id); err != nil {
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
