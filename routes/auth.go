package routes

import (
	"github.com/sundogrd/content-api/handler"

	"github.com/gin-gonic/gin"
)

// Auth ...
func Auth(r *gin.Engine) {
	r.GET("/auth", handler.Auth)
	r.GET("/login", handler.GithubLogin)
	r.GET("/loginCallBack", handler.GithubLoginCallBack)
}
