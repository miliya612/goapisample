package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func Logger(h httprouter.Handle, name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		h(w, r, ps)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.URL.Path,
			name,
			time.Since(start))
	}
}