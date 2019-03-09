package content

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/middlewares/sdsession"
	"github.com/sundogrd/content-api/services/content"
)

type CreateContentRequest struct {
	Title    string                 `json:"title"`
	Type     content.ContentType    `json:"type"`
	Status   *content.ContentStatus `json:"status"`
	Body     string                 `json:"body"`
	BodyType content.BodyType       `json:"body_type"`
}
type CreateContentResponse struct {
	content.BaseInfo
	ContentID string `json:"content_id"`
	AuthorID  string `json:"author_id"`
}

func CreateContent(c *gin.Context) {
	var request CreateContentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}
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
	res, err := content.ContentServiceInstance().Create(content.CreateRequest{
		Title:       request.Title,
		Description: "",
		AuthorID:    authorID,
		Category:    "",
		Status:      request.Status,
		Type:        request.Type,
		Body:        request.Body,
		Extra:       content.BaseInfoExtra{},
	})
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CreateContentResponse{
		BaseInfo:  res.BaseInfo,
		ContentID: strconv.FormatInt(res.ContentID, 10),
		AuthorID:  strconv.FormatInt(res.AuthorID, 10),
	})
}
