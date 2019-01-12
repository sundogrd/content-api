package user

import (
	"time"
)

// SDContent http://gorm.io/docs/models.html
type SDUser struct {
	// gorm.Model
	ID          int64      `gorm:"primary_key;AUTO_INCREMENT;not null"`
	UserID      int64      `gorm:"not null;"`
	Name        string     `gorm:"type:varchar(60);not null"`
	AvatarUrl   string     `gorm:"type:varchar(300);not null"`
	Company     string     `gorm:"type:varchar(60)"`
	Email       string     `gorm:"type:varchar(60)"`
	CreatedAt   time.Time  `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time  `gorm:""`
	DeletedAt   *time.Time `gorm:"" sql:"index"`
	Extra       string     `gorm:"type:TEXT;"`
}