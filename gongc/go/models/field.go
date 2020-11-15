package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

// FieldModel is a representation of a field of a metabaron Field
// a metabaron struct is a struct with field such as :
// - a basic type (string, int, ...)
// - a pointer to a metabaron struct (one ZeroOrOneRelationship)
// - a collection of pointers to a metabaron struct
// swagger:model FieldModel
type FieldModel struct {

	// The Name of the Struct Field
	Name string

	// the Kind of the Struct Field
	// for a metabaron struct
	Kind reflect.Kind

	// the Struct it belongs to
	StructID uint

	// the Struct Name it belongs to
	StructName string

	// if Kind is a Ptr or an Array, the Struct behind, nil otherwise
	AssociatedStructID uint

	// The Name of the Associated Struct
	AssociatedStructName string
}

// Field describres a field
// swagger:model FieldDB
type Field struct {
	gorm.Model

	FieldModel
}

// Fields arrays fields
// swagger:response fieldsResponse
type Fields []Field

// FieldResponse provides response
// swagger:response fieldResponse
type FieldResponse struct {
	Field
}
