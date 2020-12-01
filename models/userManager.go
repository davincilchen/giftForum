package models

import (
	"fmt"
	"giftForum/basemodels"
	"giftForum/gentoken"
	"time"

	"github.com/patrickmn/go-cache"
	"golang.org/x/crypto/bcrypt"
)

var userManager *UserManager

var cacheExpiration = time.Duration(72 * time.Hour)

func Initialize(buf []byte) error {
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

	ps, err := t.getLoginUser(key)
	if err != nil {
		return nil, err
	}
	t.sessions.Delete(key)
	return ps, nil
}

func (t *UserManager) getLoginUser(key string) (*LoginUser, error) {

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

func GetLoginPlayer(key string) (*LoginUser, error) {

	if userManager == nil {
		return nil, fmt.Errorf("userManager is nil")
	}
	lUser, err := userManager.getLoginUser(key)
	if err != nil {
		return nil, err
	}
	return lUser, nil
}

func CreateUserAndLogin(email, password string) (*LoginUser, error) {

	user, err := CreateUser(email, password)
	if err != nil {
		return nil, err
	}
	return userLogin(user)

}

func CreateUser(email, password string) (*basemodels.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// u := &basemodels.User{
	// 	BaseUser: basemodels.BaseUser{
	// 		Email: email,
	// 	},
	// 	Password: string(hash),
	// }

	user, err := basemodels.CreateUser(email,string(hash))
	if err != nil {
		return nil, err
	}

	return user, nil
}




func GetUserWithID(id int) (*basemodels.User, error) {
	return basemodels.GetUserWithID(id)
}

func GetUser(email, password string) (*basemodels.User, error) {

	user, err := basemodels.GetUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUserSendGift(fromID, toID int) error{
	return basemodels.CreateUserSendGift(fromID, toID)
}

func GetTopTxUser() ([]basemodels.User, error){
	return basemodels.GetTopTxUser()
}

func GetTopRxUser() ([]basemodels.User, error){
	return basemodels.GetTopRxUser()
}

func userLogin(user *basemodels.User) (*LoginUser, error) {
	if userManager == nil {
		return nil, fmt.Errorf("userManager is nil")
	}

	u := &LoginUser{
		User: *user,
		UUID: gentoken.GenToken(),
	}

	err := userManager.userSessionLogin(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func UserLogin(email, password string) (*LoginUser, error) {

	user, err := GetUser(email, password)
	if err != nil {
		return nil, err
	}
	return userLogin(user)

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

func OauthUserDone(email string) (*LoginUser, error) {

	
	user, err := basemodels.GetUser(email)
	if err == nil {
		return userLogin(user)
	}

	password := gentoken.GenToken()
	user, err = CreateUser(email,password)
	
	if err != nil {
		return nil, err
	}

	return userLogin(user)
}