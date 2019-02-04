package content

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zheng-ji/goSnowFlake"
	"os"
)

type ContentServiceInterface interface {
	// Put adds a new Greeting to the datastore
	FindOne(req FindOneRequest) (*FindOneResponse, error)
	Find(req FindRequest) (*FindResponse, error)
	Create(req CreateRequest) (*CreateResponse, error)
	Delete(req DeleteRequest) (*DeleteResponse, error)
	Update(req UpdateRequest) (*UpdateResponse, error)
	GetRecommendByContent(req GetRecommendByContentRequest) (*GetRecommendByContentResponse, error)
}

type ContentService struct {
	db        *gorm.DB
	idBuilder *goSnowFlake.IdWorker
	ContentServiceInterface
}

func newContentService(db *gorm.DB) *ContentService {
	idBuilder, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		fmt.Printf("[services/content] Init snowFlake id_builder error: %+v", err)
		os.Exit(1)
	}
	return &ContentService{
		db:        db,
		idBuilder: idBuilder,
	}
}
