package redis_test

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/redis"
	"testing"
)

func initConfig() {
	config.Init()
	viper.SetConfigType("json") // or viper.SetConfigType("YAML")

	var jsonConfig = []byte(`{
    	"redis": {
    	    "host": "127.0.0.1",
    	    "port": 6379,
    	    "password": "CCq2Si39hdgY6ajP5vHL"
    	}
	}`)
	viper.ReadConfig(bytes.NewBuffer(jsonConfig))
}

// go test -run="NoopsGetContentListWithFilter"
func TestInitRedis(t *testing.T) {
	initConfig()
	err := redis.Init(0)
	client := redis.DB(0)
	//client.Set("woca", "keke", time.Hour)
	test, err := client.Get("made").Result()
	t.Logf("get result: %+v", test)
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(resp)
	// realResp := resp.RealResponse().(*noops.GetContentListWithFilterResponse)
	// t.Log(realResp.ContentId)
}
