package content

import (
	dbUtils "github.com/sundogrd/content-api/utils/db"
	"sync"
)

var _contentService *ContentService
var _contentServiceOnce sync.Once

func ContentServiceInstance() *ContentService {
	_contentServiceOnce.Do(func() {
		db := dbUtils.Client
		hasContentTable := db.HasTable(&SDContent{})
		if hasContentTable == false {
			db.CreateTable(&SDContent{})
		} else {
			db.AutoMigrate(&SDContent{})
		}

		hasContentCountTable := db.HasTable(&SDContentCount{})
		if hasContentCountTable == false {
			db.CreateTable(&SDContentCount{})
		} else {
			db.AutoMigrate(&SDContentCount{})
		}
		_contentService = newContentService(db)
	})
	return _contentService
}
