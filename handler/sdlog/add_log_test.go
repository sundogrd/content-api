package sdlog_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sundogrd/content-api/handler/sdlog"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/db"
	"github.com/sundogrd/content-api/utils/test"
)

func initContext() (*gin.Engine, error) {
	config.Init()
	// 初始化数据库
	viper.SetConfigType("json") // or viper.SetConfigType("YAML")
	var jsonConfig = []byte(`{
	  	"db": {
			"type": "mysql",
			"options": {
				"user": "root",
				"password": "breakIN",
				"host": "localhost",
				"port": 3306,
				"dbname": "sundog",
				"connectTimeout": "10s"
			}
	  	}
	}`)
	viper.ReadConfig(bytes.NewBuffer(jsonConfig))
	_, err := db.Init()
	if err != nil {
		return nil, err
	}
	r := gin.Default()
	return r, nil
}

func TestAddLog(t *testing.T) {
	r, err := initContext()
	if err != nil {
		t.Fail()
	}

	r.POST("/statement", sdlog.AddStatement)

	var jsonStr = []byte(`
	{
		"target_id": "23223",
		"user_id": "33",
		"type": "clap",
		"extra": {
		"count": 55,
		"detail": "hahah"
		}
		}`)

	req, _ := http.NewRequest("POST", "/statement", bytes.NewBuffer(jsonStr))

	test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fail()
		}
		t.Logf("%#v", p)
		return statusOK
	})
}
