package routes

import (
	"github.com/sundogrd/content-api/handler"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Hello(r *gin.Engine) {
	r.GET("/hello", handler.Hello)
}
