package walk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const registerControllersTemplate = `package controllers

// generated code

import (
	"github.com/gin-gonic/gin"
)

// RegisterControllers register controllers
func RegisterControllers(r *gin.Engine) {
	{{Registrations}}{{AssociationControllerRegistration}}
}
`

const registrationTemplate = `
	r.GET("/{{structname}}s", Get{{Structname}}s)
	r.GET("/{{structname}}s/:id", Get{{Structname}})
	r.POST("/{{structname}}s", Post{{Structname}})
	r.PATCH("/{{structname}}s/:id", Update{{Structname}})
	r.PUT("/{{structname}}s/:id", Update{{Structname}})
	r.DELETE("/{{structname}}s/:id", Delete{{Structname}})
	`

// GenControllersRegistrations generates the setup file for the gorm
func GenControllersRegistrations(db *gorm.DB, associationRoutesAndControllers string) {

	filename := filepath.Join(ControllersPkgGenPath, "RegisterControllers.go")

	// we should use go generate
	log.Println("generating controller registering file : " + filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	registrations := fmt.Sprintf("")
	for _, _struct := range structs {

		lowerCaseStructName := strings.ToLower(_struct.Name)

		structRegistration := strings.ReplaceAll(registrationTemplate, "{{Structname}}", _struct.Name)
		structRegistration = strings.ReplaceAll(structRegistration, "{{structname}}", lowerCaseStructName)

		registrations += structRegistration
	}

	res := strings.ReplaceAll(registerControllersTemplate, "{{Registrations}}", registrations)
	res = strings.ReplaceAll(res, "{{AssociationControllerRegistration}}", associationRoutesAndControllers)

	fmt.Fprintf(f, "%s", res)

}
