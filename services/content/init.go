package content

import (
	dbUtils "github.com/sundogrd/content-api/utils/db"
	"sync"
)

var _contentRepository *ContentRepository
var _contentRepositoryOnce sync.Once

func ContentRepositoryInstance() *ContentRepository {
	_contentRepositoryOnce.Do(func() {
		db := dbUtils.Client
		hasContentTable := db.HasTable(&SDContent{})
		if hasContentTable == false {
			db.CreateTable(&SDContent{})
		}
		_contentRepository = newContentRepository(db)
	})
	return _contentRepository
}
