package walk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/gorm"
)

const errorTemplate = `
package controllers

import (
	"fmt"
)

// genQuery return the name of the column
func genQuery( columnName string) string {
	return fmt.Sprintf("%s = ?", columnName)
}

// A GenericError is the default error message that is generated.
// For certain status codes there are more appropriate error structures.
//
// swagger:response genericError
type GenericError struct {
	// in: body
	Body struct {
		Code    int32 ` + "`" + `json:"code"` + "`" + `
		Message string ` + "`" + `json:"message"` + "`" + `
	} ` + "`" + `json:"body"` + "`" + `
}

// A ValidationError is an that is generated for validation failures.
// It has the same fields as a generic error but adds a Field property.
//
// swagger:response validationError
type ValidationError struct {
	// in: body
	Body struct {
		Code    int32  ` + "`" + `json:"code"` + "`" + `
		Message string ` + "`" + `json:"message"` + "`" + `
		Field   string ` + "`" + `json:"field"` + "`" + `
	} ` + "`" + `json:"body"` + "`" + `
}
`

// GenControllersErrorCodes generates the setup file for the gorm
func GenControllersErrorCodes(db *gorm.DB) {

	// generates the generic errors file
	filename := filepath.Join(ControllersPkgGenPath, "errors.go")

	// we should use go generate
	log.Println("generating controller error file : " + filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(f, "%s", errorTemplate)

	defer f.Close()
}
