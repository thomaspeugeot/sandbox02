package diagrams

import (
	uml "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/laundromat/go/models"
)

var Diagramme_UML_Laundromat uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Machine{}),
			Position: &uml.Position{
				X: 570.000000,
				Y: 100.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Machine{}.Name,
				},
				{
					Field: models.Machine{}.State,
				},
				{
					Field: models.Machine{}.Cleanedlaundry,
				},
				{
					Field: models.Machine{}.DrumLoad,
				},
				{
					Field: models.Machine{}.RemainingTime,
				},
				{
					Field: models.Machine{}.RemainingTimeMinutes,
				},
			},
			Width:  240.000000,
			Heigth: 48.000000,
		},
		{
			Struct: &(models.Washer{}),
			Position: &uml.Position{
				X: 60.000000,
				Y: 100.000000,
			},
			Width:  240.000000,
			Heigth: 48.000000,
			Fields: []*uml.Field{
				{
					Field: models.Washer{}.Name,
				},
				{
					Field: models.Washer{}.State,
				},
				{
					Field: models.Washer{}.LaundryWeight,
				},
			},
			Links: []*uml.Link{
				{
					Field: models.Washer{}.Machine,
					Middlevertice: &uml.Vertice{
						X: 420.000000,
						Y: 70.000000,
					},
				},
			},
		},
	},
}
