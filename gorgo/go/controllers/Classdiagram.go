// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/models"
	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ClassdiagramSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type ClassdiagramSingloton struct {
	Callback ClassdiagramCallbackInterface
}

// ClassdiagramCallbackInterface is the interface that must be supported 
// by the Struct that is attached to the singloton
type ClassdiagramCallbackInterface interface {
	PostClassdiagram(classdiagram *models.Classdiagram)
}

// ClassdiagramSinglotonID is the singloton variable
var ClassdiagramSinglotonID ClassdiagramSingloton

// An ClassdiagramID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getClassdiagram updateClassdiagram deleteClassdiagram getClassdiagramPkgeltsViaClassdiagrams
type ClassdiagramID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// ClassdiagramInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postClassdiagram updateClassdiagram
type ClassdiagramInput struct {
	// The Classdiagram to submit or modify
	// in: body
	Classdiagram *orm.ClassdiagramAPI
}

// GetClassdiagrams
//
// swagger:route GET /classdiagrams classdiagrams getClassdiagrams
// 
// Get all classdiagrams
//
// Responses:
//    default: genericError
//        200: classdiagramDBsResponse
func GetClassdiagrams(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var classdiagrams []orm.ClassdiagramDB
	query := db.Find(&classdiagrams)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classdiagrams)
}

// PostClassdiagram
//
// swagger:route POST /classdiagrams classdiagrams postClassdiagram
// 
// Creates a classdiagram
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: classdiagramDBResponse
func PostClassdiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.ClassdiagramAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create classdiagram
	classdiagramDB := orm.ClassdiagramDB{}
	classdiagramDB.ClassdiagramAPI = input

	query := db.Create(&classdiagramDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if ClassdiagramSinglotonID.Callback != nil {
		ClassdiagramSinglotonID.Callback.PostClassdiagram(&(classdiagramDB.Classdiagram))
	}

	c.JSON(http.StatusOK, classdiagramDB)
}

// GetClassdiagram
//
// swagger:route GET /classdiagrams/{ID} classdiagrams getClassdiagram
//
// Gets the details for a classdiagram.
//
// Responses:
//    default: genericError
//        200: classdiagramDBResponse
func GetClassdiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get classdiagram in DB
	var classdiagram orm.ClassdiagramDB
	if err := db.First(&classdiagram, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, classdiagram)
}

// UpdateClassdiagram
// 
// swagger:route PATCH /classdiagrams/{ID} classdiagrams updateClassdiagram
//
// Update a classdiagram
//
// Responses:
//    default: genericError
//        200: classdiagramDBResponse
func UpdateClassdiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var classdiagramDB orm.ClassdiagramDB

	// fetch the classdiagram
	query := db.First(&classdiagramDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.ClassdiagramAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update 
	query = db.Model(&classdiagramDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}	

	// return status OK with the marshalling of the the classdiagramDB
	c.JSON(http.StatusOK, classdiagramDB)
}

// DeleteClassdiagram
//
// swagger:route DELETE /classdiagrams/{ID} classdiagrams deleteClassdiagram
//
// Delete a classdiagram
//
// Responses:
//    default: genericError
func DeleteClassdiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var classdiagramDB orm.ClassdiagramDB
	if err := db.First(&classdiagramDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&classdiagramDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}


// GetClassdiagramPkgeltsViaClassdiagrams swagger:route GET /classdiagrams/{ID}/pkgeltsviaclassdiagrams classdiagrams getClassdiagramPkgeltsViaClassdiagrams
//
// Gets pkgelts of classdiagram via field Classdiagrams of Pkgelt.
//
// Responses:
//    default: genericError
//        200: pkgeltDBsResponse
func GetClassdiagramPkgeltsViaClassdiagrams(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get classdiagram
	var classdiagram orm.ClassdiagramDB
	if err := db.First(&classdiagram, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get pkgelts
	var pkgeltDBs orm.PkgeltDBs
	columnName := gorm.ToColumnName("ClassdiagramsID")
	query := db.Where( genQuery(columnName), classdiagram.ID).Find(&pkgeltDBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, pkgeltDBs)
}

