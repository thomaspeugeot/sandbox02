package diagrams

import (
	"github.com/thomaspeugeot/metabaron/examples/laundromat/go/models"
	uml "github.com/thomaspeugeot/metabaron/libs/gorgo/go/models"
)

var states_Machine uml.Umlsc = uml.Umlsc{
	Activestate: string(models.MACHINE_DOOR_OPEN),
	States: []*uml.State{
		{
			X:    10.000000,
			Y:    10.000000,
			Name: string(models.MACHINE_DOOR_OPEN),
		},
		{
			X:    10.000000,
			Y:    60.000000,
			Name: string(models.MACHINE_DOOR_CLOSED_IDLE),
		},
		{
			X:    10.000000,
			Y:    110.000000,
			Name: string(models.MACHINE_DOOR_CLOSED_RUNNING),
		},
	},
}
