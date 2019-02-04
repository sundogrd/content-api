package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"net/http"
	"strconv"
)

type GetContentResponse struct {
	content.ContentInfo
}

func GetContent(c *gin.Context) {
	contentId := c.Param("contentId")
	id, err := strconv.ParseInt(contentId, 10, 64)
	if err != nil {
		panic(err)
	}
	res, err := content.ContentServiceInstance().FindOne(content.FindOneRequest{ContentID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err,
		})
		return
	}
	if res.ContentID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "ContentID: " + contentId + " Not Found",
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
