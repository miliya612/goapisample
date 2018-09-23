package main

import (
	"net/http"
)

type MyHandle func(*http.Request) Responder
