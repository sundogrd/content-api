package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/handler/sdlog"
)

// SDLog ...
func SDLog(r *gin.Engine) {
	r.POST("/log", sdlog.AddStatement)
	r.GET("/log", sdlog.GetStatement)
	r.GET("/log/count", sdlog.GetStatementCount)
}
