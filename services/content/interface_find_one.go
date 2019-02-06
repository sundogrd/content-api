package content

// FindOneRequest ...
type FindOneRequest struct {
	ID        int64
	ContentID int64
}

// FindOneResponse ...
type FindOneResponse struct {
	BaseInfo
}

// FindOne ...
func (cs ContentService) FindOne(req FindOneRequest) (*FindOneResponse, error) {
	var content SDContent
	cs.db.Where(&SDContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	}).First(&content)

	res := &FindOneResponse{
		packBaseInfo(content),
	}
	return res, nil
}
