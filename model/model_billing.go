package model

import (
	"gorm.io/gorm"
	"time"
)

type Billing struct {
	gorm.Model

	Title                     string `gorm:"type:VARCHAR(24)"`
	BillingTotal              float32
	BillingDate               time.Time
	DueDate                   time.Time
	Paid                      bool
	UserID                    uint
	User                      User
	ServicePlanSubscriptionID uint
	ServicePlanSubscription   ServicePlanSubscription
	SplitBillings             []SplitBilling
}

type SplitBilling struct {
	gorm.Model

	Title                     string `gorm:"type:VARCHAR(24)"`
	BillingTotal              float32
	BillingDate               time.Time
	DueDate                   time.Time
	Paid                      bool
	UserID                    uint
	User                      User
	Billing                   Billing
	BillingID                 uint
	ServicePackID             uint
	ServicePack               ServicePack
	ServicePlanSubscriptionID uint
	ServicePlanSubscription   ServicePlanSubscription
}
