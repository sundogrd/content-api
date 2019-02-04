package content

// FindOneRequest ...
type FindOneRequest struct {
	ID        int64
	ContentID int64
}

// FindOneResponse ...
type FindOneResponse struct {
	ContentInfo
}

// FindOne ...
func (cs ContentService) FindOne(req FindOneRequest) (*FindOneResponse, error) {
	var content SDContent
	cs.db.Where(&SDContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	}).First(&content)

	res := &FindOneResponse{
		packContentInfo(content),
	}
	return res, nil
}
