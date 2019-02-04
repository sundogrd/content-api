package user

// FindOneRequest ...
type FindOneRequest struct {
	UserID *int64
	Name   *string
}

// FindOneResponse ...
type FindOneResponse struct {
	UserInfo
}

// FindOne ...
func (cr UserService) FindOne(req FindOneRequest) *FindOneResponse {
	var user SDUser
	if req.UserID != nil {
		cr.db.Where(&SDUser{
			UserID: *req.UserID,
		}).First(&user)
	} else if req.Name != nil {
		cr.db.Where(&SDUser{
			Name: *req.Name,
		}).First(&user)
	}

	res := &FindOneResponse{
		packUserInfo(user),
	}
	if res.UserID == 0 {
		return nil
	}
	return res
}
