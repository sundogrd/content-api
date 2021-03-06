package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/env"
	"github.com/sundogrd/content-api/handler/comment"
)

func Comment(r *gin.Engine, container env.Container) {
	r.POST("/comments", comment.CreateComment(container))
	r.GET("/comments", comment.ListComment(container))
	r.GET("/subcomments", comment.ListSubComment(container))
	r.POST("/comments/like", comment.LikeComment(container))
	r.POST("/comments/hate", comment.HateComment(container))
}
