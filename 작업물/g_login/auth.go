package main

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const (
	CallBackURL = "http://localhost:1333/auth/callback"

	UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
	ScopeEmail          = "https://www.googleapis.com/auth/userinfo.email"
	ScopeProfile        = "https://www.googleapis.com/auth/userinfo.profile"
)

var OAuthConf *oauth2.Config

func init() {
	OAuthConf = &oauth2.Config{
		ClientID:     "820563016223-ctjghenakuo1a9rejqqk4cketenoaeud.apps.googleusercontent.com",
		ClientSecret: "HWXApe1bqUY2eVjVB9_TJ0up",
		RedirectURL:  CallBackURL,
		Scopes:       []string{ScopeEmail, ScopeProfile},
		Endpoint:     google.Endpoint,
	}
}

func GetLoginURL(state string) string {
	return OAuthConf.AuthCodeURL(state)
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
