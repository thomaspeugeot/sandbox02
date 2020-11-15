package controllers

// generated code

import (
	"github.com/gin-gonic/gin"
)

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	
	r.GET("/actions", GetActions)
	r.GET("/actions/:id", GetAction)
	r.POST("/actions", PostAction)
	r.PATCH("/actions/:id", UpdateAction)
	r.PUT("/actions/:id", UpdateAction)
	r.DELETE("/actions/:id", DeleteAction)
	
	r.GET("/agents", GetAgents)
	r.GET("/agents/:id", GetAgent)
	r.POST("/agents", PostAgent)
	r.PATCH("/agents/:id", UpdateAgent)
	r.PUT("/agents/:id", UpdateAgent)
	r.DELETE("/agents/:id", DeleteAgent)
	
	r.GET("/engines", GetEngines)
	r.GET("/engines/:id", GetEngine)
	r.POST("/engines", PostEngine)
	r.PATCH("/engines/:id", UpdateEngine)
	r.PUT("/engines/:id", UpdateEngine)
	r.DELETE("/engines/:id", DeleteEngine)
	
	r.GET("/agents/:id/engine", GetAgentEngine)
	r.GET("/engines/:id/agentsviaengine", GetEngineAgentsViaEngine)
}
