package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {

	fmt.Println("index ---\n")
	ctx.HTML(http.StatusOK, loginHTML, nil)
}
