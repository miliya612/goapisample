package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome!")
}

func ToDoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func ToDoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "ToDo show: %s", ps.ByName("todoId"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/todos", ToDoIndex)
	router.GET("/todos/:todoId", ToDoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}
