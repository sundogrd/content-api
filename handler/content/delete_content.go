package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"strconv"
)

func DeleteContent(c *gin.Context) {
	contentID, err := strconv.ParseInt(c.Param("contentId"), 10, 64)
	if err != nil {
		panic(err)
	}
	_, err = content.ContentServiceInstance().Delete(content.DeleteRequest{
		ContentID: contentID,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "Delete Failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "Deleted successfully",
		"data": nil,
	})
}
