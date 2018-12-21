package db

import (
	"github.com/sundogrd/content-api/utils/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Client 默认数据库实例
var Client *gorm.DB

// Init 默认初始化
func Init() (*gorm.DB, error) {
	var err error
	Client, err = Conn()
	return Client, err
}

// ConnectDB 用连接字符串连接数据库
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Connect 用配置文件连接数据库
func Connect(prePath string) (*gorm.DB, error) {
	user := config.GetString(prePath + "user")
	password := config.GetString(prePath + "password")
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	dbname := config.GetString(prePath + "dbname")
	connectTimeout := config.GetString(prePath + "connectTimeout")

	db, err := ConnectDB(user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + dbname + "?charset=utf8&parseTime=True&loc=Local&timeout=" + connectTimeout)
	if err != nil {
		return nil, err
	}
	// if config.Get
	db.LogMode(true)
	return db, nil
}

// Conn 用配置文件的默认参数连接数据库
func Conn() (*gorm.DB, error) {
	return Connect("db.options.")
}
