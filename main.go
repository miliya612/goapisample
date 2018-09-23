package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()

	log.Printf("server started at: %v", time.Now())
	log.Fatal(http.ListenAndServe(":8080", router))
}
