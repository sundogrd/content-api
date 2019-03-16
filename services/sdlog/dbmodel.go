package sdlog

import (
	"time"
)

type SDLogModel struct {
	ID        int64     `gorm:"primary_key;AUTO_INCREMENT;not null"`
	UserID    int64     `gorm:"not null;"`
	TargetID  int64     `gorm:"not null;"`
	Type      string    `gorm:"type:varchar(60);not null"`
	CreatedAt time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	Extra     string    `gorm:"type:TEXT;"`
}

func (SDLogModel) TableName() string {
	return "sd_logs"
}
