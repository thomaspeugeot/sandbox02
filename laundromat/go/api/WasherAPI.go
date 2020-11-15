// generated by genGoStructApi.go
package api

import (
	"errors"

	"github.com/thomaspeugeot/metabaron/examples/laundromat/go/models"
)

// WasherAPI is the twin in the "api" namespace of Washer of the "models" namespace
//
// swagger:model WasherAPI
type WasherAPI struct {
	models.Washer

	IDAPI uint // unique ID of the instance in the "api" namespace

	// ID generated for the implementation of the field Washer{}.Machine *Washer
	MachineIDAPI uint
}

// CreateAPIWasher creates from washer an instance in the "api" namespace
// CreateAPIWasher performs a deep copy of washer fields
func (aPIGate *APIGateStruct) CreateAPIWasher(washer *models.Washer) (IDAPI uint, err error) {

	washerAPI := new(WasherAPI)

	IDAPI = aPIGate.lastWasherAPIID + 1
	aPIGate.lastWasherAPIID = IDAPI

	// update store
	aPIGate.MapWasherIDAPI[washer] = IDAPI
	aPIGate.MapIDAPIWasher[IDAPI] = washer
	aPIGate.MapWasherAPIIDAPI[washerAPI] = IDAPI
	aPIGate.MapIDAPIWasherAPI[IDAPI] = washerAPI
	return
}

// UpdateAPIWasher updates from washer an instance in the "api" namespace
// UpdateAPIWasher performs a deep copy of washer fields
func (aPIGate *APIGateStruct) UpdateAPIWasher(washer *models.Washer) (IDAPI uint, err error) {

	// check if twin ID exists
	var ok bool
	IDAPI, ok = aPIGate.MapWasherIDAPI[washer]
	if !ok {
		return uint(0), errors.New("unknown washer")
	}

	// get the twin API
	washerAPI := aPIGate.MapIDAPIWasherAPI[IDAPI]

	// update values of washerAPI with a deep copy
	washerAPI.Washer = *washer

	// set MachineID
	if washer.Machine != nil {
		if machineIDAPI, ok := aPIGate.MapMachineIDAPI[washer.Machine]; ok {
			washerAPI.MachineIDAPI = machineIDAPI
		} else {
			return 0, errors.New("Washer : Unkown Machine with association Machine")
		}
	}

	return
}
