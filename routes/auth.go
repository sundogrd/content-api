package routes

import (
	"github.com/sundogrd/content-api/handler"

	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(r *gin.Engine) {
	r.GET("/oauth2/github/auth", handler.Auth)
	r.GET("/oauth2/github/login", handler.GithubLogin)
	r.GET("/oauth2/github/callback", handler.GithubLoginCallBack)
	r.GET("/sessions/test", handler.SessionTest)
}
