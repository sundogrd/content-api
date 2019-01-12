package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubAuth(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	req.Header.Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
	req.Header.Set("Access-Control-Expose-Headers", "Content-Length")
	router.ServeHTTP(w, req)

	response, err := http.Get("https://api.github.com/user?access_token=" + token.AccessToken)

	assert.Equal(t, "兴趣使然", w.Code)
	t.Logf("contents: %+v", w.Body.String())
}
