package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

	id, _ := strconv.Atoi(mux.Vars(r)["todoId"])

	t := RepoFindTodo(id)
	fmt.Println(t)
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

	if err := json.Unmarshal(body, &todo); err != nil {
		return Error(http.StatusInternalServerError, "failed marshalling json", err)
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	return Created(t, location)
}

func TodoDelete(r *http.Request) Responder {
	id, _ := strconv.Atoi(mux.Vars(r)["todoId"])

	if err := RepoDestroyTodo(id); err != nil {
		return Empty(http.StatusNotFound)
	}

	return Empty(http.StatusNoContent)
}
