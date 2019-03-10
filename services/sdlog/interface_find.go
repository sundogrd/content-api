package sdlog

import "time"

// FindRequest ...
type FindRequest struct {
	LogID     *int64
	TargetID  *int64
	UserID    *int64
	Type      *SDLogType
	Page      *int16 // 页数从1开始，默认1
	PageSize  *int16 // 页项，默认10
	CreatedAt *time.Time
}

// FindResponse ...
type FindResponse struct {
	List  []SDLog
	Total int64
}

// Find ...
func (ss SDLogService) Find(req FindRequest) (*FindResponse, error) {
	var page int16 = 1
	var pageSize int16 = 10
	if req.Page != nil {
		page = *req.Page
	}
	if req.PageSize != nil {
		pageSize = *req.PageSize
	}

	logModels := make([]SDLogModel, 0)
	count := int64(0)

	db := ss.db
	if req.LogID != nil {
		db.Where("id= ?", *req.LogID)
	} else {
		if req.TargetID != nil {
			db = db.Where("target_id = ?", *req.TargetID)
		}
		if req.UserID != nil {
			db = db.Where("user_id = ?", *req.UserID)
		}
		if req.Type != nil {
			db = db.Where("type = ?", *req.Type)
		}
		// if req.CreatedAt != nil {
		// 	db = db.Where("created_at = ?", *req.Type)
		// }
	}

	db.Limit(pageSize).Offset((page - 1) * (pageSize))
	if err := db.Find(&logModels).Offset(0).Limit(-1).Count(&count).Error; err != nil {
		return nil, err
	} else {
		sdlogs := make([]SDLog, 0)
		for _, v := range logModels {
			sdlogs = append(sdlogs, packSDLog(v))
		}
		res := &FindResponse{
			List:  sdlogs,
			Total: count,
		}
		return res, nil
	}
}
