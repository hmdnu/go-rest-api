package models

type Todo struct {
	TodoId string `json:"todoId"`
	Todo   string `json:"todo"`
}

type User struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}
