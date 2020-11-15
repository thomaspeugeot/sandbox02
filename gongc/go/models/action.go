package models

import (
	"github.com/jinzhu/gorm"
)

// ActionModel is a representation of a action of a metabaron Action
// a metabaron action
// swagger:model ActionModel
type ActionModel struct {

	// The Name of the Action
	Name string

	// The type of the action
	ActionType ActionType
}

// https://github.com/go-swagger/go-swagger/pull/2176

// swagger:enum ActionType
type ActionType string

// values for Action Type
const (
	WALK                    ActionType = "Walk" // iota // Parse the spinosa model (temp)
	DELETE_STRUCT_AND_FIELD ActionType = "DeleteStructAndFields"
)

// swagger:enum ActionTypeInt
type ActionTypeInt int

// values for Action Type
const (
	ACTION_TYPE_INT_0 ActionTypeInt = 0 // iota // Parse the spinosa model (temp)
	ACTION_TYPE_INT_1 ActionTypeInt = 1
)

// Action describres a action
// swagger:model ActionDB
type Action struct {
	gorm.Model

	ActionModel
}

// Actions arrays actions
// swagger:response actionsResponse
type Actions []Action

// ActionResponse provides response
// swagger:response actionResponse
type ActionResponse struct {
	Action
}
