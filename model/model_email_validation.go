package model

import (
	"gorm.io/gorm"
)

type EmailValidation struct {
	gorm.Model
	UserID        uint
	User          User
	ValidatorKey  string
	ValidatorCode string
}
