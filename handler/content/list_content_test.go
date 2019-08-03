package content_test

import (
	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/handler/content"
	"github.com/sundogrd/content-api/utils/config"
	"github.com/sundogrd/content-api/utils/test"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initListContentContext() (*gin.Engine, error) {
	config.Init()
	r := gin.Default()
	return r, nil
}

func TestListContent(t *testing.T) {
	r, err := initListContentContext()
	if err != nil {
		t.Fail()
	}

	container, err := test.InitTestContainer()
	if err != nil {
		t.Fatal(err)
	}

	r.GET("/contents", content.ListContent(container))

	req, _ := http.NewRequest("GET", "/contents", nil)

	test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {

		}
		t.Logf("%#v", p)
		return statusOK
	})
}
