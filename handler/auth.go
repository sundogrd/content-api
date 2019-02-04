package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/sundogrd/content-api/middlewares/sdsession"
	sdUserService "github.com/sundogrd/content-api/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.String(http.StatusOK, htmlIndex)
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
	sess := sdsession.GetSession(c)

	// fmt.Println("call back start")
	callbackState := c.Query("state")
	if callbackState != state {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", state, callbackState)
		c.Redirect(http.StatusTemporaryRedirect, "http://lwio.sundogrd.com")
		return
	}
	code := c.Query("code")
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "http://lwio.sundogrd.com")
		return
	}
	if !token.Valid() {
		c.JSON(http.StatusTemporaryRedirect, gin.H{
			"msg": "retreived invalid github token",
		})
		return
	}
	client := github.NewClient(githubOauthConfig.Client(oauth2.NoContext, token))
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		c.JSON(http.StatusTemporaryRedirect, gin.H{
			"msg": "failed to get github user",
			"err": err,
		})
		return
	}

	var userDataInfo sdUserService.DataInfo

	findOneRes := sdUserService.UserServiceInstance().FindOne(c, sdUserService.FindOneRequest{
		Name: user.Name,
	})
	fmt.Printf("findOneRes: %+v", findOneRes)
	if findOneRes == nil {
		createRes, err := sdUserService.UserServiceInstance().Create(c, sdUserService.CreateRequest{
			Name:      *user.Name,
			AvatarUrl: *user.AvatarURL,
			Company:   user.Company,
			Email:     user.Email,
			Extra: sdUserService.DataInfoExtra{
				GithubHome: *user.HTMLURL,
			},
		})
		if err != nil {
			c.JSON(500, gin.H{
				"msg": err,
			})
			c.Abort()
		} else {
			userDataInfo = sdUserService.DataInfo{
				UserID: createRes.UserID,
				Name:   createRes.Name,
			}
		}
	} else {
		userDataInfo = sdUserService.DataInfo{
			UserID: findOneRes.UserID,
			Name:   findOneRes.Name,
		}
	}
	sess.Set("user_id", userDataInfo.UserID)
	sess.Set("user_name", userDataInfo.Name)
	c.Redirect(http.StatusTemporaryRedirect, "http://lwio.sundogrd.com")
	return
}

func SessionTest(c *gin.Context) {
	sess := sdsession.GetSession(c)
	c.JSON(200, gin.H{
		"name": sess.Get("user_name"),
		"id":   sess.Get("user_id"),
	})
}

func I(c *gin.Context) {
	sess := sdsession.GetSession(c)
	c.JSON(200, gin.H{
		"name": sess.Get("user_name"),
		"id":   sess.Get("user_id"),
	})
}
