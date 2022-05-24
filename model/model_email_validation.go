package model

import (
	"crypto/rand"
	"fmt"
	"github.com/go-basic/uuid"
	"gorm.io/gorm"
	"ockham-api/config"
	"strings"
	"time"
)

type EmailValidation struct {
	gorm.Model
	UserID         uint
	User           User
	ValidationKey  string
	ValidationCode string
	ExpireAt       time.Time
}

func NewEmailValidation(user *User) *EmailValidation {
	expireDuration := config.EmailValidationExpireDuration
	duration, err := time.ParseDuration(expireDuration)
	if err != nil {
		panic(fmt.Sprintf("Email validation code expire duration error, config value is: '%v'!", expireDuration))
	}
	emailVerification := &EmailValidation{
		User:           *user,
		ValidationKey:  uuid.New(),
		ValidationCode: randomCode(),
		ExpireAt:       time.Now().Add(duration),
	}
	return emailVerification
}

func randomCode() string {
	c := 3
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	var builder strings.Builder
	for i := 0; i < len(b); i++ {
		builder.WriteString(fmt.Sprintf("%03d", b[i]))
	}
	return builder.String()
}
