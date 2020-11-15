package diagrams

import (
	uml "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/tests/geometry/models"
)

var Diagram2 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 10.000000,
				Y: 150.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 740.000000,
						Y: 300.000000,
					},
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 740.000000,
						Y: 40.000000,
					},
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 620.000000,
				Y: 140.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
	},
}
