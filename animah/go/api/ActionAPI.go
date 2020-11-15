// generated by genGoStructApi.go
package api

import (
	"errors"

	"github.com/thomaspeugeot/sandbox02/animah/go/models"
)

// ActionAPI is the twin in the "api" namespace of Action of the "models" namespace
//
// swagger:model ActionAPI
type ActionAPI struct {
	models.Action

	IDAPI uint // unique ID of the instance in the "api" namespace
}

// CreateAPIAction creates from action an instance in the "api" namespace
// CreateAPIAction performs a deep copy of action fields
func (aPIGate *APIGateStruct) CreateAPIAction(action *models.Action) (IDAPI uint, err error) {

	actionAPI := new(ActionAPI)

	IDAPI = aPIGate.lastActionAPIID + 1
	aPIGate.lastActionAPIID = IDAPI

	// update store
	aPIGate.MapActionIDAPI[action] = IDAPI
	aPIGate.MapIDAPIAction[IDAPI] = action
	aPIGate.MapActionAPIIDAPI[actionAPI] = IDAPI
	aPIGate.MapIDAPIActionAPI[IDAPI] = actionAPI
	return
}

// UpdateAPIAction updates from action an instance in the "api" namespace
// UpdateAPIAction performs a deep copy of action fields
func (aPIGate *APIGateStruct) UpdateAPIAction(action *models.Action) (IDAPI uint, err error) {

	// check if twin ID exists
	var ok bool
	IDAPI, ok = aPIGate.MapActionIDAPI[action]
	if !ok {
		return uint(0), errors.New("unknown action")
	}

	// get the twin API
	actionAPI := aPIGate.MapIDAPIActionAPI[IDAPI]

	// update values of actionAPI with a deep copy
	actionAPI.Action = *action

	return
}
