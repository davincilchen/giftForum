package handler

import (
	"fmt"
	"giftForum/basemodels"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResposnSuccessHtmlWithUser(ctx *gin.Context, html string, success interface{}, user *basemodels.BaseUser) {
	if success == nil {
		ResposnHtmlWithUser(ctx, html, user)
		return
	}
	if user == nil {
		ResposnSuccessHtml(ctx, html, success)
		return
	}
	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": success,
		GinHUser:  user,
	})

}

func ResposnSuccessHtml(ctx *gin.Context, html string, success interface{}) {
	if success == nil {
		ctx.HTML(http.StatusOK, html, nil)
		return
	}
	ctx.HTML(http.StatusOK, indexHTML, gin.H{
		"success": success,
	})

}
func ResposnHtmlWithUser(ctx *gin.Context, html string, user *basemodels.BaseUser) {
	fmt.Printf("user %#v \n", user)
	if user == nil {
		ctx.HTML(http.StatusOK, html, nil)
		return
	}
	ctx.HTML(http.StatusOK, html, gin.H{
		GinHUser: user,
	})
}
