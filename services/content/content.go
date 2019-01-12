package content

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	concurrentUtils "github.com/sundogrd/content-api/utils/concurrent"

	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
)

type ContentRepositoryInterface interface {
	// Put adds a new Greeting to the datastore
	FindOne(c context.Context, req FindOneRequest) *FindOneResponse
	Find(c context.Context, req FindRequest) *FindResponse
	Create(c context.Context, req CreateRequest) (*CreateResponse, error)
	Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error)
	Update(c context.Context, req UpdateRequest) (*UpdateResponse, error)
	GetRecommendByContent(c context.Context, req GetRecommendByContentRequest) (*GetRecommendByContentResponse, error)
}

type ContentRepository struct {
	db        *gorm.DB
	idBuilder *goSnowFlake.IdWorker
	ContentRepositoryInterface
}

func newContentRepository(db *gorm.DB) *ContentRepository {
	idBuilder, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Printf("[services/content] Init snowFlake id_builder error: %+v", err)
		os.Exit(1)
	}
	return &ContentRepository{
		db:        db,
		idBuilder: idBuilder,
	}
}

// FindOneRequest ...
type FindOneRequest struct {
	ID        int64
	ContentID int64
}
// FindOneResponse ...
type FindOneResponse struct {
	DataInfo
}
// FindOne ...
func (cr ContentRepository) FindOne(c context.Context, req FindOneRequest) *FindOneResponse {
	var content SDContent
	cr.db.Where(&SDContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	}).First(&content)

	res := &FindOneResponse{
		sdContentToData(content),
	}
	return res
}

type FindCreatedAtTimeRangeRequest struct {
	CreatedAtFrom *time.Time // created_at开始时间，可选
	CreatedAtTo   *time.Time // created_at结束时间，可选
}
type FindUpdatedAtTimeRangeRequest struct {
	UpdatedAtFrom *time.Time // updated_at开始时间，可选
	UpdatedAtTo   *time.Time // updated_at结束时间，可选
}
type FindDeletedAtTimeRangeRequest struct {
	DeletedAtFrom *time.Time // updated_at开始时间，可选
	DeletedAtTo   *time.Time // updated_at结束时间，可选
}

type FindContentRequest struct {
	Title       *string
	AuthorID    *string
	Description *string
	Type        *int16 // content类型
}
type FindBaseRequest struct {
	ContentIDs *[]int64
	Category   *string
}
type FindPaginationRequest struct {
	Page     *int16 // 页数从1开始，默认1
	PageSize *int16 // 页项，默认10
}

// FindRequest ...
type FindRequest struct {
	FindBaseRequest
	FindContentRequest
	FindCreatedAtTimeRangeRequest
	FindUpdatedAtTimeRangeRequest
	FindDeletedAtTimeRangeRequest
	FindPaginationRequest
}

// FindResponse ...
type FindResponse struct {
	List  []DataInfo
	Total int64
}

// Find ...
func (cr ContentRepository) Find(c context.Context, req FindRequest) (*FindResponse, error) {
	// var total int64 = 0
	contents := make([]SDContent, 0)

	db := cr.db
	// 康康参数
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	}

	if req.ContentIDs != nil && len(*req.ContentIDs) != 0 {
		db = db.Where("content_id in (?)", *req.ContentIDs)
	}
	// TODO: 搜索引擎单开一个方法，这个主要用于获取集合，所以只提供字段的查询
	if req.Title != nil {
		db = db.Where("title LIKE ?", "%"+*req.Title+"%")
	}
	if req.Description != nil {
		db = db.Where("description LIKE ?", "%"+*req.Description+"%")
	}
	if req.AuthorID != nil {
		db = db.Where("author_id = ", *req.AuthorID)
	}
	if req.Category != nil {
		db = db.Where("category = ?", *req.Category)
	}
	if req.CreatedAtFrom != nil {
		db = db.Where("created_at > ?", *req.CreatedAtFrom)
	}
	if req.CreatedAtTo != nil {
		db = db.Where("created_at < ?", *req.CreatedAtTo)
	}
	if req.UpdatedAtFrom != nil {
		db = db.Where("updated_at > ?", *req.UpdatedAtFrom)
	}

	if req.UpdatedAtTo != nil {
		db = db.Where("updated_at < ?", *req.UpdatedAtTo)
	}
	if req.DeletedAtFrom != nil {
		db = db.Where("deleted_at > ?", *req.DeletedAtFrom)
	}
	if req.DeletedAtTo != nil {
		db = db.Where("deleted_at < ?", *req.DeletedAtTo)
	}

	chanLen := 1
	errorChan := make(chan error, chanLen)

	//db.Count(&total)
	go func(db *gorm.DB) {
		if req.Page != nil && *req.Page > 0 && req.PageSize != nil && *req.PageSize > 0 {
			db.Limit(*req.PageSize).Offset((*req.Page - 1) * (*req.PageSize))
		} else {
			db.Limit(10).Offset(0)
		}
		if err := db.Find(&contents).Error; err != nil {
			errorChan <- err
		} else {
			errorChan <- nil
		}
	}(db)
	//go func(db *gorm.DB) {
	//	if err := db.Find(&SDContent{}).Count(&total).Error; err != nil {
	//		errorChan <- err
	//	}
	//}(db)
	err := concurrentUtils.WaitForError(errorChan, chanLen)

	if err != nil {
		return nil, err
	}

	res := &FindResponse{
		List:  sdContentsToDatas(contents),
		Total: int64(len(contents)),
	}
	return res, nil
}

// CreateRequest ...
type CreateRequest struct {
	Title       string
	Description string
	AuthorID    int64
	Category    string
	Type        int16
	Body        string
	Version     int16
	Extra       DataInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	DataInfo
}

// Create ...
func (cr ContentRepository) Create(c context.Context, req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	contentExtraStr, err := marshalContentExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: json marshal error: %+v", err)
		contentExtraStr, _ = marshalContentExtraJson(&DataInfoExtra{})
	}

	contentId, _ := cr.idBuilder.NextId()
	// TODO: contentType自动解析赋值
	content := SDContent{
		ContentID:   contentId,
		Title:       req.Title,
		Description: req.Description,
		AuthorID:    req.AuthorID,
		Category:    req.Category,
		Type:        1, // 先写死只有图文
		Body:        req.Body,
		BodyType:    3, // 先写死为Markdown
		Version:     req.Version,
		Extra:       contentExtraStr,
	}
	if dbc := cr.db.Create(&content); dbc.Error != nil {
		fmt.Printf("[services/content] Create: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	responseExtra, err := UnmarshalContentExtraJson(content.Extra)
	if err != nil {
		fmt.Printf("[services/content] Create: UnmarshalContentJson error: %+v", err)
		responseExtra = &DataInfoExtra{}
	}
	res := &CreateResponse{
		DataInfo: DataInfo{
			ID:          content.ID,
			ContentID:   content.ContentID,
			Title:       content.Title,
			Description: content.Description,
			AuthorID:    content.AuthorID,
			Category:    content.Category,
			Type:        1, // 写死只有图文
			Body:        content.Body,
			BodyType:    3, // 先写死为Markdown
			Version:     content.Version,
			CreatedAt:   content.CreatedAt,
			UpdatedAt:   content.UpdatedAt,
			Extra:       *responseExtra,
		},
	}
	return res, nil
}

// DeleteRequest ...
type DeleteRequest struct {
	ContentIDs []int64
}

// DeleteResponse ...
type DeleteResponse struct {
	DataInfo
}

// Delete ...
func (cr ContentRepository) Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error) {
	content := &SDContent{}
	if dbc := cr.db.Where("content_id IN (?)", req.ContentIDs).Delete(content); dbc.Error != nil {
		fmt.Printf("[services/content] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		fmt.Printf("%+v\n", dbc)
	}
	return &DeleteResponse{
		DataInfo: sdContentToData(*content),
	}, nil
}

// UpdateRequest ...
type UpdateRequest struct {
	Target SDContent

	Title       string
	Description string
	Category    string
	Type        int16
	Body        string
	//Extra       DataInfoExtra
}

// UpdateResponse ...
type UpdateResponse struct {
	DataInfo
}

// Update ...
func (cr ContentRepository) Update(c context.Context, req UpdateRequest) (*UpdateResponse, error) {
	var target SDContent

	// TODO: 加上extra的update
	cr.db.Where("content_id = ?", req.Target.ContentID).Take(&target)
	var modified bool = false
	if req.Title != "" {
		target.Title = req.Title
		modified = true
	}
	if req.Description != "" {
		target.Description = req.Description
		modified = true
	}
	if req.Category != "" {
		target.Category = req.Category
		modified = true
	}
	// TODO: content_type自动解析赋值
	//if req.Type != 0 {
	//	target.Type = req.Type
	//	modified = true
	//}
	if req.Body != "" {
		target.Body = req.Body
		modified = true
	}

	if modified {
		target.Version += 1
	} else {
		return nil, errors.New("NoContentModified")
	}
	if dbc := cr.db.Save(&target); dbc.Error != nil {
		fmt.Printf("[services/content] Update: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	return &UpdateResponse{
		DataInfo: sdContentToData(target),
	}, nil

}


type GetRecommendByContentRequest struct {
	ContentID int64
}
type GetRecommendByContentResponse struct {
	ContentList []DataInfo
}
func (cr ContentRepository) GetRecommendByContent(c context.Context, req GetRecommendByContentRequest) (*GetRecommendByContentResponse, error) {
	var recommendContents []SDContent
	if dbc := cr.db.Limit(4).Order("updated_at desc").Find(&recommendContents); dbc.Error != nil {
		fmt.Printf("[services/content] Update: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}
	return &GetRecommendByContentResponse{
		ContentList: sdContentsToDatas(recommendContents),
	}, nil
}
