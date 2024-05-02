package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hmdnubaidillah/go-rest-api/models"
)

var user models.User

func CreateUser(w http.ResponseWriter, req *http.Request) {

	// read all request
	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		http.Error(w, "cant read request body", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	// convert to json
	v, err := json.Marshal(&user)

	if err != nil {
		http.Error(w, "cant parse data to json", http.StatusInternalServerError)
		log.Fatal(err.Error())
	}

	w.Write(v)
}
