package walk

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// TODO #13 Generate association controllers (for pointers & for slices of pointers)
const ptrRevAssocCtrlTplt = `
// Get{{Structname}}{{AssocStructname}}sVia{{Fieldname}} swagger:route GET /{{structname}}s/{ID}/{{assocStructname}}svia{{fieldname}} {{structname}}s get{{Structname}}{{AssocStructname}}sVia{{Fieldname}}
//
// Gets {{assocStructname}}s of {{structname}} via field {{Fieldname}} of {{AssocStructname}}.
//
// Responses:
//    default: genericError
//        200: {{assocStructname}}DBsResponse
func Get{{Structname}}{{AssocStructname}}sVia{{Fieldname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get {{structname}}
	var {{structname}} orm.{{Structname}}DB
	if err := db.First(&{{structname}}, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get {{assocStructname}}s
	var {{assocStructname}}DBs orm.{{AssocStructname}}DBs
	columnName := gorm.ToColumnName("{{Fieldname}}ID")
	query := db.Where( genQuery(columnName), {{structname}}.ID).Find(&{{assocStructname}}DBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, {{assocStructname}}DBs)
}
`

const ptrAssocCtrlTplt = `// Get{{Structname}}{{Fieldname}} swagger:route GET /{{structname}}s/{ID}/{{fieldname}} {{structname}}s get{{Structname}}{{Fieldname}}
//
// Gets {{fieldname}} of a {{structname}}.
//
// Responses:
//    default: genericError
//        200: {{assocStructname}}DBResponse
func Get{{Structname}}{{Fieldname}}(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get {{structname}} in DB
	var {{structname}} orm.{{Structname}}DB
	if err := db.First(&{{structname}}, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Get {{assocStructname}} in DB
	var {{assocStructname}}{{Fieldname}} orm.{{AssocStructname}}DB
	if err := db.First(&{{assocStructname}}{{Fieldname}}, *{{structname}}.{{Fieldname}}ID).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, {{assocStructname}}{{Fieldname}})
}

`

const ptrRevAssocCtrlTpltReg = `
	r.GET("/{{structname}}s/:id/{{assocStructname}}svia{{fieldname}}", Get{{Structname}}{{AssocStructname}}sVia{{Fieldname}})`
const ptrAssocCtrlTpltReg = `
	r.GET("/{{structname}}s/:id/{{fieldname}}", Get{{Structname}}{{Fieldname}})`

// genControllersAssociation takes a struct and
// returns 3 generated codes :
// - the controllers for the association related to this struct (for instance r.GET( editors/ID/books, getEditorBooks() )
// - the ids of those controllers to declare them in the ID param
// - the registration code for the controlers of the struct
//
// generates 2 kind of controllers
// - the forward association, i.e.  r.GET( books/ID/editor, getBookEditor() )
// - the reverse association, i.e. r.GET( editors/ID/books, getEditorBooks() )
func genControllersAssociation(db *gorm.DB, _struct models.Struct) (string, string, string) {

	lowerCaseStructName := strings.ToLower(_struct.Name)

	var controllersCode string    // concatenation of all controllers
	var controllersCodeID string  // to append to the swagger parameter
	var controllersCodeReg string // to append to register call

	// FORWARD association
	// get fields of struct for pointers or array
	var fields models.Fields
	query := db.Model(&_struct).Related(&fields)
	if query.Error != nil {
		log.Fatal(query.Error.Error())
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name
	})

	for _, field := range fields {

		if field.Kind == reflect.Ptr && field.AssociatedStructID != 0 {

			// fetch the assoc struct
			var assocStruct models.Struct
			db.First(&assocStruct, field.AssociatedStructID)

			controllerCode := strings.ReplaceAll(ptrAssocCtrlTplt, "{{Structname}}", _struct.Name)
			controllerCode = strings.ReplaceAll(controllerCode, "{{structname}}", lowerCaseStructName)

			lowerCaseAssocStructName := strings.ToLower(assocStruct.Name)
			controllerCode = strings.ReplaceAll(controllerCode, "{{AssocStructname}}", assocStruct.Name)
			controllerCode = strings.ReplaceAll(controllerCode, "{{assocStructname}}", lowerCaseAssocStructName)

			lowerCaseFieldtName := strings.ToLower(field.Name)
			controllerCode = strings.ReplaceAll(controllerCode, "{{Fieldname}}", field.Name)
			controllerCode = strings.ReplaceAll(controllerCode, "{{fieldname}}", lowerCaseFieldtName)

			controllersCode += controllerCode

			controllersCodeID += fmt.Sprintf(" get%s%s", _struct.Name, field.Name)

			controllerCodeReg := strings.ReplaceAll(ptrAssocCtrlTpltReg, "{{Structname}}", _struct.Name)
			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{structname}}", lowerCaseStructName)

			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{AssocStructname}}", assocStruct.Name)
			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{assocStructname}}", lowerCaseAssocStructName)

			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{AssocStructname}}", assocStruct.Name)
			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{assocStructname}}", lowerCaseAssocStructName)

			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{Fieldname}}", field.Name)
			controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{fieldname}}", lowerCaseFieldtName)

			controllersCodeReg += controllerCodeReg
		}
	}

	// REVERSE relation

	// fetch all association fields worthy of a assocation path
	// ie. where the AssociatedStructID is the struct of interest
	columnName := gorm.ToColumnName("AssociatedStructID")
	// log.Output(0, fmt.Sprintf("Column name: %s", columnName))
	queryAssoc := db.Where(fmt.Sprintf("%s = ?", columnName), _struct.ID).Find(&fields)
	if queryAssoc.Error != nil {
		log.Fatal(queryAssoc.Error.Error())
	}

	for _, field := range fields {

		// fetch the assoc struct
		var assocStruct models.Struct
		db.First(&assocStruct, field.StructID)

		controllerCode := strings.ReplaceAll(ptrRevAssocCtrlTplt, "{{Structname}}", _struct.Name)
		controllerCode = strings.ReplaceAll(controllerCode, "{{structname}}", lowerCaseStructName)

		lowerCaseAssocStructName := strings.ToLower(assocStruct.Name)
		controllerCode = strings.ReplaceAll(controllerCode, "{{AssocStructname}}", assocStruct.Name)
		controllerCode = strings.ReplaceAll(controllerCode, "{{assocStructname}}", lowerCaseAssocStructName)

		lowerCaseFieldtName := strings.ToLower(field.Name)
		controllerCode = strings.ReplaceAll(controllerCode, "{{Fieldname}}", field.Name)
		controllerCode = strings.ReplaceAll(controllerCode, "{{fieldname}}", lowerCaseFieldtName)

		controllersCode += controllerCode

		controllersCodeID += fmt.Sprintf(" get%s%ssVia%s", _struct.Name, assocStruct.Name, field.Name)

		controllerCodeReg := strings.ReplaceAll(ptrRevAssocCtrlTpltReg, "{{Structname}}", _struct.Name)
		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{structname}}", lowerCaseStructName)

		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{AssocStructname}}", assocStruct.Name)
		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{assocStructname}}", lowerCaseAssocStructName)

		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{AssocStructname}}", assocStruct.Name)
		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{assocStructname}}", lowerCaseAssocStructName)

		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{Fieldname}}", field.Name)
		controllerCodeReg = strings.ReplaceAll(controllerCodeReg, "{{fieldname}}", lowerCaseFieldtName)

		controllersCodeReg += controllerCodeReg
	}

	return controllersCode, controllersCodeID, controllersCodeReg
}
