package models

import (
	"log"
	"time"
)

// AgentInterface is the interface that must be followed by part of the simulation
// swagger:model AgentInterface
type AgentInterface interface {

	// GetNextEvent provides the event with earliest fire time
	GetNextEvent() (event EventInterface, time time.Time)

	// FireNextEvent fire next Event
	FireNextEvent()

	setEngine(engine *Engine)
}

// Agent is the empty struct to perform
// generic agents chores
type Agent struct {
	TechName string

	// list of events, in increasing fire time
	// swagger:ignore
	events []EventInterface `gorm:"-"`

	// TimeLastChecked, time at which the event check state was called
	// swagger:ignore
	TimeLastChecked time.Time

	// last event time
	// swagger:ignore
	lastEventTime time.Time

	// usefull to append an agent to the Engine from an agent
	// swagger:ignore
	Engine *Engine
}

func (agent *Agent) setEngine(engine *Engine) {
	agent.Engine = engine
}

// AppendAgentToEngine append an agent to the engine
func (agent *Agent) AppendAgentToEngine(newAgent AgentInterface) {
	agent.Engine.AppendAgent(newAgent)
}

// RemoveAgentToEngine append an agent to the engine
func (agent *Agent) RemoveAgentToEngine(newAgent AgentInterface) {
	agent.Engine.RemoveAgent(newAgent)
}

// Reset removes all events from the agent and resets internal checks
func (agent *Agent) Reset() {

	agent.events = nil
	agent.lastEventTime = time.Time{}
}

// GetNextEvent provides the next event from a time point of view
// by convention 0 means infinity
func (agent *Agent) GetNextEvent() (EventInterface, time.Time) {

	if len(agent.events) == 0 {
		return nil, time.Time{}
	}

	return agent.events[0], agent.events[0].GetFireTime()
}

// GetNextEventAndRemoveIt provides the next event from a time point of view
// by convention 0 means infinity
func (agent *Agent) GetNextEventAndRemoveIt() (event EventInterface, t time.Time) {

	event, t = agent.events[0], agent.events[0].GetFireTime()
	if event == nil {
		log.Panic("cannot fire event when no event in queue")
	}

	// remove event
	agent.events = agent.events[1:len(agent.events)]

	//  update last time
	agent.lastEventTime = event.GetFireTime()
	return event, t
}

// QueueEvent is the function by which an agent queues an event from another agent (or of himself)
func (agent *Agent) QueueEvent(event EventInterface) {

	if event.GetFireTime().Before(agent.lastEventTime) {
		log.Panic("inserting event in the past")
	}

	if len(agent.events) == 0 {
		agent.events = append(agent.events, event)
		return
	}
	// parse all events and insert event when appropriate
	for idx, _event := range agent.events {
		if event.GetFireTime().Before(_event.GetFireTime()) {

			agent.events = append(agent.events, nil)       // Making space for the new element
			copy(agent.events[idx+1:], agent.events[idx:]) // Shifting elements
			agent.events[idx] = event                      // Copying/inserting the value
			return
		}
	}
	// append at the end
	agent.events = append(agent.events, event)

}
