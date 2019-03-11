package sdlog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/sdlog"
)

type GetLogRequest struct {
	LogID    string          `json:"log_id"`
	UserID   string          `json:"user_id"`
	TargetID string          `json:"target_id"`
	Type     sdlog.SDLogType `json:"type"`
	Page     string          `json:"page"`
	PageSize string          `json:"page_size"`
}

type GetLogResponse struct {
	List  []sdlog.SDLog `json:"list"`
	Total int64         `json:"total"`
}

func GetStatement(c *gin.Context) {
	var request GetLogRequest
	var logID, userID, targetID int64
	var logType sdlog.SDLogType
	var err error
	var page, pageSize int16
	var res *sdlog.FindResponse

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if request.LogID != "" {
		logID, err = strconv.ParseInt(request.LogID, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if request.TargetID != "" {
		targetID, err = strconv.ParseInt(request.TargetID, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	if request.UserID != "" {
		userID, err = strconv.ParseInt(request.UserID, 10, 64)
		if err != nil {
			panic(err)
		}
	}

	if request.Type != "" {
		logType = request.Type
	}

	if request.Page != "" {
		_page, err := strconv.ParseInt(request.Page, 10, 16)
		if err != nil {
			panic(err)
		}
		page = int16(_page)
	}
	if request.PageSize != "" {
		_pageSize, err := strconv.ParseInt(request.PageSize, 10, 16)
		if err != nil {
			panic(err)
		}
		pageSize = int16(_pageSize)
	}

	res, err = sdlog.SDLogServiceInstance().Find(sdlog.FindRequest{
		LogID:    &logID,
		UserID:   &userID,
		TargetID: &targetID,
		Type:     &logType,
		Page:     &page,
		PageSize: &pageSize,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	c.JSON(http.StatusOK, GetLogResponse{
		List:  res.List,
		Total: res.Total,
	})
}
