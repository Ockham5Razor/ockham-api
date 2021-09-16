package model

import (
	"gorm.io/gorm"
	"time"
)

type ServicePack struct {
	gorm.Model
	Title       string    `gorm:"type:VARCHAR(24)"`
	Description string    `gorm:"type:VARCHAR(255)"`
	DueTime     time.Time `gorm:"type:VARCHAR(24)"`
	UserID      int
	User        User
}
