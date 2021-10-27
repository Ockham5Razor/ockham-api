package model

import (
	"gol-c/config"
	"gol-c/util"
	"gorm.io/gorm"
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

var sessionConfig = config.GetConfig().Auth.Session

var sessionExpireDuration = time.Second * sessionConfig.ExpireSeconds

func CreateSession(user *User) *Session {
	key := util.GenString()
	body, _ := util.GenToken(user.Username, key)
	session := &Session{
		User:         *user,
		SessionKey:   key,
		SessionToken: body,
		ExpiredAt:    time.Now().Add(sessionExpireDuration),
		RenewalStock: sessionConfig.MaximumRenewalTimes,
	}
	return session
}
