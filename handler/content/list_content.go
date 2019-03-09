package content

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"log"
	"net/http"
)

type ListContentRequest struct {
	Title      *string                `form:"title" json:"title"`
	Type       *content.ContentType   `form:"type" json:"type"`
	AuthorID   *int64                 `form:"author_id" json:"author_id"`
	Status     *content.ContentStatus `form:"status" json:"status"`
	ContentIDs []int64                `form:"content_ids" json:"content_ids"`
	Page       *int16                 `form:"page" json:"page"`
	PageSize   *int16                 `form:"page_size" json:"page_size"`
}
type ListContentResponse struct {
	List  []content.BaseInfo `json:"list"`
	Total int64              `json:"total"`
}

// ListContent ...
// type title author category type created_at updated_at
func ListContent(c *gin.Context) {
	var request ListContentRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		fmt.Errorf("[handler/content] ListContent ShouldBindQuery err: %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	findReq := content.FindRequest{}

	if request.Title != nil {
		findReq.Title = request.Title
	}
	if request.Status != nil {
		findReq.Status = request.Status
	}
	if request.AuthorID != nil {
		findReq.AuthorID = request.AuthorID
	}
	if request.Type != nil {
		findReq.Type = request.Type
	}
	if request.ContentIDs != nil {
		findReq.ContentIDs = &(request.ContentIDs)
	}
	if request.Page != nil {
		findReq.Page = request.Page
	}
	if request.PageSize != nil {
		findReq.PageSize = request.PageSize
	}
	res, err := content.ContentServiceInstance().Find(findReq)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, ListContentResponse{
		List:  res.List,
		Total: res.Total,
	})
}
