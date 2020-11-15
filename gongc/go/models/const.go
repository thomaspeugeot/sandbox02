package models

import (
	"github.com/jinzhu/gorm"
)

// ConstModel is a representation of a const of a string Const in an enum
// swagger:model ConstModel
type ConstModel struct {

	// The Name of the Const
	Name string

	// The Value
	Value string

	// the Enum it belongs to
	EnumID uint

	// the Enum Name it belongs to
	EnumName string
}

// Const describres a const
// swagger:model ConstDB
type Const struct {
	gorm.Model

	ConstModel
}

// Consts arrays consts
// swagger:response constsResponse
type Consts []Const

// ConstResponse provides response
// swagger:response constResponse
type ConstResponse struct {
	Const
}
