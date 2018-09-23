package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func logger(method, uri, name string, start time.Time) {
	log.Printf("%s\t%s\t%s\t%s", method, uri, name, time.Since(start))
}

func Logging(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		logger(r.Method, r.URL.EscapedPath(), name, start)
	}
}
