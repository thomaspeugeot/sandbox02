package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gongc/go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// An StructID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getStruct updateStruct deleteStruct
type StructID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// StructInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postStruct updateStruct
type StructInput struct {
	// The Struct to submit or modify
	// in: body
	StructModel *models.StructModel
}

// GetStructs swagger:route GET /structs structs getStructs
// Get all structs
//
// Responses:
//    default: genericError
//        200: structsResponse
func GetStructs(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var structs []models.Struct
	db.Find(&structs)

	// parse paramters
	paramters := models.StructModel{}
	if err := c.ShouldBindQuery(&paramters); err != nil {
		c.JSON(200, err)
	}

	query := c.Request.URL.Query()
	log.Output(0, fmt.Sprintf("len of query %d", len(query)))

	c.JSON(http.StatusOK, structs)
}

// PostStruct swagger:route POST /structs structs postStruct
// Creates a struct
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: structResponse
func PostStruct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.StructModel

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create struct
	_struct := models.Struct{}
	_struct.StructModel = input

	createStructDB := db.Create(&_struct)
	if createStructDB.Error != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": createStructDB.Error})
		return
	}

	c.JSON(http.StatusOK, _struct)
}

// GetStruct swagger:route GET /structs/{ID} structs getStruct
//
// Gets the details for a struct.
//
// Responses:
//    default: genericError
//        200: structResponse
func GetStruct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var _struct models.Struct
	if err := db.Where("id = ?", c.Param("id")).First(&_struct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, _struct)
}

// UpdateStruct swagger:route PATCH /structs/{ID} structs updateStruct
//
// Update a struct
//
// Responses:
//    default: genericError
//        200: structResponse
func UpdateStruct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var _struct models.Struct

	// fetch the _struct
	structToUpdate := db.Where("id = ?", c.Param("id")).First(&_struct)

	if structToUpdate.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.StructModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	db.Model(&_struct).Updates(input)

	// return status OK with the marshalling of the the _struct
	c.JSON(http.StatusOK, _struct)
}

// DeleteStruct swagger:route DELETE /structs/{ID} structs deleteStruct
//
// Delete a _struct
//
// Responses:
//    default: genericError
func DeleteStruct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var _struct models.Struct
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&_struct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&_struct)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
