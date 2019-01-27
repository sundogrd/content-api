package routes

import (
	"github.com/gin-gonic/gin"
)

// Routes ...
func Routes(r *gin.Engine) {
	Hello(r)
	Content(r)
	Auth(r)
}
