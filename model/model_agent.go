package model

import (
	"gorm.io/gorm"
	"time"
)

type Agent struct {
	gorm.Model

	Name        string
	Description string
	Location    string

	ServerHost string
	ServerPort int
	ExternalID int
	ClientPort int
	WsHost     string
	WsPath     string

	Enabled bool

	LastPulse       time.Time
	SecretAccessKey string
}

func (a *Agent) IsDead() bool {
	return a.LastPulse.Add(1 * time.Minute).After(time.Now())
}
