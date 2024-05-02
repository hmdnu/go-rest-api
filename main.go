package main

import (
	"fmt"
	"net/http"

	"github.com/hmdnubaidillah/go-rest-api/controllers"
	_ "github.com/hmdnubaidillah/go-rest-api/database"
	"github.com/hmdnubaidillah/go-rest-api/middlewares"
)

func main() {
	http.HandleFunc("/todo", middlewares.Cors(controllers.TodoController))
	http.HandleFunc("/user", middlewares.Cors(controllers.UserController))

	fmt.Println("listen to port 8080")
	http.ListenAndServe(":8080", nil)
}
