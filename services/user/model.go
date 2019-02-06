package user

import (
	"time"
)

//BaseExtra ...
type BaseInfoExtra struct {
	GithubHome string `json:"github_home"`
}

// BaseInfo ...
type BaseInfo struct {
	UserID    int64         `json:"user_id"`
	Name      string        `json:"name"`
	AvatarURL string        `json:"avatar_url"`
	Company   *string       `json:"company"`
	Email     *string       `json:"email"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Extra     BaseInfoExtra `json:"extra"`
}
