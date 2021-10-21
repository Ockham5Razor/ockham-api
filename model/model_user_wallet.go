package model

import "gorm.io/gorm"

type UserWallet struct {
	gorm.Model
	Balance float32 // 余额
	UserID  uint
	User    User
}
