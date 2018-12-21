package handler_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content-api/routes"
	"github.com/sundogrd/content-api/utils/db"

	"github.com/sundogrd/content-api/utils/config"
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
	routes.Routes(r)
	return r, nil
}

func testBase(t *testing.T, method string, url string, body io.Reader) {
	router, err := initContext()
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	req.Header.Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
	req.Header.Set("Access-Control-Expose-Headers", "Content-Length")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	t.Logf("contents: %+v", w.Body.String())
}
func TestGetContent(t *testing.T) {
	testBase(t, "GET", "/contents/301523989593853952", nil)
}
func TestGetContents(t *testing.T) {
	testBase(t, "GET", "/contents?title=fuck", nil)
}

func TestUpdateContent(t *testing.T) {
	jsonStr := []byte(`{
		"title": "mdzz chnages"
	}`)
	testBase(t, "PATCH", "/contents/301523989593853952", bytes.NewBuffer(jsonStr))
}

func TestCreateContent(t *testing.T) {
	jsonStr := []byte(`{
		"title": "new test itme"
	}`)
	testBase(t, "POST", "/contents", bytes.NewBuffer(jsonStr))
}

func TestDeleteContent(t *testing.T) {
	testBase(t, "DELETE", "/contents/301835252127502336", nil)
}

func TestDeleteContents(t *testing.T) {
	testBase(t, "DELETE", "/contents?title=new", nil)
}
