package models

import (
	"github.com/jinzhu/gorm"
)

// EnumModel is a representation of go Enum
// swagger:model enumModel
type EnumModel struct {

	// The Name of the Type
	Name string
}

// Enum describres a enum
// swagger:model enumDB
type EnumDB struct {
	gorm.Model

	EnumModel
}

// Enums arrays enums
// swagger:response enumsResponse
type Enums []EnumModel

// EnumResponse provides response
// swagger:response enumResponse
type EnumResponse struct {
	EnumModel
}

// EnumMap provide a map to enum
type EnumMap map[string]*EnumModel
