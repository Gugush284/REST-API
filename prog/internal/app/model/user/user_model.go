package model_user

// User struct ...
type User struct {
	id       int
	login    string
	password string
}

func New() *User {
	return &User{}
}