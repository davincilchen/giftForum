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
		middleware.AuthLogin)
	router.GET("/", handler.Index)
	router.GET("/index", handler.Index)

	router.POST("users/google_oauth", handler.GoogleAuth)
	router.GET("users/sign_in", handler.GetUsersSignIn)
	router.POST("users/sign_in", handler.CreateUsersSignIn)
	router.POST("users/sign_out", handler.CreateUsersSignOut)
	router.GET("users/sign_up", handler.GetUsersSignUp)
	router.POST("users/sign_up", handler.CreateUsersSignUp)
	router.GET("user/:id", handler.GetUser)
	router.POST("/user/:from_id/to/:id/gift", handler.CreateUserSendGift)
	router.GET("/callback", handler.HandleGoogleCallback)

	
	return router
}
