package model

import (
	"github.com/go-basic/uuid"
	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	UserID      uint
	User        User
	SessionBody string `gorm:"type:VARCHAR(36);uniqueIndex"`
}

func CreateSession(userID uint) *Session {
	sessionBody := uuid.New()
	session := &Session{
		UserID:      userID,
		SessionBody: sessionBody,
	}
	return session
}
