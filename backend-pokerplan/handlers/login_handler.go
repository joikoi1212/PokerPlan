package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

var oauth2Config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "7d7af4243f6ebd86c80e56d38b2d08f7d0a91d5f",
	Endpoint:     oauth2github.Endpoint,
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"read:org", "user"},
}

func LoginHandler(c *gin.Context) {
	fmt.Printf("Request: %+v\n", c.Request)
	oauth2config.ClientID = c.Query("client_id")
	state := generateState()
	session := sessions.Default(c)
	session.Set("state", state)
	url := oauth2Config.AuthCodeURL(state)
	c.Redirect(http.StatusFound, url)<
}

func generateState() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
