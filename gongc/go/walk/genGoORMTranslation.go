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

const templateTranslationFile = `// generated by genORMTranslation.go
package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// TranslationImpact distinguish CREATE or UPDATE mode when translating
type TranslationImpact int

const (
	CreateMode TranslationImpact = iota // target objects are not initialized yet
	UpdateMode                          // target objects already exists and are only updated
)

// AllORMStoreStruct supports callback functions
type AllORMStoreStruct struct {
	db *gorm.DB
}

var AllORMStore AllORMStoreStruct

func (allORMStore *AllORMStoreStruct) SetDB(db *gorm.DB) {
	allORMStore.db = db
}

// AllORMToModels init reset all stores of models and orm instances
func AllModelsToORM(db *gorm.DB) (Error error) {

{{insertionPointForModelToORMCreateTranslations}}
{{insertionPointForModelToORMUpdateTranslations}}

	return nil
}

// AllORMToModels update all stores of models to orm instances
func AllModelsUpdateToORM(db *gorm.DB) (Error error) {

{{insertionPointForModelToORMUpdateTranslations}}

	return nil
}


// AllORMToModels init reset all stores of models and orm instances
func AllORMToModels(db *gorm.DB) (Error error) {

	// Init all instances of in the model store
{{insertionPointORMCreateToModelTranslations}}
{{insertionPointORMUpdateToModelTranslations}}
	return nil
}

// AllORMUpdateToModels init reset all stores of models and orm instances
func AllORMUpdateToModels(db *gorm.DB) (Error error) {

	// Init all instances of in the model store
{{insertionPointORMUpdateToModelTranslations}}
	return nil
}
`

const insertionElementORMToModelForCreateGo = `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := ORMToModel{{Structname}}Translate(
		CreateMode, 
		db); err != nil {
		err := errors.New("AllORMToModels, CreateMode Translation of {{Structname}} failed")
		return err
	}

`
const insertionElementORMToModelForUpdateGo = `	if err := ORMToModel{{Structname}}Translate(
			UpdateMode, 
			db); err != nil {
		err := errors.New("AllORMToModels, UpdateMode Translation of {{Structname}} failed")
		return err
	}
`

const insertionElementModelToORMForCreateGo = `
	map_{{Structname}}DBID_{{Structname}}DB = nil
	map_{{Structname}}Ptr_{{Structname}}DBID = nil
	map_{{Structname}}DBID_{{Structname}}Ptr = nil
	if err := ModelToORM{{Structname}}Translate(
		CreateMode,
		db); err != nil {
		return err
	}
`
const insertionElementModelToORMForUpdateGo = `	// now the pointers
	if err := ModelToORM{{Structname}}Translate(
		UpdateMode,
		db); err != nil {
		return err
	}
`

// GenGoORMTranslation generates the setup file for the gorm
func GenGoORMTranslation(db *gorm.DB) {

	// relative to the models package, swith to ../controlers package
	filename := filepath.Join(OrmPkgGenPath, "translations.go")

	// we should use go generate
	log.Println("generating orm translation file : " + filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	var res string
	{
		insertions := ""
		for _, _struct := range structs {
			insertions += strings.ReplaceAll(insertionElementORMToModelForCreateGo, "{{Structname}}", _struct.Name)
		}
		res = strings.ReplaceAll(templateTranslationFile, "{{insertionPointORMCreateToModelTranslations}}", insertions)

		insertions = "\n"
		for _, _struct := range structs {
			insertions += strings.ReplaceAll(insertionElementORMToModelForUpdateGo, "{{Structname}}", _struct.Name)
		}
		res = strings.ReplaceAll(res, "{{insertionPointORMUpdateToModelTranslations}}", insertions)
	}
	{
		insertions := ""
		for _, _struct := range structs {
			insertions += strings.ReplaceAll(insertionElementModelToORMForCreateGo, "{{Structname}}", _struct.Name)
		}
		res = strings.ReplaceAll(res, "{{insertionPointForModelToORMCreateTranslations}}", insertions)

		insertions = "\n	// and now update of the pointer fields\n"
		for _, _struct := range structs {
			insertions += strings.ReplaceAll(insertionElementModelToORMForUpdateGo, "{{Structname}}", _struct.Name)
		}
		res = strings.ReplaceAll(res, "{{insertionPointForModelToORMUpdateTranslations}}", insertions)
	}
	res = strings.ReplaceAll(res, "{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""))
	fmt.Fprintf(f, "%s", res)

	defer f.Close()
}