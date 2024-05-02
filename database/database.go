package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmdnubaidillah/go-rest-api/models"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_rest_api")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

}

func CreateTodo(t models.Todo, userId string) {

	_, err := db.Exec(fmt.Sprintf("INSERT INTO todo (id, todo, user_id) VALUES (%s, %s, %s)", t.TodoId, t.Todo, userId))

	if err != nil {
		panic(err.Error())
	}

}

func GetTodo() {

}
