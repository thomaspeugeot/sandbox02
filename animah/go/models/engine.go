package models

import (
	"log"
	"time"
)

// EngineSingloton is the default engine
var EngineSingloton Engine

// Engine describes a tiny discrete event simulation engine
// it is responsible for advancing the time
// An engine manages a set of agents
// swagger:model Engine
type Engine struct {
	// Name of the engine "laundramat" for instance
	Name string

	// StartTime is the simulation start date
	StartTime time.Time

	// EndTime is the simulatio end date
	EndTime time.Time

	// CurrentTime is the simulation current time
	CurrentTime time.Time

	// list of engine agents
	agents []AgentInterface `gorm:"-"`

	// Fired events
	Fired int

	// control mode.
	ControlMode ControlMode

	// engine state
	State EngineState

	// LastEvent ...
	LastEvent *EventInterface `gorm:"-"`

	// LastEvent agent
	LastEventAgent *AgentInterface `gorm:"-"`

	// EngineSpecificInteface supportspecific callback
	// on the engine events
	EngineSpecificInteface EngineSpecificInterface `gorm:"-"`

	// Speed compared to realtime
	Speed float64
}

// swagger:enum ControlMode
type ControlMode string

// values for ControlMode
const (
	AUTONOMOUS     ControlMode = "Autonomous" // iota
	CLIENT_CONTROL ControlMode = "ClientControl"
)

// swagger:enum EngineState
type EngineState string

// values for EngineState
const (
	RUNNING EngineState = "RUNNING" // iota
	PAUSED  EngineState = "PAUSED"
	OVER    EngineState = "OVER"
)

// AppendAgent to the engine
func (engine *Engine) AppendAgent(agent AgentInterface) {
	engine.agents = append(engine.agents, agent)
	agent.setEngine(engine)
}

// RemoveAgent to the engine
func (engine *Engine) RemoveAgent(agent AgentInterface) {
	for index, _agent := range engine.agents {

		// Order is not important
		// If you do not care about ordering, you have the much faster
		// possibility to swap the element to delete with the one at the end of the slice and then return the n-1 first elements:
		// https://stackoverflow.com/a/37335777/5803707
		if _agent == agent {
			engine.agents[index] = engine.agents[len(engine.agents)-1]
			engine.agents = engine.agents[:len(engine.agents)-1]
		}
	}
}

// GetNextEvent ...
func (engine *Engine) GetNextEvent() (agent AgentInterface, nextEventFireTime time.Time, event EventInterface) {

	firstAgent := true
	for _, _agent := range engine.agents {
		_event, agentNextEventFireTime := _agent.GetNextEvent()

		if firstAgent || agentNextEventFireTime.Before(nextEventFireTime) {
			nextEventFireTime = agentNextEventFireTime
			agent = _agent
			event = _event
		}
		firstAgent = false
	}
	return agent, nextEventFireTime, event
}

// FireNextEvent fires earliest event
// advances current time
func (engine *Engine) FireNextEvent() (agent AgentInterface, nextTimeEvent time.Time, event EventInterface) {

	agent, nextTimeEvent, event = engine.GetNextEvent()

	agent.FireNextEvent()

	engine.LastEvent = &event
	engine.LastEventAgent = &agent
	engine.CurrentTime = nextTimeEvent

	if engine.EngineSpecificInteface != nil {
		engine.EngineSpecificInteface.EventFired(engine)
	}

	return agent, nextTimeEvent, event
}

// Run will advance time till currentTime > EndTime
func (engine *Engine) Run() {
	log.Printf("time : run\n")

	_, time, _ := engine.GetNextEvent()

	for (!time.IsZero()) && time.Before(engine.EndTime) {
		engine.FireNextEvent()
		_, time, _ = engine.GetNextEvent()
		engine.Fired++
	}
}

// RunTillAnyStateHasChanged will advance time till currentTime > EndTime or one state Changed in the implementation
func (engine *Engine) RunTillAnyStateHasChanged() {
	log.Printf("time : run\n")

	_, time, _ := engine.GetNextEvent()

	hasAnyStateHasChanged := false

	for (!time.IsZero()) && time.Before(engine.EndTime) && !hasAnyStateHasChanged {
		engine.FireNextEvent()
		_, time, _ = engine.GetNextEvent()
		engine.Fired++
		if engine.EngineSpecificInteface != nil {
			hasAnyStateHasChanged = engine.EngineSpecificInteface.HasAnyStateChanged(engine)
		}
	}
}
