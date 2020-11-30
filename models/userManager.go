package models

import (
	"fmt"
	"giftForum/gentoken"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

var userManager *UserManager

var cacheExpiration = time.Duration(72 * time.Hour)

func Initialize() error {
	p := &UserManager{}
	p.sessions = cache.New(cacheExpiration, 1*time.Minute)
	userManager = p
	return nil
}

func Uninitialize() error {
	userManager = nil
	return nil
}

type UserManager struct {
	sessions *cache.Cache //key: UUID
}

func TransUserSession(userSession interface{}) (*LoginUser, error) {

	ps, ok := userSession.(*LoginUser)
	if !ok {
		err := fmt.Errorf(
			"TransUserSession userSession(*LoginUser) failed for interface %#v",
			userSession)
		return nil, err
	}
	return ps, nil
}

func (t *UserManager) userSessionLogout(key string) (*LoginUser, error) {

	s, ok := t.sessions.Get(key)
	if !ok {
		return nil, fmt.Errorf("User session is not found")
	}
	ps, err := TransUserSession(s)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (t *UserManager) userSessionLogin(user *LoginUser) error {

	t.sessions.Set(user.UUID, user, cache.DefaultExpiration)

	return nil
}

func CreateUser(email, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Email:    email,
		Password: string(hash),
	}

	user, err := createUser(u)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createUser(user *User) (*User, error) {

	return user, nil
}

func GetUserByID(id int) (*User, error) {
	return &User{}, nil
}

func GetUser(email, password string) (*User, error) {

	user, err := getUser(email, password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func getUser(email, password string) (*User, error) {
	return &User{}, nil
}

func UserLogin(email, password string) (*LoginUser, error) {
	if userManager == nil {
		return nil, fmt.Errorf("userManager is nil")
	}
	user, err := GetUser(email, password)
	if err != nil {
		return nil, err
	}
	u := &LoginUser{
		User: *user,
		UUID: gentoken.GenToken(),
	}

	err = userManager.userSessionLogin(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserLogout(uuid string) (*LoginUser, error) {
	if userManager == nil {
		return nil, fmt.Errorf("userManager is nil")
	}
	lUser, err := userManager.userSessionLogout(uuid)
	if err != nil {
		return nil, err
	}
	return lUser, nil
}
