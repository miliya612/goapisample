package main

import (
	"net/http"
)

type APIHandle func(*http.Request) Responder

func decorator(h APIHandle) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := h(r)
		result.Write(w)
	}
}
