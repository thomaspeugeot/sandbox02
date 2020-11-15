package models

// EngineSpecificInterface is the callback support for
// events that happens on the generic engine
type EngineSpecificInterface interface {
	EventFired(engine *Engine)

	// the specific engine shall implement this callback
	// it returns true if one state of the specific has changed
	HasAnyStateChanged(engine *Engine) bool
}
