package middleware

import (
	"github.com/miliya612/goapisample/presentation/decorator"
	"github.com/miliya612/goapisample/presentation/httputil"
	"log"
	"net/http"
	"time"
)

func logger(status int, method, uri, name string, start time.Time) {
	log.Printf("%d\t%s\t%s\t%s\t%s", status, method, uri, name, time.Since(start))
}

func Logging(h decorator.APIHandleFunc, name string) decorator.APIHandleFunc {
	return func(r *http.Request) httputil.Responder {
		start := time.Now()
		result := h(r)
		logger(result.Status(), r.Method, r.URL.EscapedPath(), name, start)
		return result
	}
}
