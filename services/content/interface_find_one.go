package content

// FindOneRequest ...
type FindOneRequest struct {
	ID        int64
	ContentID int64
}

// FindOneResponse ...
type FindOneResponse struct {
	FullInfo
}

// FindOne ...
func (cs ContentService) FindOne(req FindOneRequest) (*FindOneResponse, error) {
	var content SDContent
	var contentCount SDContentCount
	cs.db.Where(&SDContent{
		ID:        req.ID,
		ContentID: req.ContentID,
	}).First(&content)
	cs.db.Where(&SDContentCount{
		ContentID: req.ContentID,
		CountKey: "read_count",
	}).First(&contentCount)

	res := &FindOneResponse{
		packFullInfo(content, contentCount),
	}
	return res, nil
}
