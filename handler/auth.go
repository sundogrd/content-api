package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/sundogrd/content-api/utils/config"
	"golang.org/x/oauth2"
)

type User struct {
	Name     *string
	Bio      *string
	Location *string
	Url      *string
	Company  *string
}

var state string

const htmlIndex = `<html><body>
<a href="/login">Log in with Github</a>
</body></html>`

func getGitHubConfig(name string) string {
	prefix := "auth.github."
	return config.GetString(prefix + name)
}

var githubOauthConfig *oauth2.Config

// auth github test...
func Auth(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, htmlIndex)
}

/*github oauth2*/
func GithubLogin(c *gin.Context) {
	b := make([]byte, 16)
	rand.Read(b)

	state = base64.URLEncoding.EncodeToString(b)
	// fmt.Println(state)
	githubOauthConfig = &oauth2.Config{
		ClientID:     getGitHubConfig("client.ClientID"),
		ClientSecret: getGitHubConfig("client.ClientSecret"),
		RedirectURL:  getGitHubConfig("redirectUrl"),
		Scopes:       config.GetStringSlice("auth.github.scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  getGitHubConfig("endpoint.authorizeUrl"),
			TokenURL: getGitHubConfig("endpoint.tokenUrl"),
		},
	}
	url := githubOauthConfig.AuthCodeURL(state)
	fmt.Println(url)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// callback 获取github返回来的数据
func GithubLoginCallBack(c *gin.Context) {
	// fmt.Println("call back start")
	callbackState := c.Query("state")
	if callbackState != state {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", state, callbackState)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Println(state)
	code := c.Query("code")
	// fmt.Println(code)
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	// fmt.Println(token)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	if !token.Valid() {
		c.JSON(http.StatusTemporaryRedirect, gin.H{
			"msg": "retreived invalid token",
		})
		return
	}
	client := github.NewClient(githubOauthConfig.Client(oauth2.NoContext, token))
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		c.JSON(http.StatusTemporaryRedirect, gin.H{
			"msg": "failed to get user",
		})
		return
	}
	ret := User{
		Name:     user.Name,
		Bio:      user.Bio,
		Company:  user.Company,
		Url:      user.URL,
		Location: user.Location,
	}
	c.JSON(200, gin.H{
		"data": ret,
	})
}
