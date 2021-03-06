// generated by genORMModelDB.go
package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/animah/go/models"
)

// AgentAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model agentAPI
type AgentAPI struct {
	models.Agent

	// association fields

	// field Engine is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable a GORM `HAS ONE` association
	EngineID *uint

	// the associated Struct has a Name field, therefore it is generated to compute views with this relation
	EngineName string
}

// AgentDB describes a agent in the database
//
// It incorporates all fields : from the model, from the generated field for the API and the GORM ID
//
// swagger:model agentDB
type AgentDB struct {
	gorm.Model

	AgentAPI
}

// AgentDBs arrays agentDBs
// swagger:response agentDBsResponse
type AgentDBs []AgentDB

// AgentDBResponse provides response
// swagger:response agentDBResponse
type AgentDBResponse struct {
	AgentDB
}

// ModelToORMAgentTranslate is a translation function from models object to ORM objects
func ModelToORMAgentTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that agentStore is nil as well as agentDBs
		if map_AgentDBID_AgentPtr != nil {
			err := errors.New("In CreateMode translation, map_AgentDBID_AgentPtr should be nil")
			return err
		}

		if map_AgentDBID_AgentDB != nil {
			err := errors.New("In CreateMode translation, map_AgentDBID_AgentDB should be nil")
			return err
		}

		if map_AgentPtr_AgentDBID != nil {
			err := errors.New("In CreateMode translation, map_AgentPtr_AgentDBID should be nil")
			return err
		}

		tmp := make(map[uint]*models.Agent, 0)
		map_AgentDBID_AgentPtr = &tmp

		tmpDB := make(map[uint]*AgentDB, 0)
		map_AgentDBID_AgentDB = &tmpDB

		tmpID := make(map[*models.Agent]uint, 0)
		map_AgentPtr_AgentDBID = &tmpID

		for _, agent := range models.AllModelStore.Agents {

			// initiate agent
			var agentDB AgentDB
			agentDB.Agent = *agent

			query := db.Create(&agentDB)
			if query.Error != nil {
				return query.Error
			}

			// update stores
			(*map_AgentPtr_AgentDBID)[agent] = agentDB.ID
			(*map_AgentDBID_AgentPtr)[agentDB.ID] = agent
			(*map_AgentDBID_AgentDB)[agentDB.ID] = &agentDB
		}
	} else { // UpdateMode, update IDs of Pointer Fields of ORM object

		// check that agentStore is not nil
		if map_AgentDBID_AgentPtr == nil {
			err := errors.New("In UpdateMode translation, agentStore should not be nil")
			return err
		}

		if map_AgentDBID_AgentDB == nil {
			err := errors.New("In UpdateMode translation, agentStore should not be nil")
			return err
		}

		// update fields of agentDB with fields of agent
		for _, agent := range models.AllModelStore.Agents {
			agentDBID := (*map_AgentPtr_AgentDBID)[agent]
			agentDB := (*map_AgentDBID_AgentDB)[agentDBID]

			agentDB.Agent = *agent
		}

		// parse model objects ot update associations
		for idx, agent := range *map_AgentDBID_AgentPtr {

			// fetch matching agentDB
			if agentDB, ok := (*map_AgentDBID_AgentDB)[idx]; ok {
				// set {{Fieldname}}ID

				// set EngineID
				if agent.Engine != nil {
					if engineId, ok := (*map_EnginePtr_EngineDBID)[agent.Engine]; ok {
						agentDB.EngineID = &engineId
					}
				}

				query := db.Save(&agentDB)
				if query.Error != nil {
					return query.Error
				}

			} else {
				err := errors.New(
					fmt.Sprintf("In UpdateMode translation, agentStore should not be nil %v %v",
						agentDB, agent))
				return err
			}
		}
	}
	return nil
}

// stores AgentDB according to their gorm ID
var map_AgentDBID_AgentDB *map[uint]*AgentDB

// stores AgentDB ID according to Agent address
var map_AgentPtr_AgentDBID *map[*models.Agent]uint

// stores Agent according to their gorm ID
var map_AgentDBID_AgentPtr *map[uint]*models.Agent

// ORMToModelAgentTranslate is a translation function from ORM object to models objects
// This function used the uint ID of the ORM object to create or update (according to translationImpact)
// maps of respectively ORM and models objects
//
// In create mode,
func ORMToModelAgentTranslate(
	translationImpact TranslationImpact,
	db *gorm.DB) (Error error) {

	if translationImpact == CreateMode {

		// check that agentStores are nil

		if map_AgentDBID_AgentPtr != nil {
			err := errors.New("In CreateMode translation, Parameters agentStore should be nil")
			return err
		}

		if map_AgentDBID_AgentDB != nil {
			err := errors.New("In CreateMode translation, parameters AgentDBStore should be nil")
			return err
		}

		// init stores
		tmp := make(map[uint]*models.Agent, 0)
		map_AgentDBID_AgentPtr = &tmp

		tmpDB := make(map[uint]*AgentDB, 0)
		map_AgentDBID_AgentDB = &tmpDB

		tmpID := make(map[*models.Agent]uint, 0)
		map_AgentPtr_AgentDBID = &tmpID

		models.AllModelStore.Agents = make([]*models.Agent, 0)

		agentDBArray := make([]AgentDB, 0)
		query := db.Find(&agentDBArray)
		if query.Error != nil {
			return query.Error
		}

		// copy orm objects to the two stores
		for _, agentDB := range agentDBArray {

			// create entries in the tree maps.
			agent := agentDB.Agent
			(*map_AgentDBID_AgentPtr)[agentDB.ID] = &agent

			(*map_AgentPtr_AgentDBID)[&agent] = agentDB.ID

			agentDBCopy := agentDB
			(*map_AgentDBID_AgentDB)[agentDB.ID] = &agentDBCopy

			// append model store with the new element
			models.AllModelStore.Agents = append(models.AllModelStore.Agents, &agent)
		}
	} else { // UpdateMode
		// for later, update of the data field

		// check that agentStore is not nil
		if map_AgentDBID_AgentPtr == nil {
			err := errors.New("In UpdateMode translation, agentStore should not be nil")
			return err
		}

		if map_AgentDBID_AgentDB == nil {
			err := errors.New("In UpdateMode translation, agentStore should not be nil")
			return err
		}

		// update fields of agentDB with fields of agent
		for _, agent := range models.AllModelStore.Agents {
			agentDBID := (*map_AgentPtr_AgentDBID)[agent]
			agentDB := (*map_AgentDBID_AgentDB)[agentDBID]

			*agent = agentDB.Agent
		}

		// parse all DB instance and update all pointer fields of the translated models instance
		for _, agentDB := range *map_AgentDBID_AgentDB {
			agent := (*map_AgentDBID_AgentPtr)[agentDB.ID]
			if agent == nil {
				err := errors.New("cannot find translated instance in models store")
				return err
			}

			// Engine field
			if agentDB.EngineID != nil {
				agent.Engine = (*map_EngineDBID_EnginePtr)[*(agentDB.EngineID)]
			}

		}
	}

	return nil
}

func (allORMStoreStruct *AllORMStoreStruct) CreateORMAgent(agent *models.Agent) {

	CreateORMAgent(allORMStoreStruct.db, agent)
}

// CreateORMAgent creates ORM{{Strucname}} in DB from agent
func CreateORMAgent(
	db *gorm.DB,
	agent *models.Agent) (Error error) {

	// initiate agent
	var agentDB AgentDB
	agentDB.Agent = *agent

	query := db.Create(&agentDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*map_AgentPtr_AgentDBID)[agent] = agentDB.ID
	(*map_AgentDBID_AgentPtr)[agentDB.ID] = agent
	(*map_AgentDBID_AgentDB)[agentDB.ID] = &agentDB

	return
}

func (allORMStoreStruct *AllORMStoreStruct) DeleteORMAgent(agent *models.Agent) {

	DeleteORMAgent(allORMStoreStruct.db, agent)
}

func DeleteORMAgent(
	db *gorm.DB,
	agent *models.Agent) (Error error) {

	agentDBID := (*map_AgentPtr_AgentDBID)[agent]
	agentDB := (*map_AgentDBID_AgentDB)[agentDBID]

	query := db.Unscoped().Delete(&agentDB)
	if query.Error != nil {
		return query.Error
	}

	delete(*map_AgentPtr_AgentDBID, agent)
	delete(*map_AgentDBID_AgentPtr, agentDB.ID)
	delete(*map_AgentDBID_AgentDB, agentDBID)

	return
}
