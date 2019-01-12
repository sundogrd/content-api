package content

import (
	"time"
)

type DataInfoExtra struct {
	StarNum int64 `json:"star_num"`
}

type DataInfo struct {
	ID          int64         `json:"id"`
	ContentID   int64         `json:"content_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	AuthorID    int64         `json:"author_id"`
	Category    string        `json:"category"`
	Type        ContentType   `json:"type"`
	Body        string        `json:"body"`
	BodyType    BodyType      `json:"body_type"`
	Version     int16         `json:"version"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Extra       DataInfoExtra `json:"extra"`
}
