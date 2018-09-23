package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	f, err := os.OpenFile("tmp/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}

	log.SetOutput(f)
}

func main() {
	router := NewRouter()

	log.Printf("server started at: %v", time.Now())
	log.Fatal(http.ListenAndServe(":8080", router))
}
