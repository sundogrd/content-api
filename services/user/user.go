package user

import (
	"context"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
)

type UserServiceInterface interface {
	FindOne(c context.Context, req FindOneRequest) *FindOneResponse
	Create(c context.Context, req CreateRequest) (*CreateResponse, error)
	Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error)
}

type UserService struct {
	db        *gorm.DB
	idBuilder *goSnowFlake.IdWorker
	UserServiceInterface
}

func newUserService(db *gorm.DB) *UserService {
	idBuilder, err := goSnowFlake.NewIdWorker(3)
	if err != nil {
		fmt.Printf("[services/user] Init snowFlake id_builder error: %+v", err)
		os.Exit(1)
	}
	return &UserService{
		db:        db,
		idBuilder: idBuilder,
	}
}

// FindOneRequest ...
type FindOneRequest struct {
	UserID *int64
	Name   *string
}

// FindOneResponse ...
type FindOneResponse struct {
	DataInfo
}

// FindOne ...
func (cr UserService) FindOne(c context.Context, req FindOneRequest) *FindOneResponse {
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
		sdUserToData(user),
	}
	if res.ID == 0 {
		return nil
	}
	return res
}

// CreateRequest ...
type CreateRequest struct {
	Name      string
	AvatarUrl string
	Company   string
	Email     string
	Extra     DataInfoExtra
}

// CreateResponse ...
type CreateResponse struct {
	DataInfo
}

// Create ...
func (us UserService) Create(c context.Context, req CreateRequest) (*CreateResponse, error) {
	// TODO: Duplicate title?

	// TODO: Param validating

	userExtraStr, err := marshalUserExtraJson(&req.Extra)
	if err != nil {
		fmt.Printf("[services/user] Create: json marshal error: %+v", err)
		userExtraStr, _ = marshalUserExtraJson(&DataInfoExtra{})
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

	responseExtra, err := UnmarshalUserExtraJson(user.Extra)
	if err != nil {
		fmt.Printf("[services/user] Create: UnmarshalUserExtraJson error: %+v", err)
		responseExtra = &DataInfoExtra{}
	}
	res := &CreateResponse{
		DataInfo: DataInfo{
			ID:        user.ID,
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

// DeleteRequest ...
type DeleteRequest struct {
	UserID int64
}

// DeleteResponse ...
type DeleteResponse struct {
	DataInfo
}

// Delete ...
func (us UserService) Delete(c context.Context, req DeleteRequest) (*DeleteResponse, error) {
	user := &SDUser{}
	if dbc := us.db.Where("user_id = (?)", req.UserID).Delete(user); dbc.Error != nil {
		fmt.Printf("[services/user] Delete: db createerror: %+v", dbc.Error)
		// Create failed, do something e.g. return, panic etc.
		return nil, dbc.Error
	} else {
		fmt.Printf("%+v\n", dbc)
	}
	return &DeleteResponse{
		DataInfo: sdUserToData(*user),
	}, nil
}
