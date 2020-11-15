package events

import "github.com/thomaspeugeot/metabaron/libs/animah/go/models"

// StartProgram is to start program
// machine goes from IDLE to RUNNING
type StartProgram struct {
	models.Event
}
