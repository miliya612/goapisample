package decorator

import (
	"github.com/miliya612/goapisample/presentation/httputil"
	"net/http"
)

type APIHandleFunc func(*http.Request) httputil.Responder

func Decorator(f APIHandleFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := f(r)
		result.Write(w)
	}
}
