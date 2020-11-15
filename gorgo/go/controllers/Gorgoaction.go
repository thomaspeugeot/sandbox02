// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GorgoactionSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type GorgoactionSingloton struct {
	Callback GorgoactionCallbackInterface
}

// GorgoactionCallbackInterface is the interface that must be supported
// by the Struct that is attached to the singloton
type GorgoactionCallbackInterface interface {
	PostGorgoaction(gorgoaction *models.Gorgoaction)
}

// GorgoactionSinglotonID is the singloton variable
var GorgoactionSinglotonID GorgoactionSingloton

// An GorgoactionID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getGorgoaction updateGorgoaction deleteGorgoaction
type GorgoactionID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// GorgoactionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postGorgoaction updateGorgoaction
type GorgoactionInput struct {
	// The Gorgoaction to submit or modify
	// in: body
	Gorgoaction *orm.GorgoactionAPI
}

// GetGorgoactions
//
// swagger:route GET /gorgoactions gorgoactions getGorgoactions
//
// Get all gorgoactions
//
// Responses:
//    default: genericError
//        200: gorgoactionDBsResponse
func GetGorgoactions(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var gorgoactions []orm.GorgoactionDB
	query := db.Find(&gorgoactions)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, gorgoactions)
}

// PostGorgoaction
//
// swagger:route POST /gorgoactions gorgoactions postGorgoaction
//
// Creates a gorgoaction
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: gorgoactionDBResponse
func PostGorgoaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.GorgoactionAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create gorgoaction
	gorgoactionDB := orm.GorgoactionDB{}
	gorgoactionDB.GorgoactionAPI = input

	query := db.Create(&gorgoactionDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if GorgoactionSinglotonID.Callback != nil {
		GorgoactionSinglotonID.Callback.PostGorgoaction(&(gorgoactionDB.Gorgoaction))
	}

	c.JSON(http.StatusOK, gorgoactionDB)
}

// GetGorgoaction
//
// swagger:route GET /gorgoactions/{ID} gorgoactions getGorgoaction
//
// Gets the details for a gorgoaction.
//
// Responses:
//    default: genericError
//        200: gorgoactionDBResponse
func GetGorgoaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get gorgoaction in DB
	var gorgoaction orm.GorgoactionDB
	if err := db.First(&gorgoaction, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, gorgoaction)
}

// UpdateGorgoaction
//
// swagger:route PATCH /gorgoactions/{ID} gorgoactions updateGorgoaction
//
// Update a gorgoaction
//
// Responses:
//    default: genericError
//        200: gorgoactionDBResponse
func UpdateGorgoaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var gorgoactionDB orm.GorgoactionDB

	// fetch the gorgoaction
	query := db.First(&gorgoactionDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.GorgoactionAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update
	query = db.Model(&gorgoactionDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// return status OK with the marshalling of the the gorgoactionDB
	c.JSON(http.StatusOK, gorgoactionDB)
}

// DeleteGorgoaction
//
// swagger:route DELETE /gorgoactions/{ID} gorgoactions deleteGorgoaction
//
// Delete a gorgoaction
//
// Responses:
//    default: genericError
func DeleteGorgoaction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var gorgoactionDB orm.GorgoactionDB
	if err := db.First(&gorgoactionDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&gorgoactionDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
