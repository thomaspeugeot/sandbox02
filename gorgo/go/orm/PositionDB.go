// generated by genORMModelDB.go
package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
)

// PositionAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model positionAPI
type PositionAPI struct {
	models.Position

	// association fields

}

// PositionDB describes a position in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model positionDB
type PositionDB struct {
	gorm.Model

	PositionAPI
}

// PositionDBs arrays positionDBs
// swagger:response positionDBsResponse
type PositionDBs []PositionDB

// PositionDBResponse provides response
// swagger:response positionDBResponse
type PositionDBResponse struct {
	PositionDB
}

// ModelToORMPositionTranslate is a translation function from models object to ORM objects
func ModelToORMPositionTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that positionStore is nil as well as positionDBs
		if map_PositionDBID_PositionPtr != nil {
			err := errors.New("In CreateMode translation, map_PositionDBID_PositionPtr should be nil")
			return err
		}

		if map_PositionDBID_PositionDB != nil {
			err := errors.New("In CreateMode translation, map_PositionDBID_PositionDB should be nil")
			return err
		}

		if map_PositionPtr_PositionDBID != nil {
			err := errors.New("In CreateMode translation, map_PositionPtr_PositionDBID should be nil")
			return err
		}

		tmp := make(map[uint]*models.Position, 0)
		map_PositionDBID_PositionPtr = &tmp

		tmpDB := make(map[uint]*PositionDB, 0)
		map_PositionDBID_PositionDB = &tmpDB

		tmpID := make(map[*models.Position]uint, 0)
		map_PositionPtr_PositionDBID = &tmpID

		for _, position := range models.AllModelStore.Positions {

			// initiate position
			var positionDB PositionDB
			positionDB.Position = *position

			query := db.Create(&positionDB)
			if query.Error != nil {
				return query.Error
			}

			// update stores
			(*map_PositionPtr_PositionDBID)[position] = positionDB.ID
			(*map_PositionDBID_PositionPtr)[positionDB.ID] = position
			(*map_PositionDBID_PositionDB)[positionDB.ID] = &positionDB
		}
	} else { // UpdateMode, update IDs of Pointer Fields of ORM object

		// check that positionStore is not nil
		if map_PositionDBID_PositionPtr == nil {
			err := errors.New("In UpdateMode translation, positionStore should not be nil")
			return err
		}

		if map_PositionDBID_PositionDB == nil {
			err := errors.New("In UpdateMode translation, positionStore should not be nil")
			return err
		}

		// update fields of positionDB with fields of position
		for _, position := range models.AllModelStore.Positions {
			positionDBID := (*map_PositionPtr_PositionDBID)[position]
			positionDB := (*map_PositionDBID_PositionDB)[positionDBID]

			positionDB.Position = *position
		}

		// parse model objects ot update associations
		for idx, position := range *map_PositionDBID_PositionPtr {

			// fetch matching positionDB
			if positionDB, ok := (*map_PositionDBID_PositionDB)[idx]; ok {
				// set {{Fieldname}}ID

				query := db.Save(&positionDB)
				if query.Error != nil {
					return query.Error
				}

			} else {
				err := errors.New(
					fmt.Sprintf("In UpdateMode translation, positionStore should not be nil %v %v",
						positionDB, position))
				return err
			}
		}
	}
	return nil
}

// stores PositionDB according to their gorm ID
var map_PositionDBID_PositionDB *map[uint]*PositionDB

// stores PositionDB ID according to Position address
var map_PositionPtr_PositionDBID *map[*models.Position]uint

// stores Position according to their gorm ID
var map_PositionDBID_PositionPtr *map[uint]*models.Position

// ORMToModelPositionTranslate is a translation function from ORM object to models objects
// This function used the uint ID of the ORM object to create or update (according to translationImpact)
// maps of respectively ORM and models objects
//
// In create mode,
func ORMToModelPositionTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that positionStores are nil

		if map_PositionDBID_PositionPtr != nil {
			err := errors.New("In CreateMode translation, Parameters positionStore should be nil")
			return err
		}

		if map_PositionDBID_PositionDB != nil {
			err := errors.New("In CreateMode translation, parameters PositionDBStore should be nil")
			return err
		}

		// init stores
		tmp := make(map[uint]*models.Position, 0)
		map_PositionDBID_PositionPtr = &tmp

		tmpDB := make(map[uint]*PositionDB, 0)
		map_PositionDBID_PositionDB = &tmpDB

		tmpID := make(map[*models.Position]uint, 0)
		map_PositionPtr_PositionDBID = &tmpID

		models.AllModelStore.Positions = make([]*models.Position, 0)

		positionDBArray := make([]PositionDB, 0)
		query := db.Find(&positionDBArray)
		if query.Error != nil {
			return query.Error
		}

		// copy orm objects to the two stores
		for _, positionDB := range positionDBArray {

			// create entries in the tree maps.
			position := positionDB.Position
			(*map_PositionDBID_PositionPtr)[positionDB.ID] = &position

			(*map_PositionPtr_PositionDBID)[&position] = positionDB.ID

			positionDBCopy := positionDB
			(*map_PositionDBID_PositionDB)[positionDB.ID] = &positionDBCopy

			// append model store with the new element
			models.AllModelStore.Positions = append(models.AllModelStore.Positions, &position)
		}
	} else { // UpdateMode
		// for later, update of the data field

		// check that positionStore is not nil
		if map_PositionDBID_PositionPtr == nil {
			err := errors.New("In UpdateMode translation, positionStore should not be nil")
			return err
		}

		if map_PositionDBID_PositionDB == nil {
			err := errors.New("In UpdateMode translation, positionStore should not be nil")
			return err
		}

		// update fields of positionDB with fields of position
		for _, position := range models.AllModelStore.Positions {
			positionDBID := (*map_PositionPtr_PositionDBID)[position]
			positionDB := (*map_PositionDBID_PositionDB)[positionDBID]

			*position = positionDB.Position
		}

		// parse all DB instance and update all pointer fields of the translated models instance
		for _, positionDB := range *map_PositionDBID_PositionDB {
			position := (*map_PositionDBID_PositionPtr)[positionDB.ID]
			if position == nil {
				err := errors.New("cannot find translated instance in models store")
				return err
			}

		}
	}

	return nil
}

func (allORMStoreStruct *AllORMStoreStruct) CreateORMPosition(position *models.Position) {

	CreateORMPosition(allORMStoreStruct.db, position)
}

// CreateORMPosition creates ORM{{Strucname}} in DB from position
func CreateORMPosition(
	db *gorm.DB,
	position *models.Position) (Error error) {

	// initiate position
	var positionDB PositionDB
	positionDB.Position = *position

	query := db.Create(&positionDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*map_PositionPtr_PositionDBID)[position] = positionDB.ID
	(*map_PositionDBID_PositionPtr)[positionDB.ID] = position
	(*map_PositionDBID_PositionDB)[positionDB.ID] = &positionDB

	return
}

func (allORMStoreStruct *AllORMStoreStruct) DeleteORMPosition(position *models.Position) {

	DeleteORMPosition(allORMStoreStruct.db, position)
}

func DeleteORMPosition(
	db *gorm.DB,
	position *models.Position) (Error error) {

	positionDBID := (*map_PositionPtr_PositionDBID)[position]
	positionDB := (*map_PositionDBID_PositionDB)[positionDBID]

	query := db.Unscoped().Delete(&positionDB)
	if query.Error != nil {
		return query.Error
	}

	delete(*map_PositionPtr_PositionDBID, position)
	delete(*map_PositionDBID_PositionPtr, positionDB.ID)
	delete(*map_PositionDBID_PositionDB, positionDBID)

	return
}
