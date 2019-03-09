package sdlog

import (
	"sync"

	dbUtils "github.com/sundogrd/content-api/utils/db"
)

var _SDLogService *SDLogService
var _SDLogServiceOnce sync.Once

func CreateSDLogServiceInstance() *SDLogService {
	_SDLogServiceOnce.Do(func() {
		db := dbUtils.Client
		hasContentTable := db.HasTable(&SDLogModel{})
		if hasContentTable == false {
			db.CreateTable(&SDLogModel{})
		} else {
			db.AutoMigrate(&SDLogModel{})
		}
		_SDLogService = newSDLogService(db)
	})
	return _SDLogService
}
