package sdlog

import (
	"time"
)

type SDLogType string

const (
	Clap       SDLogType = "CLAP"
	CancelClap SDLogType = "CANCELCLAP"
	Login      SDLogType = "LOGIN"
	Logout     SDLogType = "LOGOUT"
	Publish    SDLogType = "PUBLISH"
	SIGN       SDLogType = "SIGN"
)

type SDLogExtra struct {
	Count  int64  `json:"count"`
	Detail string `json:"detail"`
}

type SDLog struct {
	LogID     int64      `json:"log_id"`
	TargetID  int64      `json:"target_id"`
	UserID    int64      `json:"user_id"`
	Type      SDLogType  `json:"type"`
	Extra     SDLogExtra `json:"extra"`
	CreatedAt time.Time  `json:"created_at"`
}
