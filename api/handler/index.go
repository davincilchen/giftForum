package handler

import (
	"giftForum/api/ginprocess"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	html := indexHTML
	user, _ := ginprocess.GetLoginUserInGin(ctx)

	if user == nil {
		ctx.HTML(http.StatusOK, html, nil)
		return
	}

	ResposnHtmlWithUser(ctx, html, &user.BaseUser)

	//ResposnSuccessHtmlWithUser(ctx, html, msg, &user.BaseUser)

}
