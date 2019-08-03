package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/di"
	"net/http"
)

// DeleteContent ...
func DeleteContent(container *di.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "not implemented",
		})
	}
}

