package diagrams

import (
	uml "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/tests/geometry/models"
)

var UmlscDiagram1 uml.Umlsc = uml.Umlsc{
	Activestate: string(models.APRES_CALCUL),
	States: []*uml.State{
		{
			X:    20.000000,
			Y:    90.000000,
			Name: string(models.APRES_CALCUL),
		},
		{
			X:    20.000000,
			Y:    30.000000,
			Name: string(models.AVANT_CALCUL),
		},
	},
}
