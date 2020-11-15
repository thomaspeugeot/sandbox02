// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PositionSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type PositionSingloton struct {
	Callback PositionCallbackInterface
}

// PositionCallbackInterface is the interface that must be supported
// by the Struct that is attached to the singloton
type PositionCallbackInterface interface {
	PostPosition(position *models.Position)
}

// PositionSinglotonID is the singloton variable
var PositionSinglotonID PositionSingloton

// An PositionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getPosition updatePosition deletePosition getPositionClassshapesViaPosition
type PositionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// PositionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postPosition updatePosition
type PositionInput struct {
	// The Position to submit or modify
	// in: body
	Position *orm.PositionAPI
}

// GetPositions
//
// swagger:route GET /positions positions getPositions
//
// Get all positions
//
// Responses:
//    default: genericError
//        200: positionDBsResponse
func GetPositions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var positions []orm.PositionDB
	query := db.Find(&positions)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, positions)
}

// PostPosition
//
// swagger:route POST /positions positions postPosition
//
// Creates a position
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: positionDBResponse
func PostPosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.PositionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create position
	positionDB := orm.PositionDB{}
	positionDB.PositionAPI = input

	query := db.Create(&positionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if PositionSinglotonID.Callback != nil {
		PositionSinglotonID.Callback.PostPosition(&(positionDB.Position))
	}

	c.JSON(http.StatusOK, positionDB)
}

// GetPosition
//
// swagger:route GET /positions/{ID} positions getPosition
//
// Gets the details for a position.
//
// Responses:
//    default: genericError
//        200: positionDBResponse
func GetPosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get position in DB
	var position orm.PositionDB
	if err := db.First(&position, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, position)
}

// UpdatePosition
//
// swagger:route PATCH /positions/{ID} positions updatePosition
//
// Update a position
//
// Responses:
//    default: genericError
//        200: positionDBResponse
func UpdatePosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var positionDB orm.PositionDB

	// fetch the position
	query := db.First(&positionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.PositionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	query = db.Model(&positionDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// return status OK with the marshalling of the the positionDB
	c.JSON(http.StatusOK, positionDB)
}

// DeletePosition
//
// swagger:route DELETE /positions/{ID} positions deletePosition
//
// Delete a position
//
// Responses:
//    default: genericError
func DeletePosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var positionDB orm.PositionDB
	if err := db.First(&positionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&positionDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetPositionClassshapesViaPosition swagger:route GET /positions/{ID}/classshapesviaposition positions getPositionClassshapesViaPosition
//
// Gets classshapes of position via field Position of Classshape.
//
// Responses:
//    default: genericError
//        200: classshapeDBsResponse
func GetPositionClassshapesViaPosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get position
	var position orm.PositionDB
	if err := db.First(&position, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get classshapes
	var classshapeDBs orm.ClassshapeDBs
	columnName := gorm.ToColumnName("PositionID")
	query := db.Where(genQuery(columnName), position.ID).Find(&classshapeDBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classshapeDBs)
}
