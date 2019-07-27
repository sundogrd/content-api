package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/env"
)

// Routes ...
func Routes(r *gin.Engine, container di.Container) {
	Hello(r)
	Content(r)
	Auth(r)
	SDLog(r)
	Comment(r, container)
	Uploader(r)
}
