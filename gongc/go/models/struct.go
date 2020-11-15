package models

import (
	"github.com/jinzhu/gorm"
)

// StructModel is a representation of go Struct
// swagger:model structModel
type StructModel struct {

	// The Name of the Type
	Name string

	// The Attributes list (not working at the moment, ...)
	Field []*Field
}

// Struct describres a struct
// swagger:model structDB
type Struct struct {
	gorm.Model

	StructModel
}

// Structs arrays structs
// swagger:response structsResponse
type Structs []Struct

// StructResponse provides response
// swagger:response structResponse
type StructResponse struct {
	Struct
}

// StructMap provide a map to struct
type StructMap map[string]*Struct
