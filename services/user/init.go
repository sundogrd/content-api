package user

import (
	dbUtils "github.com/sundogrd/content-api/utils/db"
	"sync"
)

var _userService *UserService
var _userServiceOnce sync.Once

func UserServiceInstance() *UserService {
	_userServiceOnce.Do(func() {
		db := dbUtils.Client
		hasContentTable := db.HasTable(&SDUser{})
		if hasContentTable == false {
			db.CreateTable(&SDUser{})
		}
		_userService = newUserService(db)
	})
	return _userService
}
