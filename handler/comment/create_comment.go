package comment

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/env"
	"github.com/sundogrd/content-api/grpc_gen/comment"
	"github.com/sundogrd/content-api/middlewares/sdsession"
)

type CreateCommentRequest struct {
	ContentID   string `form:"content_id" json:"content_id"`       // 内容对象id, 内容对象我们这里指文章
	ParentID    string `form:"parent_id" json:"parent_id"`         // 父级对象id, 这里一般指主评论
	ReCommentID string `form:"re_comment_id" json:"re_comment_id"` // 回复id, 对主评论下回复的回复id
	Content     string `form:"content" json:"content"`
}
type CreateCommentResponse struct {
	CommentID string `json:"comment_id"`
}

func CreateComment(container env.Container) gin.HandlerFunc {
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

		var parentIDNum int64
		if request.ParentID != "" {
			parentIDNum, err = strconv.ParseInt(request.ParentID, 10, 64)
			if err != nil {
				logrus.Errorf("[content-api/handler/content] CreateComment parse ParentID err: %+v", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "解析ParentID出错",
				})
				return
			}
		} else {
			parentIDNum = 0
		}

		var ReCommentIDNum int64
		if request.ReCommentID != "" {
			ReCommentIDNum, err = strconv.ParseInt(request.ReCommentID, 10, 64)
			if err != nil {
				logrus.Errorf("[content-api/handler/content] CreateComment parse ReCommentID err: %+v", err)
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "解析ReCommentID出错",
				})
				return
			}
		} else {
			ReCommentIDNum = 0
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
		// authorID := int64(312337740408565760)

		var params = &comment.CreateCommentRequest_CommentCreateParams{}

		params.TargetId = contentIDNum
		params.CreatorId = authorID
		params.Content = request.Content

		if parentIDNum != 0 {
			params.ParentId = parentIDNum
		}

		if ReCommentIDNum != 0 {
			params.ReCommentId = ReCommentIDNum
		}

		res, err := container.CommentGrpcClient.CreateComment(c, &comment.CreateCommentRequest{
			AppId:   "lwio",
			Comment: params,
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
