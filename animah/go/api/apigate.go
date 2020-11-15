// generated by genGoApiGateApi.go
package api

import (
	"github.com/thomaspeugeot/metabaron/libs/animah/go/models"
)

// APIGateStruct is the implementation of the APIGateImplCallback interface
type APIGateStruct struct {
	MapActionIDAPI    map[*models.Action]uint // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIAction    map[uint]*models.Action // map of ID in the "api" namespace to ID in the "models" namespace
	MapActionAPIIDAPI map[*ActionAPI]uint     // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIActionAPI map[uint]*ActionAPI     // map of ID in the "api" namespace to ID in the "models" namespace

	// lastActionAPIID is the last unique identier that has been used.
	// It is initialized at 0, therefore IDAPI starts at 1
	lastActionAPIID uint

	MapAgentIDAPI    map[*models.Agent]uint // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIAgent    map[uint]*models.Agent // map of ID in the "api" namespace to ID in the "models" namespace
	MapAgentAPIIDAPI map[*AgentAPI]uint     // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIAgentAPI map[uint]*AgentAPI     // map of ID in the "api" namespace to ID in the "models" namespace

	// lastAgentAPIID is the last unique identier that has been used.
	// It is initialized at 0, therefore IDAPI starts at 1
	lastAgentAPIID uint

	MapEngineIDAPI    map[*models.Engine]uint // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIEngine    map[uint]*models.Engine // map of ID in the "api" namespace to ID in the "models" namespace
	MapEngineAPIIDAPI map[*EngineAPI]uint     // map of pointers of models instances to ID in the "api" namespace
	MapIDAPIEngineAPI map[uint]*EngineAPI     // map of ID in the "api" namespace to ID in the "models" namespace

	// lastEngineAPIID is the last unique identier that has been used.
	// It is initialized at 0, therefore IDAPI starts at 1
	lastEngineAPIID uint
}

// APIGate is the singloton gate in the "api" namespace
var APIGate APIGateStruct = APIGateStruct{
	MapActionIDAPI:    make(map[*models.Action]uint, 0),
	MapIDAPIAction:    make(map[uint]*models.Action, 0),
	MapActionAPIIDAPI: make(map[*ActionAPI]uint, 0),
	MapIDAPIActionAPI: make(map[uint]*ActionAPI, 0),

	MapAgentIDAPI:    make(map[*models.Agent]uint, 0),
	MapIDAPIAgent:    make(map[uint]*models.Agent, 0),
	MapAgentAPIIDAPI: make(map[*AgentAPI]uint, 0),
	MapIDAPIAgentAPI: make(map[uint]*AgentAPI, 0),

	MapEngineIDAPI:    make(map[*models.Engine]uint, 0),
	MapIDAPIEngine:    make(map[uint]*models.Engine, 0),
	MapEngineAPIIDAPI: make(map[*EngineAPI]uint, 0),
	MapIDAPIEngineAPI: make(map[uint]*EngineAPI, 0),
}