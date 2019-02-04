package content

import (
	"time"
)

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
	List  []ContentInfo
	Total int64
}

// Find ...
func (cs ContentService) Find(req FindRequest) (*FindResponse, error) {
	var page int16 = 1
	var pageSize int16 = 10
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		pageSize = *req.PageSize
	}

	contents := make([]SDContent, 0)
	count := int64(0)

	db := cs.db
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

	db.Limit(pageSize).Offset((page - 1) * (pageSize))
	if err := db.Find(&contents).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	} else {
		contentInfos := make([]ContentInfo, 0)
		for _, v := range contents {
			contentInfos = append(contentInfos, packContentInfo(v))
		}
		res := &FindResponse{
			List:  contentInfos,
			Total: count,
		}
		return res, nil
	}
}
