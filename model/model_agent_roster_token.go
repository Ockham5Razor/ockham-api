package model

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AgentRosterToken struct {
	gorm.Model

	AgentId uint   `gorm:"index:idx_roster_token"`
	Token   string `gorm:"index:idx_roster_token"`
}

func b64uuid() string {
	u2, _ := uuid.NewV4()
	return base64.StdEncoding.EncodeToString(u2.Bytes())
}

func NewAgentRosterToken(agentId int) *AgentRosterToken {
	return &AgentRosterToken{
		AgentId: uint(agentId),
		Token:   b64uuid(),
	}
}
