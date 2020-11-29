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
