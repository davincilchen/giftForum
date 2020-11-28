package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func Auth(ctx *gin.Context) {
	fmt.Print("auth---\n")
}

