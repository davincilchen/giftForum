package models

type User struct {
	ID       int
	Email    string
	Password string
	RxPoint  int
	TxPoint  int
}

type LoginUser struct {
	User
	UUID string
}
