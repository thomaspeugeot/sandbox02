package uml

// ClassDiagram diagram struct store a class diagram
// temporary here
type ClassDiagram struct {
	Name       string
	ClassShape []*ClassShape
}

type ClassShape struct {
	Position Position
	Struct   interface{}
}

type Position struct {
	X, Y float64
	Name string // temporary
}
