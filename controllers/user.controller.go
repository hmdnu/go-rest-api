package controllers

import (
	"net/http"

	"github.com/hmdnubaidillah/go-rest-api/handler"
)

func UserController(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case http.MethodGet:

	case http.MethodPost:
		handler.CreateUser(w, req)

	}
}
