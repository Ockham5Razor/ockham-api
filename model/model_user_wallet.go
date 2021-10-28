package model

import (
	"gorm.io/gorm"
	"time"
)

type UserWallet struct {
	gorm.Model `json:"-"`
	Balance    float32 `json:"balance"` // 余额
	UserID     uint    `json:"-"`
	User       User    `json:"-"`
}

type UserWalletRecord struct {
	gorm.Model   `json:"-"`
	Amount       float32    `json:"amount"`                           // 金额
	Description  string     `json:"description" gorm:"type:LONGTEXT"` // 描述
	UserID       uint       `json:"-"`
	User         User       `json:"-"`
	UserWalletID uint       `json:"-"`
	UserWallet   UserWallet `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
