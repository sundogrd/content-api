package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"net/http"
	"strconv"
)

type GetContentRecommendUri struct {
	ContentID int64 `uri:"contentId" binding:"required,uuid"`
}
type GetContentRecommendAPIResponse struct {
	List []content.ContentInfo `json:"list"`
	//Total int64                 `json:"total"`
}

func GetContentRecommend(c *gin.Context) {
	contentID, err := strconv.ParseInt(c.Param("contentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err,
		})
		return
	}
	res, err := content.ContentServiceInstance().GetRecommendByContent(content.GetRecommendByContentRequest{
		ContentID: contentID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "successfully",
		"data": &GetContentRecommendAPIResponse{
			List: res.ContentList,
		},
	})
}
