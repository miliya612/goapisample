package handler

import (
	"github.com/miliya612/goapisample/presentation/httputil"
	"net/http"
)

type AppHandler interface {
	TodoHandler
}

type APIHandleFunc func(*http.Request) httputil.Responder