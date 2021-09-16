package model

import (
	"gorm.io/gorm"
	"time"
)

type Commodity struct {
	gorm.Model
	Name    string    `gorm:"type:VARCHAR(24)"`
	Stock   int32     `gorm:"type:BIGINT"`
	Price   int32     `gorm:"type:BIGINT"`
	DueTime time.Time `gorm:"type:VARCHAR(24)"`
}
