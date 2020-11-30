package basemodels

type BaseUser struct {
	ID      int
	Email   string
	RxPoint int
	TxPoint int
}
type User struct {
	BaseUser
	Password string
}
