package uploader

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/sundogrd/content-api/middlewares/sdsession"
	"github.com/sundogrd/content-api/utils/config"
	"net/http"
)

// GetQiniuTokenRequest ...
type GetQiniuTokenRequest struct {
}

// GetQiniuTokenResponse ...
type GetQiniuTokenResponse struct {
	UpToken string `json:"uptoken"`
}

// GetQiniuToken ...
func GetQiniuToken(c *gin.Context) {
	//var request GetQiniuTokenRequest

	sess := sdsession.GetSession(c)
	if sess.Get("user_id") == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "user not login",
		})
		return
	}

	bucket := config.GetString("auth.qiniu.bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(config.GetString("auth.qiniu.accessKey"), config.GetString("auth.qiniu.secretKey"))
	upToken := putPolicy.UploadToken(mac)

	c.JSON(http.StatusOK, GetQiniuTokenResponse{
		UpToken: upToken,
	})
}
