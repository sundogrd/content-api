package content

import (
	"time"
)

type BaseInfoExtra struct {
	StarNum int64 `json:"star_num"`
}

type BaseInfo struct {
	ContentID   int64         `json:"content_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	AuthorID    int64         `json:"author_id"`
	Category    string        `json:"category"`
	Type        ContentType   `json:"type"`
	Status      ContentStatus `json:"status"`
	Body        string        `json:"body"`
	BodyType    BodyType      `json:"body_type"`
	Version     int16         `json:"version"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Extra       BaseInfoExtra `json:"extra"`
}

type FullInfoExtra struct {
	StarNum int64 `json:"star_num"`
	ReadNum int64 `json:"read_num"`
}
type FullInfo struct {
	ContentID   int64             `json:"content_id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	AuthorID    int64             `json:"author_id"`
	Category    string            `json:"category"`
	Type        ContentType       `json:"type"`
	Status      ContentStatus `json:"status"`
	Body        string            `json:"body"`
	BodyType    BodyType          `json:"body_type"`
	Version     int16             `json:"version"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Extra       FullInfoExtra `json:"extra"`
}
