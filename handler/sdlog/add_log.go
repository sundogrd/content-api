package sdlog

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/sundogrd/content-api/services/sdlog"

	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	UserID    string           `json:"user_id"`
	TargetID  string           `json:"target_id"`
	Type      sdlog.SDLogType  `json:"type"`
	Extra     sdlog.SDLogExtra `json:"extra"`
	CreatedAt time.Time        `json:"created_at"`
}

type LogResponse struct {
	UserID   string           `json:"user_id"`
	TargetID string           `json:"target_id"`
	Type     sdlog.SDLogType  `json:"type"`
	ID       string           `json:"id"`
	Extra    sdlog.SDLogExtra `json:"extra"`
}

func AddStatement(c *gin.Context) {
	var request LogRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"msg": err.Error(),
		})
		return
	}
	targetID, err := strconv.ParseInt(request.TargetID, 10, 64)
	if err != nil {
		panic(err)
	}

	userID, err := strconv.ParseInt(request.UserID, 10, 64)
	if err != nil {
		panic(err)
	}
	res, err := sdlog.SDLogServiceInstance().Create(sdlog.CreateRequest{
		TargetID:  targetID,
		UserID:    userID,
		Type:      request.Type,
		CreatedAt: request.CreatedAt,
		Extra: sdlog.SDLogExtra{
			Count:  request.Extra.Count,
			Detail: request.Extra.Detail,
		},
	})
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, LogResponse{
		Type:     res.Type,
		ID:       strconv.FormatInt(res.LogID, 10),
		UserID:   strconv.FormatInt(res.UserID, 10),
		TargetID: strconv.FormatInt(res.TargetID, 10),
		Extra:    res.Extra,
	})
}
