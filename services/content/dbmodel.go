package content

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BodyType int16
const (
	Body_TEXT BodyType = 1
	Body_HTML BodyType = 2
	Body_MARKDOWN BodyType = 3
	Body_LATEX BodyType = 4
)

// TODO: 使用bit表示的话枚举可能要改下实现方式，不然命名很蛋疼
type ContentType int16
const (
	Content_TEXT ContentType = 1
	Content_AUDIO ContentType = 2
	Content_VIDEO ContentType = 4
)

// SDContent http://gorm.io/docs/models.html
type SDContent struct {
	// gorm.Model
	ID          int64      `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID   int64      `gorm:"not null;"`
	Title       string     `gorm:"type:varchar(60);not null"`
	Description string     `gorm:"type:varchar(300);not null"`
	AuthorID    int64      `gorm:"not null;"`
	Category    string     `gorm:"type:varchar(60)"`
	Type        ContentType      `gorm:"type:TINYINT;NOT NULL"`
	Body        string     `gorm:"type:TEXT;NOT NULL"`
	BodyType    BodyType      `gorm:"type:TINYINT;NOT NULL;DEFAULT:1"`
	Version     int16      `gorm:"type:INT;NOT NULL"`
	CreatedAt   time.Time  `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time  `gorm:""`
	DeletedAt   *time.Time `gorm:"" sql:"index"`
	Extra       string     `gorm:"type:TEXT;"`
}

type SDContentAudit struct {
	gorm.Model
	ID int64 `gorm:"AUTO_INCREMENT;NOT NULL"`
}
