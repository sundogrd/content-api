package content

import (
	"time"

	"github.com/jinzhu/gorm"
)

//BodyType ...
type BodyType int16

// ...
const (
	BodyTypeText     BodyType = 1
	BodyTypeHTML     BodyType = 2
	BodyTypeMarkdown BodyType = 3
	BodyTypeLatex    BodyType = 4
)

// ContentType ...
type ContentType int16

// ...
const (
	TypeText  ContentType = 1
	TypeAudio ContentType = 2
	TypeVideo ContentType = 4
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

func (SDContent) TableName() string {
	return "sd_contents"
}


// SDContentAudit ...
type SDContentAudit struct {
	gorm.Model
	ID int64 `gorm:"AUTO_INCREMENT;NOT NULL"`
}
func (SDContentAudit) TableName() string {
	return "sd_content_audits"
}


// SDContentCount ...
type SDContentCount struct {
	ID         int64  `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID  int64  `gorm:"not null;"`
	CountKey   string `gorm:"type:varchar(60);not null"`
	CountValue int64  `gorm:"type:bigint;not null"`
}
func (SDContentCount) TableName() string {
	return "sd_content_counts"
}



