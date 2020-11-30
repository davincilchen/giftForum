package handler

import (
	"giftForum/api/ginprocess"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	user, _ := ginprocess.GetLoginUserInGin(ctx)

	// user = &models.LoginUser{} //test***test //wait remove
	// user.ID = 5
	// user.Email = "ad%@gamil"

	if user == nil {
		ctx.HTML(http.StatusOK, indexHTML, nil)
		return
	}

	ResposnHtmlWithUser(ctx, indexHTML, &user.BaseUser)

}
