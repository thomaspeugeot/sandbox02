package models

import (
	"log"
	"time"
)

// ActionSingloton catch callback
// swagger:ignore
type ActionSingloton struct {
}

// PostAction is called from the Action
func (actionSingloton *ActionSingloton) PostAction(action *Action) {

	log.Printf("Post action called with %s", action.Name)

	switch action.Name {
	case string(PLAY):
		EngineSingloton.State = RUNNING

		_, simTime, _ := EngineSingloton.GetNextEvent()

		// display state for the first time
		for simTime.Before(EngineSingloton.EndTime) && EngineSingloton.State == RUNNING {

			_, newSimTime, _ := EngineSingloton.FireNextEvent()

			if newSimTime.Sub(simTime) > 0 {
				simTimeAdvance := newSimTime.Sub(simTime)

				time.Sleep(time.Duration(float64(simTimeAdvance) / EngineSingloton.Speed))
			}
			simTime = newSimTime

			EngineSingloton.Fired++
		}

		if !simTime.Before(EngineSingloton.EndTime) {
			EngineSingloton.State = OVER
		}
	case string(FIRE_EVENT_TILL_STATES_CHANGE):
		EngineSingloton.State = RUNNING

		_, simTime, _ := EngineSingloton.GetNextEvent()
		hasAnyStateHasChanged := false

		// display state for the first time
		for simTime.Before(EngineSingloton.EndTime) &&
			EngineSingloton.State == RUNNING &&
			!hasAnyStateHasChanged {

			_, newSimTime, _ := EngineSingloton.FireNextEvent()
			simTime = newSimTime

			EngineSingloton.Fired++
			if EngineSingloton.EngineSpecificInteface != nil {
				hasAnyStateHasChanged = EngineSingloton.EngineSpecificInteface.HasAnyStateChanged(&EngineSingloton)
			}
		}

		EngineSingloton.State = PAUSED
		if !simTime.Before(EngineSingloton.EndTime) {
			EngineSingloton.State = OVER
		}

		// to update display
		if EngineSingloton.EngineSpecificInteface != nil {
			EngineSingloton.EngineSpecificInteface.EventFired(&EngineSingloton)
		}

	case string(PAUSE):
		EngineSingloton.State = PAUSED

		// to update display
		if EngineSingloton.EngineSpecificInteface != nil {
			EngineSingloton.EngineSpecificInteface.EventFired(&EngineSingloton)
		}

	case string(FIRE_NEXT_EVENT):
		_, simTime, _ := EngineSingloton.GetNextEvent()

		if simTime.Before(EngineSingloton.EndTime) {
			_, simTime, _ = EngineSingloton.FireNextEvent()
			EngineSingloton.Fired++
		}

		if !simTime.Before(EngineSingloton.EndTime) {
			EngineSingloton.State = OVER
		}

		// to update display
		if EngineSingloton.EngineSpecificInteface != nil {
			EngineSingloton.EngineSpecificInteface.EventFired(&EngineSingloton)
		}

	case string(RESET):
	case string(INCREASE_SPEED_100_PERCENTS):

		EngineSingloton.Speed *= 2.0

		// to update display
		if EngineSingloton.EngineSpecificInteface != nil {
			EngineSingloton.EngineSpecificInteface.EventFired(&EngineSingloton)
		}
	case string(DECREASE_SPEED_50_PERCENTS):
		EngineSingloton.Speed *= 0.5

		// to update display
		if EngineSingloton.EngineSpecificInteface != nil {
			EngineSingloton.EngineSpecificInteface.EventFired(&EngineSingloton)
		}

	default:
		log.Panic("unkwn action " + action.Name)
	}
}
