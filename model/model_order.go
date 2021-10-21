package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	Title            string `gorm:"type:VARCHAR(24)"`
	AutomaticRenewal bool
	UserID           uint
	User             User
	ServicePlanID    uint
	ServicePlan      ServicePlan
}
