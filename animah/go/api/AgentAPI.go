// generated by genGoStructApi.go
package api

import (
	"errors"

	"github.com/thomaspeugeot/sandbox02/animah/go/models"
)

// AgentAPI is the twin in the "api" namespace of Agent of the "models" namespace
//
// swagger:model AgentAPI
type AgentAPI struct {
	models.Agent

	IDAPI uint // unique ID of the instance in the "api" namespace

	// ID generated for the implementation of the field Agent{}.Engine *Agent
	EngineIDAPI uint
}

// CreateAPIAgent creates from agent an instance in the "api" namespace
// CreateAPIAgent performs a deep copy of agent fields
func (aPIGate *APIGateStruct) CreateAPIAgent(agent *models.Agent) (IDAPI uint, err error) {

	agentAPI := new(AgentAPI)

	IDAPI = aPIGate.lastAgentAPIID + 1
	aPIGate.lastAgentAPIID = IDAPI

	// update store
	aPIGate.MapAgentIDAPI[agent] = IDAPI
	aPIGate.MapIDAPIAgent[IDAPI] = agent
	aPIGate.MapAgentAPIIDAPI[agentAPI] = IDAPI
	aPIGate.MapIDAPIAgentAPI[IDAPI] = agentAPI
	return
}

// UpdateAPIAgent updates from agent an instance in the "api" namespace
// UpdateAPIAgent performs a deep copy of agent fields
func (aPIGate *APIGateStruct) UpdateAPIAgent(agent *models.Agent) (IDAPI uint, err error) {

	// check if twin ID exists
	var ok bool
	IDAPI, ok = aPIGate.MapAgentIDAPI[agent]
	if !ok {
		return uint(0), errors.New("unknown agent")
	}

	// get the twin API
	agentAPI := aPIGate.MapIDAPIAgentAPI[IDAPI]

	// update values of agentAPI with a deep copy
	agentAPI.Agent = *agent

	// set EngineID
	if agent.Engine != nil {
		if engineIDAPI, ok := aPIGate.MapEngineIDAPI[agent.Engine]; ok {
			agentAPI.EngineIDAPI = engineIDAPI
		} else {
			return 0, errors.New("Agent : Unkown Engine with association Engine")
		}
	}

	return
}
