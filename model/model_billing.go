package model

import "gorm.io/gorm"

type Billing struct {
	gorm.Model

	Title                    string `gorm:"type:VARCHAR(24)"`
	BillingTotal             float32
	UserID                   uint
	User                     User
	OrderID                  uint
	Order                    Order
	ServicePlanUtilizationID uint
	ServicePlanUtilization   ServicePlanUtilization
}
