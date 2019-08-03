package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/di"
)

// Routes ...
func Routes(r *gin.Engine, container *di.Container) {
	Hello(r)
	Content(r, container)
	Auth(r)
	SDLog(r)
	Comment(r, container)
	Uploader(r)
}
