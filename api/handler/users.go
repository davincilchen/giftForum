package handler

import (
	"fmt"
	"errors"
	"strconv"
	"giftForum/api/ginprocess"
	"giftForum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	//googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func GetUsersSignIn(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, loginHTML, nil)
}

func CreateUserSendGift(ctx *gin.Context) {
	//TODO: 驗證是否已登入
	idString := ctx.Param("id")
	ID, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	idString = ctx.Param("from_id")
	fromID, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}


	models.CreateUserSendGift(fromID,ID)
	ctx.Redirect(http.StatusFound, "/")

}

func GetUser(ctx *gin.Context) {

	html := userHTML


	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	pageUser, err := models.GetUserWithID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}

	user, _ := ginprocess.GetLoginUserInGin(ctx)

	if user == nil {
		ctx.HTML(http.StatusOK, html, gin.H{
			GinHPageUser: pageUser,
		})
		return
	}

	if user.ID == pageUser.ID{
		ctx.HTML(http.StatusOK, html, gin.H{
			GinHPageUser: pageUser,
			GinHUser:  user,
		})
		return
	}
	///user/{{.user.id}}/to/{{.pageuser.id}}/gift
	
	
	path := fmt.Sprintf("/user/%d/to/%d/gift",user.ID, pageUser.ID)
	ctx.HTML(http.StatusOK, html, gin.H{
		GinHPageUser: pageUser,
		GinHUser:  user,
		GinHTxPotinPath: path,
	})
	
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

	loginUser, err := models.UserLogin(user.Email, user.Password)
	if err != nil {
		code = http.StatusUnauthorized
		err = errors.New("email不存在或password錯誤")
		return
	}

	ginprocess.SetUserSessionCookie(ctx, loginUser.UUID)
	ctx.Redirect(http.StatusFound, "/")
	// ctx.HTML(http.StatusOK, indexHTML, gin.H{
	// 	"success": "登入成功",
	// })
	return
}

func CreateUsersSignOut(ctx *gin.Context) {
	user, _ := ginprocess.GetLoginUserInGin(ctx)
	var err error
	defer func() {
		if err != nil {
			ctx.HTML(http.StatusOK, indexHTML, gin.H{
				"error": "登出失敗",
			})
			return
		}
	}()

	if user == nil {
		ctx.HTML(http.StatusOK, indexHTML, nil)
		return
	}
	_, err = models.UserLogout(user.UUID)
	if err != nil {
		return
	}

	ginprocess.CleanUserSessionCookie(ctx)
	ctx.Redirect(http.StatusFound, "/")
	// ctx.HTML(http.StatusOK, indexHTML, gin.H{
	// 	"success": "登出成功",
	// })
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

	loginUser, err := models.CreateUserAndLogin(user.Email, user.Password)
	if err != nil {
		code = http.StatusUnauthorized
		err = errors.New("建立帳號失敗")
		return
	}

	ginprocess.SetUserSessionCookie(ctx, loginUser.UUID)

	ctx.Redirect(http.StatusFound, "/")
	//ResposnSuccessHtmlWithUser(ctx, indexHTML, "註冊成功", &loginUser.BaseUser)

}
