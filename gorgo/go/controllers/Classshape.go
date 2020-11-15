// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/models"
	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ClassshapeSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type ClassshapeSingloton struct {
	Callback ClassshapeCallbackInterface
}

// ClassshapeCallbackInterface is the interface that must be supported 
// by the Struct that is attached to the singloton
type ClassshapeCallbackInterface interface {
	PostClassshape(classshape *models.Classshape)
}

// ClassshapeSinglotonID is the singloton variable
var ClassshapeSinglotonID ClassshapeSingloton

// An ClassshapeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getClassshape updateClassshape deleteClassshape getClassshapePosition getClassshapeClassdiagramsViaClassshapes
type ClassshapeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClassshapeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postClassshape updateClassshape
type ClassshapeInput struct {
	// The Classshape to submit or modify
	// in: body
	Classshape *orm.ClassshapeAPI
}

// GetClassshapes
//
// swagger:route GET /classshapes classshapes getClassshapes
// 
// Get all classshapes
//
// Responses:
//    default: genericError
//        200: classshapeDBsResponse
func GetClassshapes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var classshapes []orm.ClassshapeDB
	query := db.Find(&classshapes)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classshapes)
}

// PostClassshape
//
// swagger:route POST /classshapes classshapes postClassshape
// 
// Creates a classshape
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: classshapeDBResponse
func PostClassshape(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.ClassshapeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create classshape
	classshapeDB := orm.ClassshapeDB{}
	classshapeDB.ClassshapeAPI = input

	query := db.Create(&classshapeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if ClassshapeSinglotonID.Callback != nil {
		ClassshapeSinglotonID.Callback.PostClassshape(&(classshapeDB.Classshape))
	}

	c.JSON(http.StatusOK, classshapeDB)
}

// GetClassshape
//
// swagger:route GET /classshapes/{ID} classshapes getClassshape
//
// Gets the details for a classshape.
//
// Responses:
//    default: genericError
//        200: classshapeDBResponse
func GetClassshape(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get classshape in DB
	var classshape orm.ClassshapeDB
	if err := db.First(&classshape, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classshape)
}

// UpdateClassshape
// 
// swagger:route PATCH /classshapes/{ID} classshapes updateClassshape
//
// Update a classshape
//
// Responses:
//    default: genericError
//        200: classshapeDBResponse
func UpdateClassshape(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var classshapeDB orm.ClassshapeDB

	// fetch the classshape
	query := db.First(&classshapeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ClassshapeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update 
	query = db.Model(&classshapeDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}	

	// return status OK with the marshalling of the the classshapeDB
	c.JSON(http.StatusOK, classshapeDB)
}

// DeleteClassshape
//
// swagger:route DELETE /classshapes/{ID} classshapes deleteClassshape
//
// Delete a classshape
//
// Responses:
//    default: genericError
func DeleteClassshape(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var classshapeDB orm.ClassshapeDB
	if err := db.First(&classshapeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&classshapeDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetClassshapePosition swagger:route GET /classshapes/{ID}/position classshapes getClassshapePosition
//
// Gets position of a classshape.
//
// Responses:
//    default: genericError
//        200: positionDBResponse
func GetClassshapePosition(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get classshape in DB
	var classshape orm.ClassshapeDB
	if err := db.First(&classshape, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Get position in DB
	var positionPosition orm.PositionDB
	if err := db.First(&positionPosition, *classshape.PositionID).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, positionPosition)
}


// GetClassshapeClassdiagramsViaClassshapes swagger:route GET /classshapes/{ID}/classdiagramsviaclassshapes classshapes getClassshapeClassdiagramsViaClassshapes
//
// Gets classdiagrams of classshape via field Classshapes of Classdiagram.
//
// Responses:
//    default: genericError
//        200: classdiagramDBsResponse
func GetClassshapeClassdiagramsViaClassshapes(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get classshape
	var classshape orm.ClassshapeDB
	if err := db.First(&classshape, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get classdiagrams
	var classdiagramDBs orm.ClassdiagramDBs
	columnName := gorm.ToColumnName("ClassshapesID")
	query := db.Where( genQuery(columnName), classshape.ID).Find(&classdiagramDBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classdiagramDBs)
}

