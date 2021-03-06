package walk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const apiGateTemplate = `// generated by genGoApiFate.go
package models

// APIGateStruct is the struct of the APIGate in the "models" namespace
//
// APIGateStruct enables CRUD of twin instances in the "api" namespace.
//
// APIGateStruct is a gate. It is "closed" or "open".
//
// The gate is "closed" if APIGateCRUDCallbacks is null. It means all operations on an APIGate
// will have no side effect
//
// The gate is "open" if APIGateCRUDCallbacks is not null, the behavior will depends on the implementation
// of the interface
//
// 	swagger:ignore
type APIGateStruct struct {
	APIGateCRUDCallbacks APIGateCRUDCallbacksInterface // pointer to the callback interface
}

// APIGate is a the singloton gate for CRUD operation between "models" and "api" namespace
//
// APIGate is a singloton instance to performs operations between instances in the "models" namespace and instances
// in the "api" namespace
//
// 	swagger:ignore
var APIGate APIGateStruct

// APIGateCRUDCallbacksInterface is the interface that has to be supported
// by the singloton that manages instances in the "api" namespace
//
// Because pointers that implement associations in the "models" space
// cannot be communicated to the front layer and the DB layer,
// they are irrevelant in the "api" namespace and are impletemented with an
// 	ID uint
// this ID can be computed if and only if the associated instance with the pointers has his twin already
// created in the "api" namespace
// therefore, when an "api" twin is created with a call to the "Create" function, the association IDs are not computed.
// They are computed with a call to the "Update" function
//
//	swagger:ignore
type APIGateCRUDCallbacksInterface interface {
{{InsertionPoint_CRUD_Functions_Signatures}}
}

{{InsertionPoint_CRUD_Functions}}
`

const CRUDFunctionSingaturesTemplate = `
	CreateAPI{{Structname}}({{structname}} *{{Structname}}) (ID uint, err error)
	UpdateAPI{{Structname}}({{structname}} *{{Structname}}) (ID uint, err error)
`

const CRUDFunctionTemplate = `

// CreateAPI{{Structname}} creates from {{structname}} an instance in the "api" namespace
//
// CreateAPI{{Structname}} performs a deep copy of {{structname}} fields, if gate is "open"
//
// CreateAPI{{Structname}} return the ID of the "api" instance
//
// It updates aPIGateStruct maps of ID to instances and map of instances to ID
func (aPIGateStruct *APIGateStruct) CreateAPI{{Structname}}({{structname}} *{{Structname}}) (IDAPI uint, err error) {
	if aPIGateStruct.APIGateCRUDCallbacks != nil {
		return aPIGateStruct.APIGateCRUDCallbacks.CreateAPI{{Structname}}({{structname}})
	}
	return
}

// UpdateAPI{{Structname}} updates the twin of {{structname}} in the "api" namespace with values of {{structname}}
//
// UpdateAPI{{Structname}} performs a deep copy from from the former to the later
// and computes IDs for the associations
// if gate is "open" CreateAPI{{Structname}} return the ID of the "api" instance
func (aPIGateStruct *APIGateStruct) UpdateAPI{{Structname}}({{structname}} *{{Structname}}) (ID uint, err error) {
	if aPIGateStruct.APIGateCRUDCallbacks != nil {
		return aPIGateStruct.APIGateCRUDCallbacks.UpdateAPI{{Structname}}({{structname}})
	}
	return
}

`

// GenGoApiGate generates the setup file for the gorm
func GenGoApiGate(db *gorm.DB) {

	// relative to the models package, swith to ./controlers package
	filename := filepath.Join(RelativePkgPath, "apigate.go")

	// we should use go generate
	log.Println("generating all models struct file : " + filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	var insertions, res string
	res = apiGateTemplate

	insertions = ""
	for _, _struct := range structs {
		insertions += replace2(CRUDFunctionSingaturesTemplate,
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
	}
	res = strings.ReplaceAll(res, "{{InsertionPoint_CRUD_Functions_Signatures}}", insertions)

	insertions = ""
	for _, _struct := range structs {
		insertions += replace2(CRUDFunctionTemplate,
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))
	}
	res = strings.ReplaceAll(res, "{{InsertionPoint_CRUD_Functions}}", insertions)

	fmt.Fprintf(f, "%s", res)

	defer f.Close()
}
