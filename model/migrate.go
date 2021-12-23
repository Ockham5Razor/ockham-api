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
		&ServicePlanSubscription{},
		&Session{},
		&EmailValidation{},
		&RechargeCode{},
		&Billing{},
	}

	for i := 0; i < len(models); i++ {
		err := db.AutoMigrate(models[i])
		if err != nil {
			panic("Failed to migrate table.")
		}
	}
}

func InitData(db *gorm.DB) {
	initialRole0 := &Role{RoleName: "admin"}
	initialRole1 := &Role{RoleName: "user"}
	db.FirstOrCreate(initialRole0, Role{RoleName: "admin"})
	db.FirstOrCreate(initialRole1, Role{RoleName: "user"})
	initialUser := &User{
		Username:      "admin",
		Password:      util.Encrypt("admin"),
		Email:         "dave.smith@admin.com",
		EmailVerified: true,
		Roles:         []*Role{initialRole0, initialRole1},
	}
	db.FirstOrCreate(initialUser, User{Username: "admin"})
}
