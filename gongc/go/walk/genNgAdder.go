package walk

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgAdderTemplateTS ...
const NgAdderTemplateTS = `import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'
{{AssociationStructImportsTS}}

@Component({
  selector: 'app-{{structname}}-adder',
  templateUrl: './{{structname}}-adder.component.html',
  styleUrls: ['./{{structname}}-adder.component.css']
})
export class {{Structname}}AdderComponent implements OnInit {

	{{structname}} = {} as {{Structname}}DB;

{{DeclsForFieldPointerToAssStructTS}}
{{AssociationStructDeclarationSlicesTS}}
{{BooleanFieldFormControlDeclarations}}


  constructor(
    private {{structname}}Service: {{Structname}}Service, 
	{{AssociationStructDeclarationServicesTS}}  ) {
  }

  ngOnInit(): void {
{{AssociationStructGetCallsTS}}
  }

{{AssociationStructGetDefinitionsTS}}
  add(): void {
{{BooleanFieldFormRecoveries}}
{{UpdatesForFieldPointerToAssStructTS}}

    this.{{structname}}Service.post{{Structname}}( this.{{structname}} )
    .subscribe({{structname}} => {
		this.{{structname}}Service.{{Structname}}ServiceChanged.next("post")
		
		this.{{structname}} = {} // reset fields
	    console.log("{{structname}} added")
    });
  }
}
`

const ngAdderTemplateHTML = `<div *ngIf="{{structname}}">
	<h3>New {{Structname}}</h3>
{{NgEditableBasicInputFormHTML}}
{{BooleanFieldFormsHTML}}
{{AssociationStructFormsHTML}}
	<button mat-raised-button 
	(click)="add()">
	Add
	</button>
</div>
`

// NgAdderTemplateCSS ...
const NgAdderTemplateCSS = `
.mat-form {
  min-width: 150px;
  max-width: 500px;
  width: 100%;
}

.mat-full-width {
  width: 100%;
}
`

// GenNgAdder generates the ng-adder on the front
func GenNgAdder(db *gorm.DB) {

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	// generates one adder compenent per struct
	for _, _struct := range structs {

		structName := strings.ToLower(_struct.Name)

		// get fields
		var fields models.Fields
		query := db.Model(&_struct).Related(&fields)
		if query.Error != nil {
			log.Fatal(query.Error.Error())
		}
		sort.Slice(fields[:], func(i, j int) bool {
			return fields[i].Name < fields[j].Name
		})

		stringTS := NgAdderTemplateTS
		stringHTML := ngAdderTemplateHTML

		genEditableBasic(db, fields, _struct, &stringTS, &stringHTML)
		// genEditableTime(db, fields, _struct, &stringTS, &stringHTML)
		genEditablePointerToBool(db, fields, _struct, &stringTS, &stringHTML)
		genEditableReadablePointerToStruct(db, fields, _struct, &stringTS, &stringHTML, Editable)

		stringTS = strings.ReplaceAll(stringTS, "{{Structname}}", _struct.Name)
		stringTS = strings.ReplaceAll(stringTS, "{{structname}}", structName)
		stringTS = strings.ReplaceAll(stringTS, "{{PkgName}}", PkgName)

		stringHTML = strings.ReplaceAll(stringHTML, "{{Structname}}", _struct.Name)
		stringHTML = strings.ReplaceAll(stringHTML, "{{structname}}", structName)

		stringCSS := strings.ReplaceAll(NgAdderTemplateCSS, "{{Structname}}", _struct.Name)
		stringCSS = strings.ReplaceAll(stringCSS, "{{structname}}", structName)

		fileTS, fileHTML, fileCSS := createDirAndTreeFilesInNgTargetPath(_struct.Name, "-adder")
		defer fileTS.Close()
		defer fileHTML.Close()
		defer fileCSS.Close()
		fmt.Fprintf(fileTS, "%s", stringTS)
		fmt.Fprintf(fileHTML, "%s", stringHTML)
		fmt.Fprintf(fileCSS, "%s", stringCSS)

	}
}
