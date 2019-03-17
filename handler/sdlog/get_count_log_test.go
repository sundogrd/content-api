package sdlog_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sundogrd/content-api/handler/sdlog"
	"github.com/sundogrd/content-api/utils/test"
)

func TestGetLogCount(t *testing.T) {
	r, err := initContext()
	if err != nil {
		t.Fail()
	}

	r.GET("/statement/count", sdlog.GetStatementCount)

	req, _ := http.NewRequest("GET", "/statement/count?target_id=23223", nil)

	test.TestHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			t.Fail()
		}
		fmt.Println(p)
		t.Logf("%+v", p)
		return statusOK
	})
}
