package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

var app Injection

func init() {
	f, err := os.OpenFile("tmp/application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	log.SetOutput(f)

}

func main() {
	router := NewRouter()
	db, err := gorm.Open("postgres", "user=todoapp dbname=todoapp password=todopass sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Todo{})
	db.LogMode(true)

	f, err := os.OpenFile("tmp/db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	gl := gorm.Logger{
		LogWriter: log.New(f, "\r\n", 0),
	}
	db.SetLogger(gl)

	app = Inject(db)

	log.Printf("server started at: %v", time.Now())
	log.Fatal(http.ListenAndServe(":8080", router))
}
