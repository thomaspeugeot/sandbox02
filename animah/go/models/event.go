package models

import (
	"time"
)

// Event is the elementary element of a discrete event simulation
// swagger:ignore
type Event struct {
	Name string

	// Fire is the time at which the event is Fired
	FireTime time.Time

	// Duration is the difference between the current time and the fire time of the
	// event. It is handy to compute directly the fire time
	Duration time.Duration
}

// GetFireTime ...
func (event Event) GetFireTime() time.Time {
	return event.FireTime
}

// GetDuration ...
func (event Event) GetDuration() time.Duration {
	return event.Duration
}

func (event Event) SetFireTime(t time.Time) {
	event.FireTime = t
}

// EventInterface ...
// swagger:ignore
type EventInterface interface {
	GetFireTime() time.Time
	SetFireTime(time.Time)
	GetDuration() time.Duration
	GetName() string
	GetEvent() Event
}

// GetName ..
func (event Event) GetName() string { return event.Name }

func (event Event) GetEvent() Event { return event }
