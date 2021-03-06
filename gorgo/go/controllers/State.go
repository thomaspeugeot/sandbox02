// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// StateSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type StateSingloton struct {
	Callback StateCallbackInterface
}

// StateCallbackInterface is the interface that must be supported
// by the Struct that is attached to the singloton
type StateCallbackInterface interface {
	PostState(state *models.State)
}

// StateSinglotonID is the singloton variable
var StateSinglotonID StateSingloton

// An StateID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getState updateState deleteState getStateUmlscsViaStates
type StateID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StateInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postState updateState
type StateInput struct {
	// The State to submit or modify
	// in: body
	State *orm.StateAPI
}

// GetStates
//
// swagger:route GET /states states getStates
//
// Get all states
//
// Responses:
//    default: genericError
//        200: stateDBsResponse
func GetStates(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var states []orm.StateDB
	query := db.Find(&states)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, states)
}

// PostState
//
// swagger:route POST /states states postState
//
// Creates a state
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: stateDBResponse
func PostState(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.StateAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create state
	stateDB := orm.StateDB{}
	stateDB.StateAPI = input

	query := db.Create(&stateDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if StateSinglotonID.Callback != nil {
		StateSinglotonID.Callback.PostState(&(stateDB.State))
	}

	c.JSON(http.StatusOK, stateDB)
}

// GetState
//
// swagger:route GET /states/{ID} states getState
//
// Gets the details for a state.
//
// Responses:
//    default: genericError
//        200: stateDBResponse
func GetState(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get state in DB
	var state orm.StateDB
	if err := db.First(&state, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, state)
}

// UpdateState
//
// swagger:route PATCH /states/{ID} states updateState
//
// Update a state
//
// Responses:
//    default: genericError
//        200: stateDBResponse
func UpdateState(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var stateDB orm.StateDB

	// fetch the state
	query := db.First(&stateDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.StateAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	query = db.Model(&stateDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// return status OK with the marshalling of the the stateDB
	c.JSON(http.StatusOK, stateDB)
}

// DeleteState
//
// swagger:route DELETE /states/{ID} states deleteState
//
// Delete a state
//
// Responses:
//    default: genericError
func DeleteState(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var stateDB orm.StateDB
	if err := db.First(&stateDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&stateDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetStateUmlscsViaStates swagger:route GET /states/{ID}/umlscsviastates states getStateUmlscsViaStates
//
// Gets umlscs of state via field States of Umlsc.
//
// Responses:
//    default: genericError
//        200: umlscDBsResponse
func GetStateUmlscsViaStates(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get state
	var state orm.StateDB
	if err := db.First(&state, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get umlscs
	var umlscDBs orm.UmlscDBs
	columnName := gorm.ToColumnName("StatesID")
	query := db.Where(genQuery(columnName), state.ID).Find(&umlscDBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, umlscDBs)
}
