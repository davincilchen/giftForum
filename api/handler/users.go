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

	var (
		email    string
		password string
	)

	in := ""
	if in, isExist := ctx.GetPostForm("email"); !isExist || in == "" {
		err = errors.New("必須輸入email")
		return
	}
	email = in

	if in, isExist := ctx.GetPostForm("password"); !isExist || in == "" {
		err = errors.New("必須輸入password")
		return
	}
	password = in

	_, err = models.UserLogin(email, password)
	if err != nil {
		code = http.StatusUnauthorized
		err = errors.New("email不存在或password錯誤")
		return
	}

	ctx.HTML(http.StatusOK, loginHTML, gin.H{
		"success": "登入成功",
	})
	return
}

func CreateUsersSignOut(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": "登出成功",
	})
}

func GetUsersSignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetUsersSignUp")
}
func CreateUsersSignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateUsersSignUp")
}
