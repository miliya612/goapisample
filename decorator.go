package main

import (
	"net/http"
)

type APIHandle func(*http.Request) Responder
