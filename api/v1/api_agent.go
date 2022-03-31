package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"ockham-api/api/v1/util"
	"ockham-api/database"
	"ockham-api/model"
	"strconv"
)

type CreateAgentForm struct {
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
}

func CreateAgent(c *gin.Context) {
	createAgentForm := &CreateAgentForm{}
	util.FillJsonForm(c, createAgentForm)
}

func ListAgents(c *gin.Context) {

}

func GetAgent(c *gin.Context) {

}

func UpdateAgent(c *gin.Context) {

}

func DeleteAgent(c *gin.Context) {

}

// AgentPulse
// @Summary AgentPulse
// @SubscriptionDescription Send agent pulse to api server.
// @Tags agent
// @Success 200 {object} util.Pack
// @Failure 500 {object} util.Pack
// @Router /v1/agents/{agent_id}/pulse [PUT]
func AgentPulse(c *gin.Context) {
	agentKeyObj, exists := c.Get(util.ContextSignatureValue)
	if !exists {
		panic(util.ContextSignatureValue + " not exists in context!")
	}
	fmt.Println("agentKeyObj:", agentKeyObj)
}

// GetAgentSecretKey get agent secret key
func GetAgentSecretKey(resourceIdStr string) (string, error) {
	agentId, err := strconv.Atoi(resourceIdStr)
	if err != nil {
		return "", errors.New("agent_id needs to be integer")
	}

	agent := &model.Agent{}
	database.DBConn.Find(agent, agentId)
	// TODO bad way to check if agent exists.
	if agent.Name == "" {
		return "", errors.New("agent not found")
	}

	return agent.SecretAccessKey, nil
}

/*
curl -X 'PUT' \
  'http://localhost:8080/api/v1/agents/123/pulse' \
  -H 'Authorization: Signature 777888' \
  -H 'X-Timestamp: 1234567890' \
*/
