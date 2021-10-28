package model

import (
	"gol-c/api/v1/util"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	models := []interface{}{
		&Role{},
		&User{},
		&UserWallet{},
		&UserWalletRecord{},
		&ServicePack{},
		&ServicePlan{},
		&ServicePlanUtilization{},
		&Session{},
		&EmailValidation{},
		&Order{},
		&RechargeCode{},
	}

	for i := 0; i < len(models); i++ {
		err := db.AutoMigrate(models[i])
		if err != nil {
			panic("Failed to migrate table.")
		}
	}
}

func InitData(db *gorm.DB) {
	initialRole := &Role{RoleName: "admin"}
	db.FirstOrCreate(initialRole, &Role{RoleName: "admin"})
	initialUser := &User{
		Username:      "admin",
		Password:      util.Encrypt("admin"),
		Email:         "dave.smith@admin.com",
		EmailVerified: true,
		Roles:         []Role{*initialRole},
	}
	db.FirstOrCreate(initialUser, &User{Username: "admin"})
}
