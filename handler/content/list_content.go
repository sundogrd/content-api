package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/di"
	contentGrpc "github.com/sundogrd/content-api/grpc_gen/content"
	"github.com/sundogrd/content-api/services/content"
	"net/http"
)

type ListContentRequest struct {
	Title      *string                `form:"title" json:"title"`
	Type       *content.ContentType   `form:"type" json:"type"`
	AuthorID   *int64                 `form:"author_id" json:"author_id"`
	Status     *content.ContentStatus `form:"status" json:"status"`
	ContentIDs []int64                `form:"content_ids" json:"content_ids"`
	Page       *int32                 `form:"page" json:"page"`
	PageSize   *int32                 `form:"page_size" json:"page_size"`
}
type ListContentResponse struct {
	List  []*contentGrpc.Content `json:"list"`
	Total int64              `json:"total"`
}

func ListContent(container *di.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ListContentRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 500126,
				"msg": err.Error(),
			})
			return
		}
		page := int32(1)
		pageSize := int32(10)
		if req.Page != nil {
			page = *req.Page
		}
		if req.PageSize != nil {
			pageSize = *req.PageSize
		}


		listRes, err := container.ContentGrpcClient.ListContents(c, &contentGrpc.ListContentsRequest{
			AppId: "lwio",
			Page: page,
			PageSize: pageSize,
			State: contentGrpc.EContentState_PUBLISHED,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500126,
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, ListContentResponse{
			List:  listRes.Contents,
			Total: listRes.Total,
		})
		return
	}
}