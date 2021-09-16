package model

import (
	"gorm.io/gorm"
)

type ServicePlan struct {
	gorm.Model
	Title                    string `gorm:"type:VARCHAR(24)"`
	TotalTraffic             int64
	TrafficResetIntervalDays int16
}
