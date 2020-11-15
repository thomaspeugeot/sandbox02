package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thomaspeugeot/sandbox02/gongc/go/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// An DiagramID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getDiagram updateDiagram deleteDiagram
type DiagramID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// DiagramInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postDiagram updateDiagram
type DiagramInput struct {
	// The Diagram to submit or modify
	// in: body
	DiagramModel *models.DiagramModel
}

// GetDiagrams swagger:route GET /diagrams diagrams getDiagrams
// Get all diagrams
//
// Responses:
//    default: genericError
//        200: diagramsResponse
func GetDiagrams(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var diagrams []models.Diagram
	db.Find(&diagrams)

	// parse paramters
	paramters := models.DiagramModel{}
	if err := c.ShouldBindQuery(&paramters); err != nil {
		c.JSON(200, err)
	}

	query := c.Request.URL.Query()
	log.Output(0, fmt.Sprintf("len of query %d", len(query)))

	c.JSON(http.StatusOK, diagrams)
}

// PostDiagram swagger:route POST /diagrams diagrams postDiagram
// Creates a diagram
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: diagramResponse
func PostDiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.DiagramModel

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create diagram
	diagram := models.Diagram{}
	diagram.DiagramModel = input

	createDiagramDB := db.Create(&diagram)
	if createDiagramDB.Error != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": createDiagramDB.Error})
		return
	}

	c.JSON(http.StatusOK, diagram)
}

// GetDiagram swagger:route GET /diagrams/{ID} diagrams getDiagram
//
// Gets the details for a diagram.
//
// Responses:
//    default: genericError
//        200: diagramResponse
func GetDiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var diagram models.Diagram
	if err := db.Where("id = ?", c.Param("id")).First(&diagram).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, diagram)
}

// UpdateDiagram swagger:route PATCH /diagrams/{ID} diagrams updateDiagram
//
// Update a diagram
//
// Responses:
//    default: genericError
//        200: diagramResponse
func UpdateDiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var diagram models.Diagram

	// fetch the diagram
	diagramToUpdate := db.Where("id = ?", c.Param("id")).First(&diagram)

	if diagramToUpdate.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.DiagramModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//
	db.Model(&diagram).Updates(input)

	// return status OK with the marshalling of the the diagram
	c.JSON(http.StatusOK, diagram)
}

// DeleteDiagram swagger:route DELETE /diagrams/{ID} diagrams deleteDiagram
//
// Delete a diagram
//
// Responses:
//    default: genericError
func DeleteDiagram(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var diagram models.Diagram
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&diagram).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&diagram)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
