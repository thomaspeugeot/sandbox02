package models

// swagger:enum Typeaction
type Typeaction string

// values for Action Type
const (
	UNMARSHALL_ALL_DIAGRAMS Typeaction = "UNMARSHALL_ALL_DIAGRAMS" // iota // Parse the spinosa model (temp)
	MARSHALL_ALL_DIAGRAMS   Typeaction = "MARSHALL_ALL_DIAGRAMS"
	PRINT_ALL_DOCUMENTS     Typeaction = "PRINT_ALL_DOCUMENTS"
)

// Gorgoaction is a representation of a action of a metabaron Gorgoaction
// a metabaron action
// swagger:model Gorgoaction
type Gorgoaction struct {

	// The Name of the Action
	Name string

	// The type of the action
	Type Typeaction
}
