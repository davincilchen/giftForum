package models

type UserManager struct {
}

//var userManager UserManager

// func GetUserManager() *UserManager {
// 	return &userManager
// }

func CreateUser(email, password string) *User {
	return &User{}
}

func UserLogin(user *User) (*LoginUser, error) {
	u := &LoginUser{
		User:  *user,
		Token: "",
	}
	return u, nil
}

func UserLogout(user *LoginUser) error {
	return nil
}
