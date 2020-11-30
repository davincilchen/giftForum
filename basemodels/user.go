package basemodels

type BaseUser struct {
	ID       int
	Email    string
	Password string
	RxPoint  int
	TxPoint  int
}
type User struct {
	BaseUser
	Password string
}
