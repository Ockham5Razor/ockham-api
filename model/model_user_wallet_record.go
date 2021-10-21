package model

import "gorm.io/gorm"

type UserWalletRecord struct {
	gorm.Model
	Amount       float32 // 金额
	Description  string  `gorm:"type:LONGTEXT"` // 描述
	UserID       uint
	User         User
	UserWalletID uint
	UserWallet   UserWallet
}
