package main

import (
	"net/http"
)

type APIHandle func(*http.Request) Responder

func decorator(h APIHandle) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result := h(r)
		result.Write(w)
	})
}
