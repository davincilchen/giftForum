package handler

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/json"
	"giftForum/config"
	"giftForum/models"
	"giftForum/api/ginprocess"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"


)


type GoogleUser struct {
    ID string `json:"id"`
    Name string `json:"name"`
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Profile string `json:"profile"`
    Picture string `json:"picture"`
    Email string `json:"email"`
    EmailVerified bool `json:"verified_email"`
    Gender string `json:"gender"`
}
 

func GoogleAuth(ctx *gin.Context) {


	googleOauthConfig := config.GetGoogleOauth2Config()

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}





func HandleGoogleCallback(ctx *gin.Context) {
	
	content, err := getUserInfo(ctx.Request.URL.Query().Get("state"), ctx.Request.URL.Query().Get("code"))
	//content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	u := GoogleUser{}
	err = json.Unmarshal(content, &u)
	if err != nil {
		fmt.Println("HandleGoogleCallback GoogleUser Unmarshal error",err)
		return 
	}

	loginUser, err := models.OauthUserDone(u.Email)
	if err != nil {
		fmt.Println(err.Error())
		ctx.Redirect(http.StatusFound, "/") //TODO:
		return
	}
	
	// fmt.Printf("Content: %#v\n", content)
	// fmt.Printf("Content: %#v\n",string( content))
	// fmt.Printf("Content: %#v\n", u)
	ginprocess.SetUserSessionCookie(ctx, loginUser.UUID)
	ctx.Redirect(http.StatusFound, "/")
}





func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	googleOauthConfig := config.GetGoogleOauth2Config()//2020aaa
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
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
