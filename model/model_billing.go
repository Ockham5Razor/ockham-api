package model

import (
	"gorm.io/gorm"
	"time"
)

type Billing struct {
	gorm.Model

	Title         string `gorm:"type:VARCHAR(24)"`
	Description   string `gorm:"type:LONGTEXT"`
	BillingTotal  float32
	BillingDate   time.Time
	DueDate       time.Time
	Paid          bool
	UserID        uint
	User          User
	SplitPayment  bool
	SplitBillings []*Billing `gorm:"many2many:split_billings"`
}
