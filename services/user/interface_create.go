package user

import "fmt"

// CreateRequest ...
type CreateRequest struct {
	Name      string
	AvatarURL string
	Company   *string
	Email     *string
	Extra     BaseInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	BaseInfo
}

// Create ...
func (us UserService) Create(req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	userExtraStr, err := marshalUserExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/user] Create: json marshal error: %+v", err)
		userExtraStr, _ = marshalUserExtraJson(&BaseInfoExtra{})
	}

	userID, _ := us.idBuilder.NextId()
	user := SDUser{
		UserID:    userID,
		Name:      req.Name,
		AvatarURL: req.AvatarURL,
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
		responseExtra = &BaseInfoExtra{}
	}
	res := &CreateResponse{
		BaseInfo: BaseInfo{
			UserID:    user.UserID,
			Name:      user.Name,
			AvatarURL: user.AvatarURL,
			Company:   user.Company,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Extra:     *responseExtra,
		},
	}
	return res, nil
}
