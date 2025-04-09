package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v70/github"
	"golang.org/x/oauth2"
)

func CallbackHandler(c *gin.Context) {

	code := c.Query("code")

	token, err := oauth2Config.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
		fmt.Println(err)
		fmt.Println("###############################")
		fmt.Println("###################################       ", code)
		return
	}

	githubUser, err := getGithubUser(token)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Failed to get user info"})
		return
	}

	fmt.Println("User ID:", *githubUser.Login)
	fmt.Println("User Name:", *githubUser.Name)

	c.Redirect(http.StatusFound, "http://localhost:5173/dashboard")

}

func getGithubUser(token *oauth2.Token) (*github.User, error) {
	ctx := context.Background()
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)))

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}
	return user, nil
}
