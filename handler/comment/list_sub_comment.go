package comment

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/content-api/di"
	"github.com/sundogrd/content-api/grpc_gen/comment"
)

type ListSubCommentRequest struct {
	ContentID string `form:"content_id" json:"content_id"`
	TargetId  string `form:"target_id" json:"target_id"`
	Page      *int32 `form:"page" json:"page"`
	PageSize  *int32 `form:"page_size" json:"page_size"`
}
type ListSubCommentResponse struct {
	List  []*comment.Comment `json:"list"`
	Total int64              `json:"total"`
}

// ListContent ...
// type title author category type created_at updated_at
func ListSubComment(container *di.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request ListSubCommentRequest
		if err := c.ShouldBindQuery(&request); err != nil {
			logrus.Errorf("[content-api/handler/content] ListComment ShouldBindQuery err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}
		contentIDNum, err := strconv.ParseInt(request.ContentID, 10, 64)
		if err != nil {
			logrus.Errorf("[content-api/handler/content] ListComment parse contentIDNum err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "解析ContentID出错",
			})
			return
		}

		targetIDNum, err := strconv.ParseInt(request.TargetId, 10, 64)
		if err != nil {
			logrus.Errorf("[content-api/handler/content] ListComment parse targetIDNum err: %+v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "解析TargetID出错",
			})
			return
		}

		listReq := &comment.ListCommentsRequest{
			Page:     1,
			PageSize: 10,
			TargetId: contentIDNum,
			ParentId: targetIDNum,
			AppId:    "lwio",
		}

		if request.Page != nil {
			listReq.Page = *request.Page
		}
		if request.PageSize != nil {
			listReq.PageSize = *request.PageSize
		}

		res, err := container.CommentGrpcClient.ListComments(c, listReq)
		if err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
		}
		if res.Total == 0 {
			c.JSON(http.StatusOK, ListSubCommentResponse{
				List:  []*comment.Comment{},
				Total: res.Total,
			})
			return
		}
		c.JSON(http.StatusOK, ListSubCommentResponse{
			List:  res.Comments,
			Total: res.Total,
		})
	}
}
