package apiv

import (
	"giftForum/api/handler"
	"giftForum/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(middleware.Logger, 
		gin.Recovery(), 
		middleware.Auth)
	router.GET("/index", handler.Index)

	
	
	return router
}
