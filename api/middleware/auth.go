package middleware

import (
	"giftForum/api/ginprocess"
	"giftForum/models"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {

	uuid, err := ginprocess.GetUserSessionCookie(ctx)
	if err != nil {
		return
	}

	loginPlayer, err := models.GetLoginPlayer(*uuid)
	if err != nil {
		return
	}
	ginprocess.CachePlayerSessionInGin(ctx, loginPlayer)

}
