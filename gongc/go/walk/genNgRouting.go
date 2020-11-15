package walk

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgRoutingTemplateTS ...
const NgRoutingTemplateTS = `import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

{{Imports}}
const routes: Routes = [{{Routings}}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
`

const routingTemplate = `
	{ path: '{{structname}}s', component: {{Structname}}sTableComponent, outlet: 'table' },
	{ path: '{{structname}}-adder', component: {{Structname}}AdderComponent, outlet: 'editor' },
	{ path: '{{structname}}-detail/:id', component: {{Structname}}DetailComponent, outlet: 'editor' },
	{ path: '{{structname}}-presentation/:id', component: {{Structname}}PresentationComponent, outlet: 'presentation' },
	{ path: '{{structname}}-presentation-special/:id', component: {{Structname}}PresentationComponent, outlet: '{{structname}}pres' },
`

// GenNgRouting generates the routing on the front
func GenNgRouting(db *gorm.DB) {

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	routings := ""
	imports := ""
	// generates one table compenent per struct
	for _, _struct := range structs {

		lowerCaseStructName := strings.ToLower(_struct.Name)

		res := ""
		res = strings.ReplaceAll(routingTemplate, "{{Structname}}", _struct.Name)
		res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
		routings += res

		res = strings.ReplaceAll(importComponentTemplate, "{{Structname}}", _struct.Name)
		res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
		imports += res

	}

	file := createSingleFileInNgTargetPath("app-routing.module.ts")
	defer file.Close()

	res := strings.ReplaceAll(NgRoutingTemplateTS, "{{Routings}}", routings)
	res = strings.ReplaceAll(res, "{{Imports}}", imports)
	res = strings.ReplaceAll(res, "{{PkgName}}", PkgName)

	fmt.Fprint(file, res)
}
