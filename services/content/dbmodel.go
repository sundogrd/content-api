package content

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BodyType int16

const (
	BODY_TYPE_TEXT     BodyType = 1
	BODY_TYPE_HTML     BodyType = 2
	BODY_TYPE_MARKDOWN BodyType = 3
	BODY_TYPE_LATEX    BodyType = 4
)

type ContentType int16

const (
	CONTENT_TYPE_TEXT  ContentType = 1
	CONTENT_TYPE_AUDIO ContentType = 2
	CONTENT_TYPE_VIDEO ContentType = 4
)

// SDContent http://gorm.io/docs/models.html
type SDContent struct {
	// gorm.Model
	ID          int64       `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID   int64       `gorm:"not null;"`
	Title       string      `gorm:"type:varchar(60);not null"`
	Description string      `gorm:"type:varchar(300);not null"`
	AuthorID    int64       `gorm:"not null;"`
	Category    string      `gorm:"type:varchar(60)"`
	Type        ContentType `gorm:"type:TINYINT;NOT NULL"`
	Body        string      `gorm:"type:TEXT;NOT NULL"`
	BodyType    BodyType    `gorm:"type:TINYINT;NOT NULL;DEFAULT:1"`
	Version     int16       `gorm:"type:INT;NOT NULL"`
	CreatedAt   time.Time   `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time   `gorm:""`
	DeletedAt   *time.Time  `gorm:"" sql:"index"`
	Extra       string      `gorm:"type:TEXT;"`
}

type SDContentAudit struct {
	gorm.Model
	ID int64 `gorm:"AUTO_INCREMENT;NOT NULL"`
}
