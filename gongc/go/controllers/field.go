package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gongc/go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// An FieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getField updateField deleteField
type FieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postField updateField
type FieldInput struct {
	// The Field to submit or modify
	// in: body
	FieldModel *models.FieldModel
}

// GetFields swagger:route GET /fields fields getFields
// Get all fields
//
// Responses:
//    default: genericError
//        200: fieldsResponse
func GetFields(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var fields []models.Field
	db.Find(&fields)

	// parse paramters
	paramters := models.FieldModel{}
	if err := c.ShouldBindQuery(&paramters); err != nil {
		c.JSON(200, err)
	}

	query := c.Request.URL.Query()
	log.Output(0, fmt.Sprintf("len of query %d", len(query)))

	c.JSON(http.StatusOK, fields)
}

// PostField swagger:route POST /fields fields postField
// Creates a field
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: fieldResponse
func PostField(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.FieldModel

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create field
	field := models.Field{}
	field.FieldModel = input

	createFieldDB := db.Create(&field)
	if createFieldDB.Error != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": createFieldDB.Error})
		return
	}

	c.JSON(http.StatusOK, field)
}

// GetField swagger:route GET /fields/{ID} fields getField
//
// Gets the details for a field.
//
// Responses:
//    default: genericError
//        200: fieldResponse
func GetField(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var field models.Field
	if err := db.Where("id = ?", c.Param("id")).First(&field).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, field)
}

// UpdateField swagger:route PATCH /fields/{ID} fields updateField
//
// Update a field
//
// Responses:
//    default: genericError
//        200: fieldResponse
func UpdateField(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var field models.Field

	// fetch the field
	fieldToUpdate := db.Where("id = ?", c.Param("id")).First(&field)

	if fieldToUpdate.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.FieldModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	db.Model(&field).Updates(input)

	// return status OK with the marshalling of the the field
	c.JSON(http.StatusOK, field)
}

// DeleteField swagger:route DELETE /fields/{ID} fields deleteField
//
// Delete a field
//
// Responses:
//    default: genericError
func DeleteField(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var field models.Field
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&field).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&field)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
