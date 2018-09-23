package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type MyHandle func(*http.Request) Responder

func IDShouldBeInt(h func(r *http.Request) Responder, name string) MyHandle {
	return Logging(func(r *http.Request) Responder {
		_, err := strconv.Atoi(mux.Vars(r)["todoId"])
		if err != nil {
			return Error(422, "todoId should be number", err)
		}

		return h(r)
	}, name)
}
