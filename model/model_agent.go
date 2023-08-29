package model

import (
	"gorm.io/gorm"
	"time"
)

type Agent interface {
	IsDead() bool
}

type VmessAgent struct {
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

func (a *VmessAgent) IsDead() bool {
	return a.LastPulse.Add(1 * time.Minute).After(time.Now())
}

// VlessAgent TODO
type VlessAgent struct {
}

func (a *VlessAgent) IsDead() bool {
	return false
}

// TrojanAgent TODO
type TrojanAgent struct {
}

func (a *TrojanAgent) IsDead() bool {
	return false
}

// ShadowsocksAgent TODO
type ShadowsocksAgent struct {
}

func (a *ShadowsocksAgent) IsDead() bool {
	return false
}
