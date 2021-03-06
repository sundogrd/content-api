package content

// FindRequest ...
type FindRequest struct {
	ContentIDs *[]int64
	Title      *string
	AuthorID   *int64
	Status     *ContentStatus
	Type       *ContentType // content类型
	Page       *int16       // 页数从1开始，默认1
	PageSize   *int16       // 页项，默认10
}

// FindResponse ...
type FindResponse struct {
	List  []BaseInfo
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
	if req.Title != nil {
		db = db.Where("title LIKE ?", "%"+*req.Title+"%")
	}
	if req.AuthorID != nil {
		db = db.Where("author_id = ", *req.AuthorID)
	}
	if req.Type != nil {
		db = db.Where("type = ?", *req.Type)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	db.Limit(pageSize).Offset((page - 1) * (pageSize))
	if err := db.Find(&contents).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	} else {
		BaseInfos := make([]BaseInfo, 0)
		for _, v := range contents {
			BaseInfos = append(BaseInfos, packBaseInfo(v))
		}
		res := &FindResponse{
			List:  BaseInfos,
			Total: count,
		}
		return res, nil
	}
}
