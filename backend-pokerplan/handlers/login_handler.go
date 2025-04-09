package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

var oauth2Config = &oauth2.Config{
	ClientID:     "Ov23liIn3aTcuubUuzs4",
	ClientSecret: "3bfef7fb7e9af47c29aab5425fa6fb66bd64f851",
	Endpoint:     oauth2github.Endpoint,
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"read:org", "user"},
}

func LoginHandler(c *gin.Context) {
	oauth2Config.ClientID = "Ov23liIn3aTcuubUuzs4"
	state := generateState()
	//session := sessions.Default(c)
	//session.Set("state", []byte("123"))
	//session.Save()
	url := oauth2Config.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)
}

func generateState() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
