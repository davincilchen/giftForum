package config

import (
	"fmt"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"csecret"`
}

var cred Credentials
type GConfig struct {
	Domain string
}
type Config struct {
	GConfig GConfig `json:"general" env:"general"`
}

var gConfig GConfig
var gPort   string

func Initialize(buf []byte) error {

	c := Config{}
	err := json.Unmarshal(buf, &c)
	if err != nil {
		return err
	}

	gConfig = c.GConfig

	return nil
}

func Uninitialize() error {

	return nil

}

func GetDomain() string {
	return gConfig.Domain
}

func SetCredentials(c Credentials)  {
	cred = c
}

func GetCredentials() Credentials {
	return cred
}

func SetPort(c string)  {
	gPort = c
}

func GetPort() string {
	return gPort
}


func GetGoogleOauth2Config() *oauth2.Config {
	RedirectURL := fmt.Sprintf("http://localhost%s/callback",GetPort())
	googleOauthConfig := &oauth2.Config{
		RedirectURL: RedirectURL,
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return googleOauthConfig
}