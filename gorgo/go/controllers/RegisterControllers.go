package controllers

// generated code

import (
	"github.com/gin-gonic/gin"
)

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	
	r.GET("/classdiagrams", GetClassdiagrams)
	r.GET("/classdiagrams/:id", GetClassdiagram)
	r.POST("/classdiagrams", PostClassdiagram)
	r.PATCH("/classdiagrams/:id", UpdateClassdiagram)
	r.PUT("/classdiagrams/:id", UpdateClassdiagram)
	r.DELETE("/classdiagrams/:id", DeleteClassdiagram)
	
	r.GET("/classshapes", GetClassshapes)
	r.GET("/classshapes/:id", GetClassshape)
	r.POST("/classshapes", PostClassshape)
	r.PATCH("/classshapes/:id", UpdateClassshape)
	r.PUT("/classshapes/:id", UpdateClassshape)
	r.DELETE("/classshapes/:id", DeleteClassshape)
	
	r.GET("/fields", GetFields)
	r.GET("/fields/:id", GetField)
	r.POST("/fields", PostField)
	r.PATCH("/fields/:id", UpdateField)
	r.PUT("/fields/:id", UpdateField)
	r.DELETE("/fields/:id", DeleteField)
	
	r.GET("/gorgoactions", GetGorgoactions)
	r.GET("/gorgoactions/:id", GetGorgoaction)
	r.POST("/gorgoactions", PostGorgoaction)
	r.PATCH("/gorgoactions/:id", UpdateGorgoaction)
	r.PUT("/gorgoactions/:id", UpdateGorgoaction)
	r.DELETE("/gorgoactions/:id", DeleteGorgoaction)
	
	r.GET("/links", GetLinks)
	r.GET("/links/:id", GetLink)
	r.POST("/links", PostLink)
	r.PATCH("/links/:id", UpdateLink)
	r.PUT("/links/:id", UpdateLink)
	r.DELETE("/links/:id", DeleteLink)
	
	r.GET("/pkgelts", GetPkgelts)
	r.GET("/pkgelts/:id", GetPkgelt)
	r.POST("/pkgelts", PostPkgelt)
	r.PATCH("/pkgelts/:id", UpdatePkgelt)
	r.PUT("/pkgelts/:id", UpdatePkgelt)
	r.DELETE("/pkgelts/:id", DeletePkgelt)
	
	r.GET("/positions", GetPositions)
	r.GET("/positions/:id", GetPosition)
	r.POST("/positions", PostPosition)
	r.PATCH("/positions/:id", UpdatePosition)
	r.PUT("/positions/:id", UpdatePosition)
	r.DELETE("/positions/:id", DeletePosition)
	
	r.GET("/states", GetStates)
	r.GET("/states/:id", GetState)
	r.POST("/states", PostState)
	r.PATCH("/states/:id", UpdateState)
	r.PUT("/states/:id", UpdateState)
	r.DELETE("/states/:id", DeleteState)
	
	r.GET("/umlscs", GetUmlscs)
	r.GET("/umlscs/:id", GetUmlsc)
	r.POST("/umlscs", PostUmlsc)
	r.PATCH("/umlscs/:id", UpdateUmlsc)
	r.PUT("/umlscs/:id", UpdateUmlsc)
	r.DELETE("/umlscs/:id", DeleteUmlsc)
	
	r.GET("/vertices", GetVertices)
	r.GET("/vertices/:id", GetVertice)
	r.POST("/vertices", PostVertice)
	r.PATCH("/vertices/:id", UpdateVertice)
	r.PUT("/vertices/:id", UpdateVertice)
	r.DELETE("/vertices/:id", DeleteVertice)
	
	r.GET("/classdiagrams/:id/pkgeltsviaclassdiagrams", GetClassdiagramPkgeltsViaClassdiagrams)
	r.GET("/classshapes/:id/position", GetClassshapePosition)
	r.GET("/classshapes/:id/classdiagramsviaclassshapes", GetClassshapeClassdiagramsViaClassshapes)
	r.GET("/fields/:id/classshapesviafields", GetFieldClassshapesViaFields)
	r.GET("/links/:id/middlevertice", GetLinkMiddlevertice)
	r.GET("/links/:id/classshapesvialinks", GetLinkClassshapesViaLinks)
	r.GET("/positions/:id/classshapesviaposition", GetPositionClassshapesViaPosition)
	r.GET("/states/:id/umlscsviastates", GetStateUmlscsViaStates)
	r.GET("/umlscs/:id/pkgeltsviaumlscs", GetUmlscPkgeltsViaUmlscs)
	r.GET("/vertices/:id/linksviamiddlevertice", GetVerticeLinksViaMiddlevertice)
}
