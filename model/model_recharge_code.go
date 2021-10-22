package model

import (
	"database/sql"
	"gol-c/utils"
	"gorm.io/gorm"
)

type RechargeCode struct {
	gorm.Model
	PackageName    string `gorm:"type:VARCHAR(36)"`
	RechargeCode   string `gorm:"type:VARCHAR(36);uniqueIndex"`
	RechargeAmount float32
	Used           bool
	ExportedAt     sql.NullTime
}

func GenRechargeCode(packageName string, rechargeAmount float32) *RechargeCode {
	return &RechargeCode{
		PackageName:    packageName,
		RechargeCode:   utils.GenString(),
		RechargeAmount: rechargeAmount,
		Used:           false,
	}
}
