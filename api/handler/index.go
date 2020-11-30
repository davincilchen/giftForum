package handler

import (
	"giftForum/api/ginprocess"
	"giftForum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	user, _ := ginprocess.GetLoginUserInGin(ctx)

	user = &models.LoginUser{} //test***test
	user.ID = 5
	user.Email = "ad%@gamil"

	if user == nil {
		ctx.HTML(http.StatusOK, indexHTML, nil)
	}

	ResposnHtmlWithUser(ctx, indexHTML, user)

}
