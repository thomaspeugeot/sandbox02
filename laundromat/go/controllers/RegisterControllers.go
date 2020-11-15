package controllers

// generated code

import (
	"github.com/gin-gonic/gin"
)

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	
	r.GET("/machines", GetMachines)
	r.GET("/machines/:id", GetMachine)
	r.POST("/machines", PostMachine)
	r.PATCH("/machines/:id", UpdateMachine)
	r.PUT("/machines/:id", UpdateMachine)
	r.DELETE("/machines/:id", DeleteMachine)
	
	r.GET("/washers", GetWashers)
	r.GET("/washers/:id", GetWasher)
	r.POST("/washers", PostWasher)
	r.PATCH("/washers/:id", UpdateWasher)
	r.PUT("/washers/:id", UpdateWasher)
	r.DELETE("/washers/:id", DeleteWasher)
	
	r.GET("/machines/:id/washersviamachine", GetMachineWashersViaMachine)
	r.GET("/washers/:id/machine", GetWasherMachine)
}
