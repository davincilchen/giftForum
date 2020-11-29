package models

type User struct {
	ID    int
	Email string
}

type LoginUser struct {
	User
	Token string
}
