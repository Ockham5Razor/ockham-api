package model

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type AgentRosterToken struct {
	gorm.Model

	AgentId uint
	Token   string
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
