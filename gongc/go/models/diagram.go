package models

import (
	"github.com/jinzhu/gorm"
)

// DiagramModel is a representation of a diagram
// swagger:model DiagramModel
type DiagramModel struct {

	// The Name of the Struct Diagram
	Name string

	// The Storage as a json string
	Storage string
}

// Diagram describres a diagram
// swagger:model DiagramDB
type Diagram struct {
	gorm.Model

	DiagramModel
}

// Diagrams arrays diagrams
// swagger:response diagramsResponse
type Diagrams []Diagram

// DiagramResponse provides response
// swagger:response diagramResponse
type DiagramResponse struct {
	Diagram
}
