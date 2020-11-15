// generated by genORMModelDB.go
package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
)

// UmlscAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model umlscAPI
type UmlscAPI struct {
	models.Umlsc

	// association fields

	// ID generated for the implementation of the field Umlsc{}.Umlscs []*Pkgelt
	Pkgelt_UmlscsDBID uint
}

// UmlscDB describes a umlsc in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model umlscDB
type UmlscDB struct {
	gorm.Model

	UmlscAPI
}

// UmlscDBs arrays umlscDBs
// swagger:response umlscDBsResponse
type UmlscDBs []UmlscDB

// UmlscDBResponse provides response
// swagger:response umlscDBResponse
type UmlscDBResponse struct {
	UmlscDB
}

// ModelToORMUmlscTranslate is a translation function from models object to ORM objects
func ModelToORMUmlscTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that umlscStore is nil as well as umlscDBs
		if map_UmlscDBID_UmlscPtr != nil {
			err := errors.New("In CreateMode translation, map_UmlscDBID_UmlscPtr should be nil")
			return err
		}

		if map_UmlscDBID_UmlscDB != nil {
			err := errors.New("In CreateMode translation, map_UmlscDBID_UmlscDB should be nil")
			return err
		}

		if map_UmlscPtr_UmlscDBID != nil {
			err := errors.New("In CreateMode translation, map_UmlscPtr_UmlscDBID should be nil")
			return err
		}

		tmp := make(map[uint]*models.Umlsc, 0)
		map_UmlscDBID_UmlscPtr = &tmp

		tmpDB := make(map[uint]*UmlscDB, 0)
		map_UmlscDBID_UmlscDB = &tmpDB

		tmpID := make(map[*models.Umlsc]uint, 0)
		map_UmlscPtr_UmlscDBID = &tmpID

		for _, umlsc := range models.AllModelStore.Umlscs {

			// initiate umlsc
			var umlscDB UmlscDB
			umlscDB.Umlsc = *umlsc

			query := db.Create(&umlscDB)
			if query.Error != nil {
				return query.Error
			}

			// update stores
			(*map_UmlscPtr_UmlscDBID)[umlsc] = umlscDB.ID
			(*map_UmlscDBID_UmlscPtr)[umlscDB.ID] = umlsc
			(*map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDB
		}
	} else { // UpdateMode, update IDs of Pointer Fields of ORM object

		// check that umlscStore is not nil
		if map_UmlscDBID_UmlscPtr == nil {
			err := errors.New("In UpdateMode translation, umlscStore should not be nil")
			return err
		}

		if map_UmlscDBID_UmlscDB == nil {
			err := errors.New("In UpdateMode translation, umlscStore should not be nil")
			return err
		}

		// update fields of umlscDB with fields of umlsc
		for _, umlsc := range models.AllModelStore.Umlscs {
			umlscDBID := (*map_UmlscPtr_UmlscDBID)[umlsc]
			umlscDB := (*map_UmlscDBID_UmlscDB)[umlscDBID]

			umlscDB.Umlsc = *umlsc
		}

		// parse model objects ot update associations
		for idx, umlsc := range *map_UmlscDBID_UmlscPtr {

			// fetch matching umlscDB
			if umlscDB, ok := (*map_UmlscDBID_UmlscDB)[idx]; ok {
				// set {{Fieldname}}ID

				// set StatesIDs reverse pointer to State
				for _, State := range umlsc.States {
					if StateDBID, ok := (*map_StatePtr_StateDBID)[State]; ok {
						if StateDB, ok := (*map_StateDBID_StateDB)[StateDBID]; ok {
							StateDB.Umlsc_StatesDBID = umlscDB.ID
							if q := db.Save(&StateDB); q.Error != nil {
								return q.Error
							}
						}
					}
				}

				query := db.Save(&umlscDB)
				if query.Error != nil {
					return query.Error
				}

			} else {
				err := errors.New(
					fmt.Sprintf("In UpdateMode translation, umlscStore should not be nil %v %v",
						umlscDB, umlsc))
				return err
			}
		}
	}
	return nil
}

// stores UmlscDB according to their gorm ID
var map_UmlscDBID_UmlscDB *map[uint]*UmlscDB

// stores UmlscDB ID according to Umlsc address
var map_UmlscPtr_UmlscDBID *map[*models.Umlsc]uint

// stores Umlsc according to their gorm ID
var map_UmlscDBID_UmlscPtr *map[uint]*models.Umlsc

// ORMToModelUmlscTranslate is a translation function from ORM object to models objects
// This function used the uint ID of the ORM object to create or update (according to translationImpact)
// maps of respectively ORM and models objects
//
// In create mode,
func ORMToModelUmlscTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that umlscStores are nil

		if map_UmlscDBID_UmlscPtr != nil {
			err := errors.New("In CreateMode translation, Parameters umlscStore should be nil")
			return err
		}

		if map_UmlscDBID_UmlscDB != nil {
			err := errors.New("In CreateMode translation, parameters UmlscDBStore should be nil")
			return err
		}

		// init stores
		tmp := make(map[uint]*models.Umlsc, 0)
		map_UmlscDBID_UmlscPtr = &tmp

		tmpDB := make(map[uint]*UmlscDB, 0)
		map_UmlscDBID_UmlscDB = &tmpDB

		tmpID := make(map[*models.Umlsc]uint, 0)
		map_UmlscPtr_UmlscDBID = &tmpID

		models.AllModelStore.Umlscs = make([]*models.Umlsc, 0)

		umlscDBArray := make([]UmlscDB, 0)
		query := db.Find(&umlscDBArray)
		if query.Error != nil {
			return query.Error
		}

		// copy orm objects to the two stores
		for _, umlscDB := range umlscDBArray {

			// create entries in the tree maps.
			umlsc := umlscDB.Umlsc
			(*map_UmlscDBID_UmlscPtr)[umlscDB.ID] = &umlsc

			(*map_UmlscPtr_UmlscDBID)[&umlsc] = umlscDB.ID

			umlscDBCopy := umlscDB
			(*map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDBCopy

			// append model store with the new element
			models.AllModelStore.Umlscs = append(models.AllModelStore.Umlscs, &umlsc)
		}
	} else { // UpdateMode
		// for later, update of the data field

		// check that umlscStore is not nil
		if map_UmlscDBID_UmlscPtr == nil {
			err := errors.New("In UpdateMode translation, umlscStore should not be nil")
			return err
		}

		if map_UmlscDBID_UmlscDB == nil {
			err := errors.New("In UpdateMode translation, umlscStore should not be nil")
			return err
		}

		// update fields of umlscDB with fields of umlsc
		for _, umlsc := range models.AllModelStore.Umlscs {
			umlscDBID := (*map_UmlscPtr_UmlscDBID)[umlsc]
			umlscDB := (*map_UmlscDBID_UmlscDB)[umlscDBID]

			*umlsc = umlscDB.Umlsc
		}

		// parse all DB instance and update all pointer fields of the translated models instance
		for _, umlscDB := range *map_UmlscDBID_UmlscDB {
			umlsc := (*map_UmlscDBID_UmlscPtr)[umlscDB.ID]
			if umlsc == nil {
				err := errors.New("cannot find translated instance in models store")
				return err
			}

			// parse all StateDB and redeem the array of poiners to Umlsc
			for _, StateDB := range *map_StateDBID_StateDB {
				if StateDB.Umlsc_StatesDBID == umlscDB.ID {
					State := (*map_StateDBID_StatePtr)[StateDB.ID]
					umlsc.States = append(umlsc.States, State)
				}
			}

		}
	}

	return nil
}

func (allORMStoreStruct *AllORMStoreStruct) CreateORMUmlsc(umlsc *models.Umlsc) {

	CreateORMUmlsc(allORMStoreStruct.db, umlsc)
}

// CreateORMUmlsc creates ORM{{Strucname}} in DB from umlsc
func CreateORMUmlsc(
	db *gorm.DB,
	umlsc *models.Umlsc) (Error error) {

	// initiate umlsc
	var umlscDB UmlscDB
	umlscDB.Umlsc = *umlsc

	query := db.Create(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*map_UmlscPtr_UmlscDBID)[umlsc] = umlscDB.ID
	(*map_UmlscDBID_UmlscPtr)[umlscDB.ID] = umlsc
	(*map_UmlscDBID_UmlscDB)[umlscDB.ID] = &umlscDB

	return
}

func (allORMStoreStruct *AllORMStoreStruct) DeleteORMUmlsc(umlsc *models.Umlsc) {

	DeleteORMUmlsc(allORMStoreStruct.db, umlsc)
}

func DeleteORMUmlsc(
	db *gorm.DB,
	umlsc *models.Umlsc) (Error error) {

	umlscDBID := (*map_UmlscPtr_UmlscDBID)[umlsc]
	umlscDB := (*map_UmlscDBID_UmlscDB)[umlscDBID]

	query := db.Unscoped().Delete(&umlscDB)
	if query.Error != nil {
		return query.Error
	}

	delete(*map_UmlscPtr_UmlscDBID, umlsc)
	delete(*map_UmlscDBID_UmlscPtr, umlscDB.ID)
	delete(*map_UmlscDBID_UmlscDB, umlscDBID)

	return
}
