package handler

import (
	"encoding/json"
	"log"
	"net/http"

	helper "github.com/hmdnubaidillah/go-rest-api/helpers"
	"github.com/hmdnubaidillah/go-rest-api/models"
)

var todo models.Todo

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	todo.TodoId = helper.GenerateUUID()

	err := json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		http.Error(w, "cant read request body", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	b, err := json.Marshal(todo)

	if err != nil {
		http.Error(w, "cant parse data to json", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}
