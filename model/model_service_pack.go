package model

import (
	"gorm.io/gorm"
	"time"
)

type ServicePack struct {
	gorm.Model

	Title            string `gorm:"type:VARCHAR(24)"`
	Description      string `gorm:"type:VARCHAR(255)"`
	ServiceStartTime time.Time
	ServiceEndTime   time.Time
	UserID           int
	User             User
	Enabled          bool
}
