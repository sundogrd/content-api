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
		hasContentTable := db.HasTable(&PFContent{})
		if hasContentTable == false {
			db.CreateTable(&PFContent{})
		}
		_contentRepository = newContentRepository(db)
		// &ContentRepository{
		// KitcClient: kitcClient,
		// cache??
		// }
	})
	return _contentRepository
}
