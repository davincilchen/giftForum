package handler

import (
	"errors"
	"giftForum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersSignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, loginHTML, nil)
}

func GetUser(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, userHTML, nil)
}

type User struct {
	Password string
	Email    string
}

func GetUserForm(ctx *gin.Context) (*User, error) {
	user := &User{}

	in := ""
	in, isExist := ctx.GetPostForm("email")
	if !isExist || in == "" {
		return nil, errors.New("必須輸入email")

	}
	user.Email = in

	in = ""
	in, isExist = ctx.GetPostForm("password")
	if !isExist || in == "" {
		return nil, errors.New("必須輸入password")

	}

	user.Password = in

	return user, nil

}

func CreateUsersSignIn(ctx *gin.Context) {
	var err error
	code := http.StatusBadRequest

	defer func() {
		if err == nil {
			return
		}
		ctx.HTML(code, loginHTML, gin.H{
			"error": err,
		})
	}()

	user, err := GetUserForm(ctx)
	if err != nil {
		return
	}

	_, err = models.UserLogin(user.Email, user.Password)
	if err != nil {
		code = http.StatusUnauthorized
		err = errors.New("email不存在或password錯誤")
		return
	}

	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": "登入成功",
	})
	return
}

func CreateUsersSignOut(ctx *gin.Context) {
	//UserLogout()
	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": "登出成功",
	})
}

func GetUsersSignUp(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, registerHTML, nil)
}
func CreateUsersSignUp(ctx *gin.Context) {
	var err error
	code := http.StatusBadRequest

	defer func() {
		if err == nil {
			return
		}
		ctx.HTML(code, registerHTML, gin.H{
			"error": err,
		})
	}()

	user, err := GetUserForm(ctx)
	if err != nil {
		return
	}

	in := ""
	in, _ = ctx.GetPostForm("checkpassword")
	if user.Password != in {
		err = errors.New("輸入的password不一致")
		return
	}

	_, err = models.CreateUser(user.Email, user.Password)
	if err != nil {
		code = http.StatusUnauthorized
		err = errors.New("建立帳號失敗")
		return
	}

	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": "註冊成功",
	})
}
