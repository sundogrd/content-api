package routes

import (
	"github.com/gin-gonic/gin"
)

// Routes ...
func Routes(r *gin.Engine) {
	View(r)
	Hello(r)
	Content(r)
	Auth(r)
}
