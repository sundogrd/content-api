package content

import (
	"fmt"
	"github.com/sundogrd/content-api/di"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
	"github.com/sundogrd/content-api/services/user"
	"github.com/sundogrd/content-api/utils/pointer"
)

// GetContentResponseAuthor ...
type GetContentResponseAuthor struct {
	UserID    string             `json:"user_id"`
	Name      string             `json:"name"`
	AvatarURL string             `json:"avatar_url"`
	Company   *string            `json:"company"`
	Email     *string            `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	Extra     user.BaseInfoExtra `json:"extra"`
}

// GetContentResponse ...
type GetContentResponse struct {
	ContentID   string                   `json:"content_id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Author      GetContentResponseAuthor `json:"author"`
	Category    string                   `json:"category"`
	Type        content.ContentType      `json:"type"`
	Status      content.ContentStatus    `json:"status"`
	Body        string                   `json:"body"`
	BodyType    content.BodyType         `json:"body_type"`
	Version     int16                    `json:"version"`
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
	Extra       content.FullInfoExtra    `json:"extra"`
}

// GetContent ...
func GetContent(container *di.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		contentID := c.Param("contentId")
		id, err := strconv.ParseInt(contentID, 10, 64)
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
				"msg": "ContentID: " + contentID + " Not Found",
			})
			return
		}
		_, err = content.ContentServiceInstance().Read(content.ReadRequest{
			ContentID: id,
		})
		if err != nil {
			fmt.Errorf("read %d error", id)
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
				AvatarURL: userFindOneRes.AvatarURL,
				Company:   userFindOneRes.Company,
				Email:     userFindOneRes.Email,
				CreatedAt: userFindOneRes.CreatedAt,
				UpdatedAt: userFindOneRes.UpdatedAt,
				Extra:     userFindOneRes.Extra,
			},
			Category:  contentFindOneRes.Category,
			Status:    contentFindOneRes.Status,
			Type:      contentFindOneRes.Type,
			Body:      contentFindOneRes.Body,
			BodyType:  contentFindOneRes.BodyType,
			Version:   contentFindOneRes.Version,
			CreatedAt: contentFindOneRes.CreatedAt,
			UpdatedAt: contentFindOneRes.UpdatedAt,
			Extra:     contentFindOneRes.Extra,
		})
	}
}
