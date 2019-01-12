package sdsession

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	redisUtil "github.com/sundogrd/content-api/utils/redis"
	"sync"
	"time"
)

var Redis_Prefix = "sundog:sessions:"
var Redis_life = time.Hour

type redisStore struct {
	redis *redis.Client
	lock  sync.RWMutex
}

func (r *redisStore) Get(k string, sid string) interface{} {
	r.lock.RLock()
	defer r.lock.RUnlock()
	jsondata, err := r.getSessionData(sid)

	if err != nil {
		return nil
	}
	return jsondata[k]
}
func (r *redisStore) Set(k string, v interface{}, sid string) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	jsondata, err := r.getSessionData(sid)
	if err != nil {
		fmt.Printf("[middlewares/sdsession] set error: %+v", err)
		fmt.Println()
		jsondata = make(map[string]interface{})
	}

	jsondata[k] = v
	jsonString, err := json.Marshal(jsondata)
	if err != nil {
		fmt.Println(err)
		return err
	}
	redisstring := string(jsonString)

	r.redis.Set(Redis_Prefix+sid, redisstring, Redis_life)
	return nil
}
func (r *redisStore) getSessionData(sid string) (jsondata map[string]interface{}, err error) {
	if r.redis == nil {
		fmt.Println("redis == nil")
		r.initRedisConn()
	}
	fmt.Printf("redis = %+v \n", r.redis)
	data, err := r.redis.Get(Redis_Prefix + sid).Result()
	if err != nil {
		fmt.Printf("[middlewares/sdsession] getSessionData error, sid: %s, err: %+v\n", sid, err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(data), &jsondata); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsondata, nil
}
func (r *redisStore) initRedisConn() {
	client := redisUtil.DB(0)

	r.redis = client
}
func init() {
	driver := new(redisStore)
	Register("redis", driver)
}
