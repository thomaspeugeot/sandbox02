package walk

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgModuleMaterialAppTemplateTS ...
const NgModuleMaterialAppTemplateTS = `import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion'; 

import { AngularSplitModule, SplitComponent } from 'angular-split';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent} from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

{{Imports}}

@NgModule({
  declarations: [

{{Declarations}}

	SplitterComponent,
	SidebarComponent
  ],
  imports: [
	FormsModule,
	ReactiveFormsModule,
    CommonModule,
    RouterModule,

    AppRoutingModule,

    MatSliderModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
	MatToolbarModule,
	MatExpansionModule,
	MatListModule,

    AngularSplitModule,
  ],
  exports: [  
	  
	{{Declarations}}

  	SplitterComponent,
  	SidebarComponent,

  ]
})
export class {{PkgName}}Module { }
`

const declatationComponentTemplate = `
	{{Structname}}sTableComponent,
	{{Structname}}AdderComponent,
	{{Structname}}DetailComponent,
	{{Structname}}PresentationComponent,
`

const importComponentTemplate = `
import { {{Structname}}sTableComponent } from './{{structname}}s-table/{{structname}}s-table.component'
import { {{Structname}}AdderComponent } from './{{structname}}-adder/{{structname}}-adder.component'
import { {{Structname}}DetailComponent } from './{{structname}}-detail/{{structname}}-detail.component'
import { {{Structname}}PresentationComponent } from './{{structname}}-presentation/{{structname}}-presentation.component'
`

const ngPackageTemplate = `{
	"$schema": "./node_modules/ng-packagr/ng-package.schema.json",
	"lib": {
	  "entryFile": "index.ts"
	}
  }
`

// GenNgMatModuleApp generates module definition of the front
func GenNgMatModuleApp(db *gorm.DB) {

	{ // create the list of structs
		var structs []models.Struct
		db.Find(&structs)

		declarations := ""
		imports := ""
		// generates one table compenent per struct
		for _, _struct := range structs {

			lowerCaseStructName := strings.ToLower(_struct.Name)

			res := ""
			res = strings.ReplaceAll(declatationComponentTemplate, "{{Structname}}", _struct.Name)
			res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
			declarations += res

			res = strings.ReplaceAll(importComponentTemplate, "{{Structname}}", _struct.Name)
			res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
			imports += res
		}

		file := createSingleFileInNgTargetPath(PkgName + ".module.ts")
		defer file.Close()

		res := strings.ReplaceAll(NgModuleMaterialAppTemplateTS, "{{Declarations}}", declarations)
		res = strings.ReplaceAll(res, "{{Imports}}", imports)
		res = strings.ReplaceAll(res, "{{PkgName}}", strings.Title(PkgName))

		fmt.Fprint(file, res)
	}
	// {
	// 	file := genTemlateFile("ng-package.json")
	// 	defer file.Close()
	// 	fmt.Fprint(file, ngPackageTemplate)
	// }

}
