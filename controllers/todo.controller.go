package controllers

import (
	"net/http"

	"github.com/hmdnubaidillah/go-rest-api/handler"
)

func TodoController(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case http.MethodGet:

	case http.MethodPost:
		handler.CreateTodo(w, req)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
