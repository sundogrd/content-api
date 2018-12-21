package content

import (
	"time"

	"github.com/jinzhu/gorm"
)

// PFContent http://gorm.io/docs/models.html
type PFContent struct {
	// gorm.Model
	ID          int64      `gorm:"primary_key;AUTO_INCREMENT;not null"`
	ContentID   int64      `gorm:"not null;"`
	Title       string     `gorm:"type:varchar(60);not null"`
	Description string     `gorm:"type:varchar(300);not null"`
	AuthorID    int64      `gorm:"not null;"`
	Category    string     `gorm:"type:varchar(60)"`
	Type        int16      `gorm:"type:TINYINT;NOT NULL"`
	Body        string     `gorm:"type:TEXT;NOT NULL"`
	Version     int16      `gorm:"type:INT;NOT NULL"`
	CreatedAt   time.Time  `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time  `gorm:""`
	DeletedAt   *time.Time `gorm:"" sql:"index"`
	Extra       string     `gorm:"type:TEXT;"`
}

type PFContentAudit struct {
	gorm.Model
	ID int64 `gorm:"AUTO_INCREMENT;NOT NULL"`
}
