package events

import "github.com/thomaspeugeot/sandbox02/animah/go/models"

// OpenDoor is an event whose role is open the door
// of the machine
type OpenDoor struct {
	models.Event
}
