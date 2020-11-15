// generated by genController.go
package controllers

import (
	"net/http"

	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/models"
	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/orm"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// VerticeSingloton is the type of the singloton of the controllers package
// this singloton allows for the attachment of callbacks to controllers function
type VerticeSingloton struct {
	Callback VerticeCallbackInterface
}

// VerticeCallbackInterface is the interface that must be supported 
// by the Struct that is attached to the singloton
type VerticeCallbackInterface interface {
	PostVertice(vertice *models.Vertice)
}

// VerticeSinglotonID is the singloton variable
var VerticeSinglotonID VerticeSingloton

// An VerticeID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getVertice updateVertice deleteVertice getVerticeLinksViaMiddlevertice
type VerticeID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// VerticeInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postVertice updateVertice
type VerticeInput struct {
	// The Vertice to submit or modify
	// in: body
	Vertice *orm.VerticeAPI
}

// GetVertices
//
// swagger:route GET /vertices vertices getVertices
// 
// Get all vertices
//
// Responses:
//    default: genericError
//        200: verticeDBsResponse
func GetVertices(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var vertices []orm.VerticeDB
	query := db.Find(&vertices)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, vertices)
}

// PostVertice
//
// swagger:route POST /vertices vertices postVertice
// 
// Creates a vertice
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: verticeDBResponse
func PostVertice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input orm.VerticeAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create vertice
	verticeDB := orm.VerticeDB{}
	verticeDB.VerticeAPI = input

	query := db.Create(&verticeDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	if VerticeSinglotonID.Callback != nil {
		VerticeSinglotonID.Callback.PostVertice(&(verticeDB.Vertice))
	}

	c.JSON(http.StatusOK, verticeDB)
}

// GetVertice
//
// swagger:route GET /vertices/{ID} vertices getVertice
//
// Gets the details for a vertice.
//
// Responses:
//    default: genericError
//        200: verticeDBResponse
func GetVertice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get vertice in DB
	var vertice orm.VerticeDB
	if err := db.First(&vertice, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, vertice)
}

// UpdateVertice
// 
// swagger:route PATCH /vertices/{ID} vertices updateVertice
//
// Update a vertice
//
// Responses:
//    default: genericError
//        200: verticeDBResponse
func UpdateVertice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var verticeDB orm.VerticeDB

	// fetch the vertice
	query := db.First(&verticeDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Validate input
	var input orm.VerticeAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update 
	query = db.Model(&verticeDB).Updates(input)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}	

	// return status OK with the marshalling of the the verticeDB
	c.JSON(http.StatusOK, verticeDB)
}

// DeleteVertice
//
// swagger:route DELETE /vertices/{ID} vertices deleteVertice
//
// Delete a vertice
//
// Responses:
//    default: genericError
func DeleteVertice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var verticeDB orm.VerticeDB
	if err := db.First(&verticeDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&verticeDB)

	c.JSON(http.StatusOK, gin.H{"data": true})
}


// GetVerticeLinksViaMiddlevertice swagger:route GET /vertices/{ID}/linksviamiddlevertice vertices getVerticeLinksViaMiddlevertice
//
// Gets links of vertice via field Middlevertice of Link.
//
// Responses:
//    default: genericError
//        200: linkDBsResponse
func GetVerticeLinksViaMiddlevertice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get vertice
	var vertice orm.VerticeDB
	if err := db.First(&vertice, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get links
	var linkDBs orm.LinkDBs
	columnName := gorm.ToColumnName("MiddleverticeID")
	query := db.Where( genQuery(columnName), vertice.ID).Find(&linkDBs)

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	c.JSON(http.StatusOK, linkDBs)
}
