package model

import (
	"fmt"
	"gorm.io/gorm"
	"ockham-api/api/v1/util"
)

func Migrate(db *gorm.DB) {
	models := []interface{}{
		&Role{},
		&User{},
		&UserWallet{},
		&UserWalletRecord{},
		&ServicePlan{},
		&TrafficPlan{},
		&ServicePlanSubscription{},
		&TrafficPlanSubscription{},
		&Session{},
		&EmailValidation{},
		&RechargeCode{},
		&Billing{},
		&VmessAgent{},
		&AgentRosterToken{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(fmt.Sprintf("Failed to migrate table, model is: %#v", model))
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
