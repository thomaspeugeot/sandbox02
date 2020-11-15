package diagrams

import (
	uml "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/tests/geometry/models"
)

var Diagram1 uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Line{}),
			Position: &uml.Position{
				X: 20.000000,
				Y: 80.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Links: []*uml.Link{
				{
					Field: models.Line{}.End,
					Middlevertice: &uml.Vertice{
						X: 380.000000,
						Y: 250.000000,
					},
				},
				{
					Field: models.Line{}.Start,
					Middlevertice: &uml.Vertice{
						X: 360.000000,
						Y: 10.000000,
					},
				},
			},
		},
		{
			Struct: &(models.Point{}),
			Position: &uml.Position{
				X: 500.000000,
				Y: 90.000000,
			},
			Width:  240.000000,
			Heigth: 78.000000,
			Fields: []*uml.Field{
				{
					Field: models.Point{}.X,
				},
				{
					Field: models.Point{}.Y,
				},
			},
		},
	},
}
