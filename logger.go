package main

import (
	"log"
	"net/http"
	"time"
)

func logger(status int, method, uri, name string, start time.Time) {
	log.Printf("%d\t%s\t%s\t%s\t%s", status, method, uri, name, time.Since(start))
}

func Logging(h APIHandle, name string) APIHandle {
	return func(r *http.Request) Responder {
		start := time.Now()
		result := h(r)
		logger(result.Status(), r.Method, r.URL.EscapedPath(), name, start)
		return result
	}
}
