package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/di"
	"github.com/sundogrd/content-api/handler/comment"
)

func Comment(r *gin.Engine, container *di.Container) {
	r.POST("/comments", comment.CreateComment(container))
	r.GET("/comments", comment.ListComment(container))
	r.GET("/subcomments", comment.ListSubComment(container))
	r.POST("/comments/like", comment.LikeComment(container))
	r.POST("/comments/hate", comment.HateComment(container))
}
