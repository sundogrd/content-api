package content

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"strconv"
)

func GetContent(c *gin.Context) {
	contentId := c.Param("contentId")
	id, err := strconv.ParseInt(contentId, 10, 64)
	if err != nil {
		panic(err)
	}
	res := content.ContentRepositoryInstance().FindOne(content.FindOneRequest{ContentID: id})
	if res.ID == 0 {
		c.JSON(404, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(200, res)
}

// GetContent ...
func GetContent(c *gin.Context) {
	contentId := c.Param("contentId")
	res, err := getContentById(contentId)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(200, res)
}
