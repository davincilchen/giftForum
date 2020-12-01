package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"


)

// func CreateGoogleOAuthURL() string {
// 	// 使用 lib 產生一個特定 config instance
// 	config := &oauth2.Config{
// 		//憑證的 client_id
// 		ClientID:
// 		//憑證的 client_secret
// 		ClientSecret: ,
// 		//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
// 		// RedirectURL:  "http://localhost:8080/google-login-callback",
// 		RedirectURL: "http://localhost:8081/",
// 		//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
// 		Scopes: []string{
// 			"https://www.googleapis.com/auth/userinfo.email",
// 			"https://www.googleapis.com/auth/userinfo.profile",
// 		},
// 		//指的是 Google auth server 的 endpoint，用 lib 預設值即可
// 		Endpoint: google.Endpoint,
// 	}

// 	//產生出 config instance 後，就可以使用 func AuthCodeURL 建立請求網址
// 	return config.AuthCodeURL("state")
// }

func CreateGoogleConfig() *oauth2.Config {
	// 使用 lib 產生一個特定 config instance
	config := &oauth2.Config{
		//憑證的 client_id

		//憑證的 client_secret

		//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
		// RedirectURL:  "http://localhost:8080/google-login-callback",
		RedirectURL: "http://localhost:8081/",
		//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		//指的是 Google auth server 的 endpoint，用 lib 預設值即可
		Endpoint: google.Endpoint,
	}

	return config
}


func GoogleAuth(ctx *gin.Context) {

	// config := CreateGoogleConfig()
	// client, err := oauth2ns.AuthenticateUser(config)
	// if err != nil {
	// 	fmt.Println("oauth2ns.AuthenticateUser error : ", err)
	// }

	//fmt.Printf("user %#v", user)
	//spew.Dump(client)

	ctx.HTML(http.StatusOK, loginHTML, nil)
}




// func CreateUsersSignIn(ctx *gin.Context) {
// 	var err error
// 	code := http.StatusBadRequest

// 	defer func() {
// 		if err == nil {
// 			return
// 		}
// 		ctx.HTML(code, loginHTML, gin.H{
// 			"error": err,
// 		})
// 	}()

// 	user, err := GetUserForm(ctx)
// 	if err != nil {
// 		return
// 	}

// 	loginUser, err := models.UserLogin(user.Email, user.Password)
// 	if err != nil {
// 		code = http.StatusUnauthorized
// 		err = errors.New("email不存在或password錯誤")
// 		return
// 	}

// 	ginprocess.SetUserSessionCookie(ctx, loginUser.UUID)
// 	ctx.Redirect(http.StatusFound, "/")
// 	// ctx.HTML(http.StatusOK, indexHTML, gin.H{
// 	// 	"success": "登入成功",
// 	// })
// 	return
// }
