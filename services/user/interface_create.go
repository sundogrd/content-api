package user

import "fmt"

// CreateRequest ...
type CreateRequest struct {
	Name      string
	AvatarUrl string
	Company   *string
	Email     *string
	Extra     UserInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	UserInfo
}

// Create ...
func (us UserService) Create(req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	userExtraStr, err := marshalUserExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/user] Create: json marshal error: %+v", err)
		userExtraStr, _ = marshalUserExtraJson(&UserInfoExtra{})
	}

	userId, _ := us.idBuilder.NextId()
	user := SDUser{
		UserID:    userId,
		Name:      req.Name,
		AvatarUrl: req.AvatarUrl,
		Company:   req.Company,
		Email:     req.Email,
		Extra:     userExtraStr,
	}
	if dbc := us.db.Create(&user); dbc.Error != nil {
		fmt.Printf("[services/user] Create: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	}

	responseExtra, err := unmarshalUserExtraJson(user.Extra)
	if err != nil {
		fmt.Printf("[services/user] Create: UnmarshalUserExtraJson error: %+v", err)
		responseExtra = &UserInfoExtra{}
	}
	res := &CreateResponse{
		UserInfo: UserInfo{
			UserID:    user.UserID,
			Name:      user.Name,
			AvatarUrl: user.AvatarUrl,
			Company:   user.Company,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Extra:     *responseExtra,
		},
	}
	return res, nil
}
