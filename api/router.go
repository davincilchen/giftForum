package apiv

import (
	"giftForum/api/handler"
	"giftForum/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.LoadHTMLGlob("template/html/*")
	router.Static("/assets", "./template/assets")

	// .. //
	router.Use(middleware.Logger,
		gin.Recovery(),
		middleware.Auth)
	router.GET("/", handler.Index)
	router.GET("/index", handler.Index)

	router.GET("users/sign_in", handler.GetUsersSignIn)
	router.POST("users/sign_in", handler.CreateUsersSignIn)
	router.DELETE("users/:sign_out", handler.DeleteUsersSignOut)
	router.GET("users/sign_up", handler.GetUsersSignUp)
	router.POST("users/sign_up", handler.CreateUsersSignUp)
	router.GET("user/:id", handler.GetUser)

	return router
}
