package content

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"github.com/sundogrd/content-api/services/user"
	"github.com/sundogrd/content-api/utils/pointer"
	"net/http"
	"strconv"
	"time"
)

type GetContentResponseAuthor struct {
	UserID    string             `json:"user_id"`
	Name      string             `json:"name"`
	AvatarUrl string             `json:"avatar_url"`
	Company   *string            `json:"company"`
	Email     *string            `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Extra     user.UserInfoExtra `json:"extra"`
}
type GetContentResponse struct {
	ContentID   string                   `json:"content_id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Author      GetContentResponseAuthor `json:"author"`
	Category    string                   `json:"category"`
	Type        content.ContentType      `json:"type"`
	Body        string                   `json:"body"`
	BodyType    content.BodyType         `json:"body_type"`
	Version     int16                    `json:"version"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	Extra       content.ContentInfoExtra `json:"extra"`
}

func GetContent(c *gin.Context) {
	contentId := c.Param("contentId")
	id, err := strconv.ParseInt(contentId, 10, 64)
	if err != nil {
		panic(err)
	}
	contentFindOneRes, err := content.ContentServiceInstance().FindOne(content.FindOneRequest{ContentID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": err,
		})
		return
	}
	if contentFindOneRes.ContentID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "ContentID: " + contentId + " Not Found",
		})
		return
	}
	userFindOneRes, err := user.UserServiceInstance().FindOne(user.FindOneRequest{
		UserID: pointer.Int64(contentFindOneRes.AuthorID),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, GetContentResponse{
		ContentID:   strconv.FormatInt(contentFindOneRes.ContentID, 10),
		Title:       contentFindOneRes.Title,
		Description: contentFindOneRes.Description,
		Author: GetContentResponseAuthor{
			UserID:    strconv.FormatInt(userFindOneRes.UserID, 10),
			Name:      userFindOneRes.Name,
			AvatarUrl: userFindOneRes.AvatarUrl,
			Company:   userFindOneRes.Company,
			Email:     userFindOneRes.Email,
			CreatedAt: userFindOneRes.CreatedAt,
			UpdatedAt: userFindOneRes.UpdatedAt,
			Extra:     userFindOneRes.Extra,
		},
		Category:  contentFindOneRes.Category,
		Type:      contentFindOneRes.Type,
		Body:      contentFindOneRes.Body,
		BodyType:  contentFindOneRes.BodyType,
		Version:   contentFindOneRes.Version,
		CreatedAt: contentFindOneRes.CreatedAt,
		UpdatedAt: contentFindOneRes.UpdatedAt,
		Extra:     contentFindOneRes.Extra,
	})
}
