package sdlog

type CountRequest struct {
	TargetID *int64
	UserID   *int64
	Type     *SDLogType
}

type CountResponse struct {
	Count int64
}

// Count...
func (ss SDLogService) Count(req CountRequest) (*CountResponse, error) {

	logModels := make([]SDLogModel, 0)
	count := int64(0)

	// fmt.Printf("[Service/Count] request is %+v", *req.TargetID)

	db := ss.db

	if req.TargetID != nil {
		db = db.Where("target_id = ?", *req.TargetID)
	}
	if req.UserID != nil {
		db = db.Where("user_id = ?", *req.UserID)
	}
	if req.Type != nil {
		db = db.Where("type = ?", *req.Type)
	}

	if err := db.Find(&logModels).Count(&count).Error; err != nil {
		return nil, err
	} else {
		res := &CountResponse{
			Count: count,
		}
		return res, nil
	}
}
