package model

import (
	"gorm.io/gorm"
	"ockham-api/config"
	"ockham-api/util"
	"time"
)

type Session struct {
	gorm.Model
	UserID       uint
	User         User
	SessionKey   string `gorm:"type:VARCHAR(36);uniqueIndex"`
	SessionToken string `gorm:"type:LONGTEXT"`
	ExpiredAt    time.Time
	RenewalStock int
}

var sessionExpireDuration = time.Second * time.Duration(config.AuthSessionExpireSeconds)

func CreateSession(user *User) *Session {
	key := util.GenString()
	body, _ := util.GenToken(user.Username, key)
	session := &Session{
		User:         *user,
		SessionKey:   key,
		SessionToken: body,
		ExpiredAt:    time.Now().Add(sessionExpireDuration),
		RenewalStock: config.AuthSessionMaximumRenewalTimes,
	}
	return session
}
