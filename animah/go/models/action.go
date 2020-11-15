package models

// swagger:enum ActionType
type ActionType string

// values for EnumType
const (
	PLAY                          ActionType = "PLAY"
	PAUSE                         ActionType = "PAUSE"
	FIRE_NEXT_EVENT               ActionType = "FIRE_NEXT_EVENT"
	FIRE_EVENT_TILL_STATES_CHANGE ActionType = "FIRE_EVENT_TILL_STATES_CHANGE"
	RESET                         ActionType = "RESET"
	INCREASE_SPEED_100_PERCENTS   ActionType = "INCREASE_SPEED_100_PERCENTS"
	DECREASE_SPEED_50_PERCENTS    ActionType = "DECREASE_SPEED_50_PERCENTS"
)

// Action to control the engine
// swagger:model Action
type Action struct {
	Name string
}
