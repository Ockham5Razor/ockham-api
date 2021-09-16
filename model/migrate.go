package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	models := []interface{}{
		&User{},
		&ServicePack{},
		&Commodity{},
		&ServicePlan{},
	}

	for i := 0; i < len(models); i++ {
		err := db.AutoMigrate(models[i])
		if err != nil {
			panic("Failed to migrate table.")
		}
	}
}
