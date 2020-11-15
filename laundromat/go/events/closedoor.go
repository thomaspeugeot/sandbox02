package events

import "github.com/thomaspeugeot/sandbox02/animah/go/models"

// CloseDoor is an event whose role is close the door
// of the machine
type CloseDoor struct {
	models.Event
}
