package sdlog

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
)

type SDLogServiceInterface interface {
	Find(req FindRequest) (*FindResponse, error)
	Create(req CreateRequest) (*CreateResponse, error)
	// Delete(req DeleteRequest) (*DeleteResponse, error)
}

type SDLogService struct {
	db        *gorm.DB
	idBuilder *goSnowFlake.IdWorker
	SDLogServiceInterface
}

func newSDLogService(db *gorm.DB) *SDLogService {
	idBuilder, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Printf("[services/sdlog] Init snowFlake id_builder error: %+v", err)
		os.Exit(1)
	}
	return &SDLogService{
		db:        db,
		idBuilder: idBuilder,
	}
}
