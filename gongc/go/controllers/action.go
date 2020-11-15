package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
	"github.com/thomaspeugeot/sandbox02/gongc/go/walk"
)

// ActionInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postAction
type ActionInput struct {
	// The Action to submit or modify
	// in: body
	ActionModel *models.ActionModel
}

// PostAction swagger:route POST /actions actions postAction
// Creates a action
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: actionResponse
func PostAction(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.ActionModel

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create action
	action := models.Action{}
	action.ActionModel = input

	switch action.ActionType {
	case models.WALK:
		// traverse spinosa and fill up db
		walk.Walk(db)
	case models.DELETE_STRUCT_AND_FIELD:

		// https://gorm.io/docs/delete.html
		// if the primary key field is blank, GORM will delete all records for the model
		blankStruct := models.Struct{}
		blankField := models.Field{}
		db.Delete(&blankStruct)
		db.Delete(&blankField)
	}

	createActionDB := db.Create(&action)
	if createActionDB.Error != nil {
		// http.StatusBadRequest is a 400
		c.JSON(http.StatusBadRequest, gin.H{"error": createActionDB.Error})
		return
	}

	c.JSON(http.StatusOK, action)
}
