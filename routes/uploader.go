package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/handler/uploader"
)

// Uploader ...
func Uploader(r *gin.Engine) {
	r.GET("/qiniu/token", uploader.GetQiniuToken)
}
