package handler

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
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

func gAuth() {

	//key := "Secret-session-key"  // Replace with your SESSION_SECRET or similar

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New("our-google-client-id",
			"our-google-client-secret",
			"http://localhost:8081/auth/google/callback", "email", "profile"),
	)

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		fmt.Printf("***1 user %#v\n", user)
		// t, _ := template.ParseFiles("templates/success.html")
		// t.Execute(res, user)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		// t, _ := template.ParseFiles("templates/index.html")
		// t.Execute(res, false)
		fmt.Printf("*** no *** \n")
	})

	log.Println("listening on localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", p))
}

func GoogleAuth(ctx *gin.Context) {
	gAuth()
	// config := CreateGoogleConfig()
	// client, err := oauth2ns.AuthenticateUser(config)
	// if err != nil {
	// 	fmt.Println("oauth2ns.AuthenticateUser error : ", err)
	// }

	user, err := FetchUserInfo(client)
	if err != nil {
		fmt.Println("oauth2ns.AuthenticateUser error : ", err)
	}
	fmt.Printf("user %#v", user)
	//spew.Dump(client)
	ctx.HTML(http.StatusOK, loginHTML, nil)
}

var client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          1000,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, //取消驗證憑證
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

type UserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

func FetchUserInfo(client *http.Client) (*UserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result UserInfo
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return &result, nil
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
