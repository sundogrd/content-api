package user

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
)

type UserServiceInterface interface {
	FindOne(req FindOneRequest) (*FindOneResponse, error)
	Create(req CreateRequest) (*CreateResponse, error)
	Delete(req DeleteRequest) (*DeleteResponse, error)
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
