package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR(24)"`
	Password string `gorm:"type:VARCHAR(128)"`
	Email    string `gorm:"type:VARCHAR(128)"`
}

type Role struct {
	gorm.Model
	RoleName string `gorm:"type:VARCHAR(24)"`
}

func CreateUser(username string, rawPassword string, email string) {

}
