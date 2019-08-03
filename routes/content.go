package routes

import (
	"github.com/sundogrd/content-api/di"
	contentHandler "github.com/sundogrd/content-api/handler/content"

	"github.com/gin-gonic/gin"
)

// Hello ...
func Content(r *gin.Engine, container *di.Container) {
	r.GET("/contents/:contentId", contentHandler.GetContent(container))
	r.GET("/contents", contentHandler.ListContent(container))
	r.GET("/contents/:contentId/recommends", contentHandler.GetContentRecommend(container))
	r.POST("/contents", contentHandler.CreateContent(container))
	//r.PATCH("/contents/:contentId", handler.UpdateContent)
	r.DELETE("/contents/:contentId", contentHandler.DeleteContent(container))
}
