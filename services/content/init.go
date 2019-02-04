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
		}
		_contentService = newContentService(db)
	})
	return _contentService
}
