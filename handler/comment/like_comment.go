package comment

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/env"
	"github.com/sundogrd/content-api/grpc_gen/comment"
)

type LikeCommentRequest struct {
	CommentID string `form:"comment_id" json:"comment_id"`
}
type LikeCommentResponse struct {
	CommentID string `form:"comment_id" json:"comment_id"`
}

// ListContent ...
// type title author category type created_at updated_at
func LikeComment(container env.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request LikeCommentRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			logrus.Errorf("[content-api/handler/content] ListComment ShouldBindQuery err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		commentIDNum, err := strconv.ParseInt(request.CommentID, 10, 64)
		if err != nil {
			logrus.Errorf("[content-api/handler/content] ListComment parse commentIDNum err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "解析CommentID出错",
			})
			return
		}

		res, err := container.CommentGrpcClient.Like(c, &comment.LikeRequest{
			CommentId: commentIDNum,
		})
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, LikeCommentResponse{
			CommentID: strconv.Itoa(int(res.CommentId)),
		})
	}
}
