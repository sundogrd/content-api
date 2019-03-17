package sdlog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/sdlog"
)

type LogCountRequest struct {
	UserID   string          `form:"user_id"`
	TargetID string          `form:"target_id"`
	Type     sdlog.SDLogType `form:"type"`
}

type LogCountResponse struct {
	Count int64 `json:"count"`
}

func GetStatementCount(c *gin.Context) {
	var request LogCountRequest = LogCountRequest{
		TargetID: "",
		UserID:   "",
		Type:     "",
	}
	var userID, targetID int64
	var logType sdlog.SDLogType
	var (
		user    *int64
		target  *int64
		logtype *sdlog.SDLogType
	)
	var err error
	var res *sdlog.CountResponse

	if c.BindQuery(&request) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if request.TargetID != "" {
		targetID, err = strconv.ParseInt(request.TargetID, 10, 64)
		if err != nil {
			panic(err)
		}
		target = &targetID
	}
	if request.UserID != "" {
		userID, err = strconv.ParseInt(request.UserID, 10, 64)
		if err != nil {
			panic(err)
		}
		user = &userID
	}

	if request.Type != "" {
		logType = request.Type
		logtype = &logType
	}

	res, err = sdlog.SDLogServiceInstance().Count(sdlog.CountRequest{
		UserID:   user,
		TargetID: target,
		Type:     logtype,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, LogCountResponse{
		Count: res.Count,
	})
}
