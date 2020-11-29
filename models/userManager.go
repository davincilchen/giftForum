package models

type UserManager struct {
}

func CreateUser(email, password string) *User {
	return &User{}
}

func GetUser(id int) (*User, error) {
	return &User{}, nil
}

func UserLogin(email, password string) (*LoginUser, error) {
	u := &LoginUser{
		User: User{
			Email: email,
		},
		Token: "",
	}
	return u, nil
}

func UserLogout(user *LoginUser) error {
	return nil
}
