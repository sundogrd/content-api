package routes

import (
	contentHandler "github.com/sundogrd/content-api/handler/content"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Content(r *gin.Engine) {
	r.GET("/contents/:contentId", contentHandler.GetContent)
	r.GET("/contents", contentHandler.ListContent)
	r.GET("/contents/:contentId/recommends", contentHandler.GetContentRecommend)
	r.POST("/contents", contentHandler.CreateContent)
	//r.PATCH("/contents/:contentId", handler.UpdateContent)
	r.DELETE("/contents/:contentId", contentHandler.DeleteContent)
}
