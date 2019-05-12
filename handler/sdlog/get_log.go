package sdlog

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sundogrd/content-api/services/sdlog"
)

type GetLogRequest struct {
	LogID    string          `form:"log_id"`
	UserID   string          `form:"user_id"`
	TargetID string          `form:"target_id"`
	Type     sdlog.SDLogType `form:"type"`
	Page     string          `form:"page"`
	PageSize string          `form:"page_size"`
}

type GetLogResponse struct {
	List  []sdlog.SDLog `json:"list"`
	Total int64         `json:"total"`
}

// GetStatementBase...
func GetStatementBase(c *gin.Context) (*sdlog.FindResponse, error) {
	var request GetLogRequest
	var logID, userID, targetID int64
	var logType sdlog.SDLogType
	var err error
	var page, pageSize int16
	var res *sdlog.FindResponse

	var log, user, target *int64
	var p, ps *int16
	var logtype *sdlog.SDLogType

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return nil, err
	}
	// fmt.Printf("%+v", request)
	if request.LogID != "" {
		logID, err = strconv.ParseInt(request.LogID, 10, 64)
		if err != nil {
			panic(err)
		}
		log = &logID
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

	if request.Page != "" {
		_page, err := strconv.ParseInt(request.Page, 10, 16)
		if err != nil {
			panic(err)
		}
		page = int16(_page)
		p = &page
	}
	if request.PageSize != "" {
		_pageSize, err := strconv.ParseInt(request.PageSize, 10, 16)
		if err != nil {
			panic(err)
		}
		pageSize = int16(_pageSize)
		ps = &pageSize
	}

	res, err = sdlog.SDLogServiceInstance().Find(sdlog.FindRequest{
		LogID:    log,
		UserID:   user,
		TargetID: target,
		Type:     logtype,
		Page:     p,
		PageSize: ps,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return nil, err
	}
	return res, nil
}

// GetStatement...
func GetStatement(c *gin.Context) {
	res, err := GetStatementBase(c)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, GetLogResponse{
		List:  res.List,
		Total: res.Total,
	})
}
