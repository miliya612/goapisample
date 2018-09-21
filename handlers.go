package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var KB int64 = 1024

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idParam := ps.ByName("todoId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}

func TodoCreate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024 * KB)) // 1MB
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.Header().Add("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}
