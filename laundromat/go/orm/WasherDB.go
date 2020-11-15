// generated by genORMModelDB.go
package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/laundromat/go/models"
)

// WasherAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model washerAPI
type WasherAPI struct {
	models.Washer

	// association fields

	// field Machine is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable a GORM `HAS ONE` association
	MachineID *uint

	// the associated Struct has a Name field, therefore it is generated to compute views with this relation
	MachineName string
}

// WasherDB describes a washer in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model washerDB
type WasherDB struct {
	gorm.Model

	WasherAPI
}

// WasherDBs arrays washerDBs
// swagger:response washerDBsResponse
type WasherDBs []WasherDB

// WasherDBResponse provides response
// swagger:response washerDBResponse
type WasherDBResponse struct {
	WasherDB
}

// ModelToORMWasherTranslate is a translation function from models object to ORM objects
func ModelToORMWasherTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that washerStore is nil as well as washerDBs
		if map_WasherDBID_WasherPtr != nil {
			err := errors.New("In CreateMode translation, map_WasherDBID_WasherPtr should be nil")
			return err
		}

		if map_WasherDBID_WasherDB != nil {
			err := errors.New("In CreateMode translation, map_WasherDBID_WasherDB should be nil")
			return err
		}

		if map_WasherPtr_WasherDBID != nil {
			err := errors.New("In CreateMode translation, map_WasherPtr_WasherDBID should be nil")
			return err
		}

		tmp := make(map[uint]*models.Washer, 0)
		map_WasherDBID_WasherPtr = &tmp

		tmpDB := make(map[uint]*WasherDB, 0)
		map_WasherDBID_WasherDB = &tmpDB

		tmpID := make(map[*models.Washer]uint, 0)
		map_WasherPtr_WasherDBID = &tmpID

		for _, washer := range models.AllModelStore.Washers {

			// initiate washer
			var washerDB WasherDB
			washerDB.Washer = *washer

			query := db.Create(&washerDB)
			if query.Error != nil {
				return query.Error
			}

			// update stores
			(*map_WasherPtr_WasherDBID)[washer] = washerDB.ID
			(*map_WasherDBID_WasherPtr)[washerDB.ID] = washer
			(*map_WasherDBID_WasherDB)[washerDB.ID] = &washerDB
		}
	} else { // UpdateMode, update IDs of Pointer Fields of ORM object

		// check that washerStore is not nil
		if map_WasherDBID_WasherPtr == nil {
			err := errors.New("In UpdateMode translation, washerStore should not be nil")
			return err
		}

		if map_WasherDBID_WasherDB == nil {
			err := errors.New("In UpdateMode translation, washerStore should not be nil")
			return err
		}

		// update fields of washerDB with fields of washer
		for _, washer := range models.AllModelStore.Washers {
			washerDBID := (*map_WasherPtr_WasherDBID)[washer]
			washerDB := (*map_WasherDBID_WasherDB)[washerDBID]

			washerDB.Washer = *washer
		}

		// parse model objects ot update associations
		for idx, washer := range *map_WasherDBID_WasherPtr {

			// fetch matching washerDB
			if washerDB, ok := (*map_WasherDBID_WasherDB)[idx]; ok {
				// set {{Fieldname}}ID

				// set MachineID
				if washer.Machine != nil {
					if machineId, ok := (*map_MachinePtr_MachineDBID)[washer.Machine]; ok {
						washerDB.MachineID = &machineId
					}
				}

				query := db.Save(&washerDB)
				if query.Error != nil {
					return query.Error
				}

			} else {
				err := errors.New(
					fmt.Sprintf("In UpdateMode translation, washerStore should not be nil %v %v",
						washerDB, washer))
				return err
			}
		}
	}
	return nil
}

// stores WasherDB according to their gorm ID
var map_WasherDBID_WasherDB *map[uint]*WasherDB

// stores WasherDB ID according to Washer address
var map_WasherPtr_WasherDBID *map[*models.Washer]uint

// stores Washer according to their gorm ID
var map_WasherDBID_WasherPtr *map[uint]*models.Washer

// ORMToModelWasherTranslate is a translation function from ORM object to models objects
// This function used the uint ID of the ORM object to create or update (according to translationImpact)
// maps of respectively ORM and models objects
//
// In create mode,
func ORMToModelWasherTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that washerStores are nil

		if map_WasherDBID_WasherPtr != nil {
			err := errors.New("In CreateMode translation, Parameters washerStore should be nil")
			return err
		}

		if map_WasherDBID_WasherDB != nil {
			err := errors.New("In CreateMode translation, parameters WasherDBStore should be nil")
			return err
		}

		// init stores
		tmp := make(map[uint]*models.Washer, 0)
		map_WasherDBID_WasherPtr = &tmp

		tmpDB := make(map[uint]*WasherDB, 0)
		map_WasherDBID_WasherDB = &tmpDB

		tmpID := make(map[*models.Washer]uint, 0)
		map_WasherPtr_WasherDBID = &tmpID

		models.AllModelStore.Washers = make([]*models.Washer, 0)

		washerDBArray := make([]WasherDB, 0)
		query := db.Find(&washerDBArray)
		if query.Error != nil {
			return query.Error
		}

		// copy orm objects to the two stores
		for _, washerDB := range washerDBArray {

			// create entries in the tree maps.
			washer := washerDB.Washer
			(*map_WasherDBID_WasherPtr)[washerDB.ID] = &washer

			(*map_WasherPtr_WasherDBID)[&washer] = washerDB.ID

			washerDBCopy := washerDB
			(*map_WasherDBID_WasherDB)[washerDB.ID] = &washerDBCopy

			// append model store with the new element
			models.AllModelStore.Washers = append(models.AllModelStore.Washers, &washer)
		}
	} else { // UpdateMode
		// for later, update of the data field

		// check that washerStore is not nil
		if map_WasherDBID_WasherPtr == nil {
			err := errors.New("In UpdateMode translation, washerStore should not be nil")
			return err
		}

		if map_WasherDBID_WasherDB == nil {
			err := errors.New("In UpdateMode translation, washerStore should not be nil")
			return err
		}

		// update fields of washerDB with fields of washer
		for _, washer := range models.AllModelStore.Washers {
			washerDBID := (*map_WasherPtr_WasherDBID)[washer]
			washerDB := (*map_WasherDBID_WasherDB)[washerDBID]

			*washer = washerDB.Washer
		}

		// parse all DB instance and update all pointer fields of the translated models instance
		for _, washerDB := range *map_WasherDBID_WasherDB {
			washer := (*map_WasherDBID_WasherPtr)[washerDB.ID]
			if washer == nil {
				err := errors.New("cannot find translated instance in models store")
				return err
			}

			// Machine field
			if washerDB.MachineID != nil {
				washer.Machine = (*map_MachineDBID_MachinePtr)[*(washerDB.MachineID)]
			}

		}
	}

	return nil
}

func (allORMStoreStruct *AllORMStoreStruct) CreateORMWasher(washer *models.Washer) {

	CreateORMWasher(allORMStoreStruct.db, washer)
}

// CreateORMWasher creates ORM{{Strucname}} in DB from washer
func CreateORMWasher(
	db *gorm.DB,
	washer *models.Washer) (Error error) {

	// initiate washer
	var washerDB WasherDB
	washerDB.Washer = *washer

	query := db.Create(&washerDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*map_WasherPtr_WasherDBID)[washer] = washerDB.ID
	(*map_WasherDBID_WasherPtr)[washerDB.ID] = washer
	(*map_WasherDBID_WasherDB)[washerDB.ID] = &washerDB

	return
}

func (allORMStoreStruct *AllORMStoreStruct) DeleteORMWasher(washer *models.Washer) {

	DeleteORMWasher(allORMStoreStruct.db, washer)
}

func DeleteORMWasher(
	db *gorm.DB,
	washer *models.Washer) (Error error) {

	washerDBID := (*map_WasherPtr_WasherDBID)[washer]
	washerDB := (*map_WasherDBID_WasherDB)[washerDBID]

	query := db.Unscoped().Delete(&washerDB)
	if query.Error != nil {
		return query.Error
	}

	delete(*map_WasherPtr_WasherDBID, washer)
	delete(*map_WasherDBID_WasherPtr, washerDB.ID)
	delete(*map_WasherDBID_WasherDB, washerDBID)

	return
}
