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

	Host       string
	ServerPort int64
	ClientPort int64
	ExternalID int8
	ListenHost string
	GrpcHost   string

	Enabled bool

	LastPulse       time.Time
	SecretAccessKey string
}
