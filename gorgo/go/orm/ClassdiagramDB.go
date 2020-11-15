// generated by genORMModelDB.go
package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
)

// ClassdiagramAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model classdiagramAPI
type ClassdiagramAPI struct {
	models.Classdiagram

	// association fields

	// ID generated for the implementation of the field Classdiagram{}.Classdiagrams []*Pkgelt
	Pkgelt_ClassdiagramsDBID uint
}

// ClassdiagramDB describes a classdiagram in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model classdiagramDB
type ClassdiagramDB struct {
	gorm.Model

	ClassdiagramAPI
}

// ClassdiagramDBs arrays classdiagramDBs
// swagger:response classdiagramDBsResponse
type ClassdiagramDBs []ClassdiagramDB

// ClassdiagramDBResponse provides response
// swagger:response classdiagramDBResponse
type ClassdiagramDBResponse struct {
	ClassdiagramDB
}

// ModelToORMClassdiagramTranslate is a translation function from models object to ORM objects
func ModelToORMClassdiagramTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that classdiagramStore is nil as well as classdiagramDBs
		if map_ClassdiagramDBID_ClassdiagramPtr != nil {
			err := errors.New("In CreateMode translation, map_ClassdiagramDBID_ClassdiagramPtr should be nil")
			return err
		}

		if map_ClassdiagramDBID_ClassdiagramDB != nil {
			err := errors.New("In CreateMode translation, map_ClassdiagramDBID_ClassdiagramDB should be nil")
			return err
		}

		if map_ClassdiagramPtr_ClassdiagramDBID != nil {
			err := errors.New("In CreateMode translation, map_ClassdiagramPtr_ClassdiagramDBID should be nil")
			return err
		}

		tmp := make(map[uint]*models.Classdiagram, 0)
		map_ClassdiagramDBID_ClassdiagramPtr = &tmp

		tmpDB := make(map[uint]*ClassdiagramDB, 0)
		map_ClassdiagramDBID_ClassdiagramDB = &tmpDB

		tmpID := make(map[*models.Classdiagram]uint, 0)
		map_ClassdiagramPtr_ClassdiagramDBID = &tmpID

		for _, classdiagram := range models.AllModelStore.Classdiagrams {

			// initiate classdiagram
			var classdiagramDB ClassdiagramDB
			classdiagramDB.Classdiagram = *classdiagram

			query := db.Create(&classdiagramDB)
			if query.Error != nil {
				return query.Error
			}

			// update stores
			(*map_ClassdiagramPtr_ClassdiagramDBID)[classdiagram] = classdiagramDB.ID
			(*map_ClassdiagramDBID_ClassdiagramPtr)[classdiagramDB.ID] = classdiagram
			(*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDB.ID] = &classdiagramDB
		}
	} else { // UpdateMode, update IDs of Pointer Fields of ORM object

		// check that classdiagramStore is not nil
		if map_ClassdiagramDBID_ClassdiagramPtr == nil {
			err := errors.New("In UpdateMode translation, classdiagramStore should not be nil")
			return err
		}

		if map_ClassdiagramDBID_ClassdiagramDB == nil {
			err := errors.New("In UpdateMode translation, classdiagramStore should not be nil")
			return err
		}

		// update fields of classdiagramDB with fields of classdiagram
		for _, classdiagram := range models.AllModelStore.Classdiagrams {
			classdiagramDBID := (*map_ClassdiagramPtr_ClassdiagramDBID)[classdiagram]
			classdiagramDB := (*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDBID]

			classdiagramDB.Classdiagram = *classdiagram
		}

		// parse model objects ot update associations
		for idx, classdiagram := range *map_ClassdiagramDBID_ClassdiagramPtr {

			// fetch matching classdiagramDB
			if classdiagramDB, ok := (*map_ClassdiagramDBID_ClassdiagramDB)[idx]; ok {
				// set {{Fieldname}}ID

				// set ClassshapesIDs reverse pointer to Classshape
				for _, Classshape := range classdiagram.Classshapes {
					if ClassshapeDBID, ok := (*map_ClassshapePtr_ClassshapeDBID)[Classshape]; ok {
						if ClassshapeDB, ok := (*map_ClassshapeDBID_ClassshapeDB)[ClassshapeDBID]; ok {
							ClassshapeDB.Classdiagram_ClassshapesDBID = classdiagramDB.ID
							if q := db.Save(&ClassshapeDB); q.Error != nil {
								return q.Error
							}
						}
					}
				}

				query := db.Save(&classdiagramDB)
				if query.Error != nil {
					return query.Error
				}

			} else {
				err := errors.New(
					fmt.Sprintf("In UpdateMode translation, classdiagramStore should not be nil %v %v",
						classdiagramDB, classdiagram))
				return err
			}
		}
	}
	return nil
}

// stores ClassdiagramDB according to their gorm ID
var map_ClassdiagramDBID_ClassdiagramDB *map[uint]*ClassdiagramDB

// stores ClassdiagramDB ID according to Classdiagram address
var map_ClassdiagramPtr_ClassdiagramDBID *map[*models.Classdiagram]uint

// stores Classdiagram according to their gorm ID
var map_ClassdiagramDBID_ClassdiagramPtr *map[uint]*models.Classdiagram

// ORMToModelClassdiagramTranslate is a translation function from ORM object to models objects
// This function used the uint ID of the ORM object to create or update (according to translationImpact)
// maps of respectively ORM and models objects
//
// In create mode,
func ORMToModelClassdiagramTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that classdiagramStores are nil

		if map_ClassdiagramDBID_ClassdiagramPtr != nil {
			err := errors.New("In CreateMode translation, Parameters classdiagramStore should be nil")
			return err
		}

		if map_ClassdiagramDBID_ClassdiagramDB != nil {
			err := errors.New("In CreateMode translation, parameters ClassdiagramDBStore should be nil")
			return err
		}

		// init stores
		tmp := make(map[uint]*models.Classdiagram, 0)
		map_ClassdiagramDBID_ClassdiagramPtr = &tmp

		tmpDB := make(map[uint]*ClassdiagramDB, 0)
		map_ClassdiagramDBID_ClassdiagramDB = &tmpDB

		tmpID := make(map[*models.Classdiagram]uint, 0)
		map_ClassdiagramPtr_ClassdiagramDBID = &tmpID

		models.AllModelStore.Classdiagrams = make([]*models.Classdiagram, 0)

		classdiagramDBArray := make([]ClassdiagramDB, 0)
		query := db.Find(&classdiagramDBArray)
		if query.Error != nil {
			return query.Error
		}

		// copy orm objects to the two stores
		for _, classdiagramDB := range classdiagramDBArray {

			// create entries in the tree maps.
			classdiagram := classdiagramDB.Classdiagram
			(*map_ClassdiagramDBID_ClassdiagramPtr)[classdiagramDB.ID] = &classdiagram

			(*map_ClassdiagramPtr_ClassdiagramDBID)[&classdiagram] = classdiagramDB.ID

			classdiagramDBCopy := classdiagramDB
			(*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDB.ID] = &classdiagramDBCopy

			// append model store with the new element
			models.AllModelStore.Classdiagrams = append(models.AllModelStore.Classdiagrams, &classdiagram)
		}
	} else { // UpdateMode
		// for later, update of the data field

		// check that classdiagramStore is not nil
		if map_ClassdiagramDBID_ClassdiagramPtr == nil {
			err := errors.New("In UpdateMode translation, classdiagramStore should not be nil")
			return err
		}

		if map_ClassdiagramDBID_ClassdiagramDB == nil {
			err := errors.New("In UpdateMode translation, classdiagramStore should not be nil")
			return err
		}

		// update fields of classdiagramDB with fields of classdiagram
		for _, classdiagram := range models.AllModelStore.Classdiagrams {
			classdiagramDBID := (*map_ClassdiagramPtr_ClassdiagramDBID)[classdiagram]
			classdiagramDB := (*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDBID]

			*classdiagram = classdiagramDB.Classdiagram
		}

		// parse all DB instance and update all pointer fields of the translated models instance
		for _, classdiagramDB := range *map_ClassdiagramDBID_ClassdiagramDB {
			classdiagram := (*map_ClassdiagramDBID_ClassdiagramPtr)[classdiagramDB.ID]
			if classdiagram == nil {
				err := errors.New("cannot find translated instance in models store")
				return err
			}

			// parse all ClassshapeDB and redeem the array of poiners to Classdiagram
			for _, ClassshapeDB := range *map_ClassshapeDBID_ClassshapeDB {
				if ClassshapeDB.Classdiagram_ClassshapesDBID == classdiagramDB.ID {
					Classshape := (*map_ClassshapeDBID_ClassshapePtr)[ClassshapeDB.ID]
					classdiagram.Classshapes = append(classdiagram.Classshapes, Classshape)
				}
			}

		}
	}

	return nil
}

func (allORMStoreStruct *AllORMStoreStruct) CreateORMClassdiagram(classdiagram *models.Classdiagram) {

	CreateORMClassdiagram(allORMStoreStruct.db, classdiagram)
}

// CreateORMClassdiagram creates ORM{{Strucname}} in DB from classdiagram
func CreateORMClassdiagram(
	db *gorm.DB,
	classdiagram *models.Classdiagram) (Error error) {

	// initiate classdiagram
	var classdiagramDB ClassdiagramDB
	classdiagramDB.Classdiagram = *classdiagram

	query := db.Create(&classdiagramDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*map_ClassdiagramPtr_ClassdiagramDBID)[classdiagram] = classdiagramDB.ID
	(*map_ClassdiagramDBID_ClassdiagramPtr)[classdiagramDB.ID] = classdiagram
	(*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDB.ID] = &classdiagramDB

	return
}

func (allORMStoreStruct *AllORMStoreStruct) DeleteORMClassdiagram(classdiagram *models.Classdiagram) {

	DeleteORMClassdiagram(allORMStoreStruct.db, classdiagram)
}

func DeleteORMClassdiagram(
	db *gorm.DB,
	classdiagram *models.Classdiagram) (Error error) {

	classdiagramDBID := (*map_ClassdiagramPtr_ClassdiagramDBID)[classdiagram]
	classdiagramDB := (*map_ClassdiagramDBID_ClassdiagramDB)[classdiagramDBID]

	query := db.Unscoped().Delete(&classdiagramDB)
	if query.Error != nil {
		return query.Error
	}

	delete(*map_ClassdiagramPtr_ClassdiagramDBID, classdiagram)
	delete(*map_ClassdiagramDBID_ClassdiagramPtr, classdiagramDB.ID)
	delete(*map_ClassdiagramDBID_ClassdiagramDB, classdiagramDBID)

	return
}
