package comment

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/env"
	"github.com/sundogrd/content-api/grpc_gen/comment"
	"github.com/sundogrd/content-api/middlewares/sdsession"
	"log"
	"net/http"
	"strconv"
)

type CreateCommentRequest struct {
	ContentID string `form:"content_id" json:"content_id"`
	Content   string `form:"content" json:"content"`
}
type CreateCommentResponse struct {
	CommentID string `json:"comment_id"`
}

func CreateComment(container di.Container) gin.HandlerFunc {
	logrus.Info("-1")
	return func(c *gin.Context) {
		logrus.Info("0")
		var request CreateCommentRequest
		logrus.Info("1")
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"msg": err.Error(),
			})
			return
		}
		logrus.Info("2")
		contentIDNum, err := strconv.ParseInt(request.ContentID, 10, 64)
		if err != nil {
			logrus.Errorf("[content-api/handler/content] CreateComment parse contentIDNum err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "解析ContentID出错",
			})
			return
		}
		logrus.Info("3")
		sess := sdsession.GetSession(c)
		if sess.Get("user_id") == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "user not login",
			})
			return
		}

		authorID, err := sess.Get("user_id").(json.Number).Int64()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		//authorID := int64(312337740408565760)
		res, err := container.CommentGrpcClient.CreateComment(c, &comment.CreateCommentRequest{
			AppId: "lwio",
			Comment: &comment.CreateCommentRequest_CommentCreateParams{
				TargetId:  contentIDNum,
				CreatorId: authorID,
				Content:   request.Content,
			},
		})
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, CreateCommentResponse{
			CommentID: strconv.FormatInt(res.Comment.CommentId, 10),
		})
		return
	}
}
