package ginprocess

import (
	"giftForum/config"

	"github.com/gin-gonic/gin"
)

func CleanUserSessionCookie(ctx *gin.Context) {
	ctx.SetCookie("userUUID", "", -1, "/", config.GetDomain(), false, true)
}

func SetUserSessionCookie(ctx *gin.Context, uuid string) {
	ctx.SetCookie("userUUID", uuid, 0, "/", config.GetDomain(), false, true)
}

func GetUserSessionCookie(ctx *gin.Context) (*string, error) {
	uuid, err := ctx.Cookie("userUUID")
	if err != nil {
		return nil, err
	}
	return &uuid, nil
}
