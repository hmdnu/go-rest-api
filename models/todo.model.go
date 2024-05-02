package models

type Todo struct {
	TodoId string `json:"todoId"`
	Todo   string `json:"todo"`
	UserId string `json:"userId"`
}

type User struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
}
