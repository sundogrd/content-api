package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/content"
)

type UpdatePayload struct {
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Category    string `form:"category" json:"category"`
	Type        int16  `form:"type" json:"type"`
	Body        string `form:"body" json:"body"`
}

type FindRequestPayload struct {
	UpdatePayload
	AuthorID      string    `form:"author_id" json:"author_id"`
	ContentIDs    []int64   `form:"content_ids" json:"content_ids"`
	CreatedAtFrom time.Time `form:"created_at_from" json:"created_at_from" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	CreatedAtTo   time.Time `form:"created_at_to" json:"created_at_to" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	UpdatedAtFrom time.Time `form:"updated_at_from" json:"updated_at_from" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	UpdatedAtTo   time.Time `form:"updated_at_to" json:"updated_at_to" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	DeletedAtFrom time.Time `form:"deleted_at_from" json:"deleted_at_from" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	DeletedAtTo   time.Time `form:"deleted_at_to" json:"deleted_at_to" time_format:"2006-01-02 15:04:05" time_utc:"1"`
	Page          int16     `form:"page" json:"page"`
	PageSize      int16     `form:"page_size" json:"page_size"`
}

func getContentById(sId string) (*content.FindOneResponse, error) {
	ctx := context.Background()
	id, err := strconv.ParseInt(sId, 10, 64)
	if err != nil {
		panic(err)
	}
	res := content.ContentRepositoryInstance().FindOne(ctx, content.FindOneRequest{ContentID: id})
	if res.ID == 0 {
		return &content.FindOneResponse{}, errors.New("Not Found")
	}
	return res, nil
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
	c.JSON(200, gin.H{
		"msg":  "successfully",
		"data": res,
	})
}

func getLists(c *gin.Context) (*content.FindResponse, error) {
	ctx := context.Background()
	var request FindRequestPayload
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, errors.New("Parameter parse error!")
	}
	req := &content.FindRequest{}
	// 去TM的反射动态赋值
	// elem := reflect.ValueOf(req).Elem()

	// t := reflect.TypeOf(request)
	// v := reflect.ValueOf(request)

	// for k := 0; k < t.NumField(); k++ {
	// 	elem.FieldByName(t.Field(k).Name).Set(v.Field(k).Interface())
	// }

	// 好恶心啊
	if request.Title != "" {
		req.Title = &(request.Title)
	}
	if request.AuthorID != "" {
		req.AuthorID = &(request.AuthorID)
	}
	if request.Description != "" {
		req.Description = &(request.Description)
	}
	if request.Type != 0 {
		req.Type = &(request.Type)
	}
	if request.ContentIDs != nil {
		req.ContentIDs = &(request.ContentIDs)
	}
	if request.Category != "" {
		req.Category = &(request.Category)
	}
	if !request.CreatedAtFrom.IsZero() {
		req.CreatedAtFrom = &(request.CreatedAtFrom)
	}
	if !request.CreatedAtTo.IsZero() {
		req.CreatedAtTo = &(request.CreatedAtTo)
	}
	if !request.UpdatedAtFrom.IsZero() {
		req.UpdatedAtFrom = &(request.UpdatedAtFrom)
	}
	if !request.UpdatedAtTo.IsZero() {
		req.UpdatedAtTo = &(request.UpdatedAtTo)
	}
	if !request.DeletedAtFrom.IsZero() {
		req.DeletedAtFrom = &(request.DeletedAtFrom)
	}
	if !request.DeletedAtTo.IsZero() {
		req.DeletedAtTo = &(request.DeletedAtTo)
	}
	if request.Page != 0 {
		req.Page = &(request.Page)
	}
	if request.PageSize != 0 {
		req.PageSize = &(request.PageSize)
	}
	res, err := content.ContentRepositoryInstance().Find(ctx, *req)
	if err != nil {
		log.Fatalln(err)
	}
	return res, nil
}

// ListContent ...
// type title author category type created_at updated_at
func ListContent(c *gin.Context) {
	res, err := getLists(c)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"msg":  "Get Successfully",
		"data": res,
	})
}

func CreateContent(c *gin.Context) {
	ctx := context.Background()
	var request content.CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := content.ContentRepositoryInstance().Create(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(err)
	c.JSON(200, gin.H{
		"data": res,
	})
}

func UpdateContent(c *gin.Context) {
	var payload UpdatePayload
	conId := c.Param("contentId")
	con, err := getContentById(conId)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": err,
		})
		return
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request := content.UpdateRequest{
		Target:      content.PFContent{ContentID: con.ContentID},
		Title:       payload.Title,
		Description: payload.Description,
		Category:    payload.Category,
		Type:        payload.Type,
		Body:        payload.Body,
	}
	res, err := content.ContentRepositoryInstance().Update(context.Background(), request)
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "Internal Error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "successful updated!",
		"data": res,
	})
}

func DeleteContentById(c *gin.Context) {
	conId := c.Param("contentId")
	ctx := context.Background()
	id, err := strconv.ParseInt(conId, 10, 64)
	if err != nil {
		panic(err)
	}
	ids := []int64{id}
	_, err = content.ContentRepositoryInstance().Delete(ctx, content.DeleteRequest{
		ContentIDs: ids,
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
func DeleteContents(c *gin.Context) {
	ctx := context.Background()

	res, err := getLists(c)
	if err != nil {
		return
	}
	var ids []int64
	for _, con := range res.List {
		ids = append(ids, con.ContentID)
	}
	if int64(len(ids)) != res.Total {
		c.JSON(204, gin.H{
			"msg": "no content", // 用个啥http码呢？
		})
		return
	}
	_, err = content.ContentRepositoryInstance().Delete(ctx, content.DeleteRequest{
		ContentIDs: ids,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"msg": "Delete Failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "Successfully deleted",
		"data": ids,
	})
}
