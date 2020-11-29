package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type APILog struct {
	RequestURI   string
	Method       string
	Duration     time.Duration
	DurationText string
}

func Logger(ctx *gin.Context) {
	now := time.Now()

	ctx.Next()
	log := APILog{}
	log.RequestURI = ctx.Request.RequestURI
	log.Method = ctx.Request.Method
	log.Duration = time.Since(now)
	log.DurationText = fmt.Sprintf("%v", log.Duration)

	fmt.Printf("%+v\n", log)

}
