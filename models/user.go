package models

type User struct {
	Email string
}

type LoginUser struct {
	User
	Token string
}
