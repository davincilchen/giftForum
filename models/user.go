package models

type User struct {
}

type LoginUser struct {
	User
	Token string
}
