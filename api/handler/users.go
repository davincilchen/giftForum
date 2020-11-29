package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUsersSignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateUsersSignIn")
}

func GetUsersSignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetUsersSignUp")
}
func CreateUsersSignUp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateUsersSignUp")
}
