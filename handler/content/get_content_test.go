package content_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/db"
	"testing"
)

func initContext() (*gin.Engine, error) {
	config.Init()
	// 初始化数据库
	viper.SetConfigType("json") // or viper.SetConfigType("YAML")
	var jsonConfig = []byte(`{
	  	"db": {
			"type": "mysql",
			"options": {
				"user": "sundog",
				"password": "sundogPwd",
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

func TestGetContent(t *testing.T) {
	//r, err := initContext()
	//if err != nil {
	//	t.Fail()
	//}
	//
	//r.GET("/contents/:contentId", content.GetContent)
	//
	//req, _ := http.NewRequest("GET", "/contents/303983137602408448", nil)
	//
	//test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
	//	statusOK := w.Code == http.StatusOK
	//
	//	p, err := ioutil.ReadAll(w.Body)
	//	if err != nil {
	//
	//	}
	//	t.Logf("%#v", p)
	//	return statusOK
	//})
}
