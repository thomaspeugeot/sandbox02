package events

import "github.com/thomaspeugeot/sandbox02/animah/go/models"

// StartProgram is to start program
// machine goes from IDLE to RUNNING
type StartProgram struct {
	models.Event
}
