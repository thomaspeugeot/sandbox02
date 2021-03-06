package diagrams

import (
	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	uml "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
)

var DiagramUMLduUML uml.Classdiagram = uml.Classdiagram{
	Classshapes: []*uml.Classshape{
		{
			Struct: &(models.Classdiagram{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 290.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Classdiagram{}.Classshapes,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 420.000000,
					},
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Classdiagram{}.Name,
				},
			},
		},
		{
			Struct: &(models.Classshape{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 490.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Classshape{}.Fields,
					Middlevertice: &uml.Vertice{
						X: 550.000000,
						Y: 630.000000,
					},
				},
				{
					Field: models.Classshape{}.Links,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 620.000000,
					},
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Classshape{}.Structname,
				},
			},
		},
		{
			Struct: &(models.Field{}),
			Position: &uml.Position{
				X: 430.000000,
				Y: 680.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Field{}.Fieldname,
				},
			},
		},
		{
			Struct: &(models.Link{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 680.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Link{}.Middlevertice,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 790.000000,
					},
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Link{}.Fieldname,
				},
			},
		},
		{
			Struct: &(models.Pkgelt{}),
			Position: &uml.Position{
				X: 350.000000,
				Y: 90.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Pkgelt{}.Classdiagrams,
					Middlevertice: &uml.Vertice{
						X: 200.000000,
						Y: 170.000000,
					},
				},
				{
					Field: models.Pkgelt{}.Umlscs,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 180.000000,
					},
				},
			},
		},
		{
			Struct: &(models.State{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 480.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.State{}.Name,
				},
			},
		},
		{
			Struct: &(models.Umlsc{}),
			Position: &uml.Position{
				X: 580.000000,
				Y: 290.000000,
			},
			Links: []*uml.Link{
				{
					Field: models.Umlsc{}.States,
					Middlevertice: &uml.Vertice{
						X: 700.000000,
						Y: 420.000000,
					},
				},
			},
			Fields: []*uml.Field{
				{
					Field: models.Umlsc{}.Name,
				},
			},
		},
		{
			Struct: &(models.Vertice{}),
			Position: &uml.Position{
				X: 80.000000,
				Y: 850.000000,
			},
			Fields: []*uml.Field{
				{
					Field: models.Vertice{}.X,
				},
				{
					Field: models.Vertice{}.Y,
				},
			},
		},
	},
}
