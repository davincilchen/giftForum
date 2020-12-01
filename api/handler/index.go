package handler

import (
	"giftForum/api/ginprocess"
	"giftForum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	//TODO:重構
	html := indexHTML
	user, _ := ginprocess.GetLoginUserInGin(ctx)

	topTxUser, err:= models.GetTopTxUser()
	if err != nil{
		if user == nil {
			ctx.HTML(http.StatusOK, html, nil)
			return
		}
		ResposnHtmlWithUser(ctx, html, &user.BaseUser)
		return 
	}

	topRxUser, err:= models.GetTopRxUser()
	if err != nil{
		if user == nil {
			ctx.HTML(http.StatusOK, html, nil)
			return
		}
		ResposnHtmlWithUser(ctx, html, &user.BaseUser)
		return
	}	

	if user == nil {
		ctx.HTML(http.StatusOK, html, gin.H{		
			GinHTopTxUsers: topTxUser,
			GinHTopRxUsers: topRxUser,
		})
		return
	}

	// topTxUser, err:= models.GetTopTxUser()
	// if err != nil{
	// 	ResposnHtmlWithUser(ctx, html, &user.BaseUser)
	// 	return 
	// }

	// topRxUser, err:= models.GetTopRxUser()
	// if err != nil{
	// 	ResposnHtmlWithUser(ctx, html, &user.BaseUser)
	// 			return
	// }	
	
	ctx.HTML(http.StatusOK, html, gin.H{		
		GinHUser:  &user.BaseUser,
		GinHTopTxUsers: topTxUser,
		GinHTopRxUsers: topRxUser,
	})
}
