package user

import (
	"time"
)

type DataInfoExtra struct {
	GithubHome string `json:"github_home"`
}

type DataInfo struct {
	ID        int64         `json:"id"`
	UserID    int64         `json:"user_id"`
	Name      string        `json:"name"`
	AvatarUrl string        `json:"avatar_url"`
	Company   string        `json:"company"`
	Email     string        `json:"email"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Extra     DataInfoExtra `json:"extra"`
}
