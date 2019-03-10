package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/handler/sdlog"
)

// SDLog ...
func SDLog(r *gin.Engine) {
	r.POST("/statement", sdlog.AddStatement)
	r.GET("/statement", sdlog.GetStatement)
}
