package ginprocess

import (
	"fmt"
	"giftForum/models"

	"github.com/gin-gonic/gin"
)

func GetLoginUserInGin(ctx *gin.Context) (*models.LoginUser, error) {

	info, ok := ctx.Get(GinKeyLoginUser)
	if !ok {
		err := fmt.Errorf("GinKeyLoginUser is not find")
		return nil, err
	}

	player, ok := info.(*models.LoginUser)
	if !ok {
		err := fmt.Errorf("Trans GinKeyLoginUser.(*LoginUser) failed from gin cache")
		return nil, err
	}

	return player, nil
}

func CachePlayerSessionInGin(ctx *gin.Context, user *models.LoginUser) {
	ctx.Set(GinKeyLoginUser, user)
}

func GetMessageInGin(ctx *gin.Context) (*string, error) {

	info, ok := ctx.Get(GinKeyMessage)
	if !ok {
		err := fmt.Errorf("GinKeyMessage is not find")
		return nil, err
	}

	msg, ok := info.(string)
	if !ok {
		err := fmt.Errorf("Trans GinKeyMessage.(string) failed from gin cache")
		return nil, err
	}

	return &msg, nil
}

func CacheMessageInGin(ctx *gin.Context, message string) {
	ctx.Set(GinKeyMessage, message)
}
