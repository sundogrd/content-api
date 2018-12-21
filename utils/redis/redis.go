package redis

import (
	"github.com/go-redis/redis"

	"github.com/sundogrd/content-api/utils/config"
	// "github.com/sundogrd/content-api/utils/log"
)

var dbs map[int]*redis.Client = make(map[int]*redis.Client)

// Client Redis client
type Client = redis.Client

// Options ...
type Options = redis.Options

// func init() error {
// 	dbs = make(map[int]*redis.Client)
// 	return err
// }

// Init 初始化数据库
func Init(db int) error {
	var err error
	dbs[db], err = Conn(db)
	return err
}

// DB 获取连接的实例
func DB(db int) *redis.Client {
	if v, ok := dbs[db]; ok {
		return v
	}
	return nil
}

// ConnectDB 连接redis
func ConnectDB(opt *redis.Options) (*redis.Client, error) {
	client := redis.NewClient(opt)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Connect 用配置文件连接数据库
func Connect(db int, prePath string) (*redis.Client, error) {
	host := config.GetString(prePath + "host")
	port := config.GetString(prePath + "port")
	password := config.GetString(prePath + "password")

	client, err := ConnectDB(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Conn 用配置文件的默认参数连接数据库
func Conn(db int) (*redis.Client, error) {
	return Connect(db, "redis.")
}
