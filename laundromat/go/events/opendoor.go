package events

import "github.com/thomaspeugeot/metabaron/libs/animah/go/models"

// OpenDoor is an event whose role is open the door
// of the machine
type OpenDoor struct {
	models.Event
}
