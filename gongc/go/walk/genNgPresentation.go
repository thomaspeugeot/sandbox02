package walk

import (
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgPresentationTemplateTS ...
const NgPresentationTemplateTS = `import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup} from '@angular/forms';

import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'
{{AssociationStructImportsTS}}

import { Router, RouterState, ActivatedRoute } from '@angular/router';

@Component({
	selector: 'app-{{structname}}-presentation',
	templateUrl: './{{structname}}-presentation.component.html',
	styleUrls: ['./{{structname}}-presentation.component.css']
})
export class {{Structname}}PresentationComponent implements OnInit {

	{{structname}}: {{Structname}}DB;
{{DeclsForFieldPointerToAssStructTS}}
{{AssociationStructDeclarationSlicesTS}}
{{ReversePointerToStructDeclarations}}

	constructor(
		private {{structname}}Service: {{Structname}}Service,
{{AssociationStructDeclarationServicesTS}}
		private route: ActivatedRoute,
		private router: Router,
	) {
			this.router.routeReuseStrategy.shouldReuseRoute = function () {
				return false;
			};
	}

	ngOnInit(): void {
		this.get{{Structname}}();
{{AssociationStructGetCallsTS}}

		// observable for changes in 
		this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.get{{Structname}}()
					{{AssociationStructGetCallsTS}}
				}
			}
		)
	}

  get{{Structname}}(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.{{structname}}Service.get{{Structname}}(id)
		.subscribe( 
			{{structname}} => 
			{ 
					this.{{structname}} = {{structname}}
        	}
  	);
  }

{{AssociationStructGetDefinitionsTS}}

	// set presentation outlet
	setPresentationRouterOutlet(structName :string, ID: number) {
		this.router.navigate([{
	  	outlets: {
			presentation: [structName + "-presentation", ID]
	  	}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
	 		outlets: {
	   			editor: ["{{structname}}-detail", ID]
	 	}
   	}]);
 }

}
`

// NgPresentationTemplateHTML is used for the generation
const NgPresentationTemplateHTML = `<div *ngIf="{{structname}}">
    <div class="mat-table">
        <div class="mat-header-row  pointer-to-struct-row" (click)="setEditorRouterOutlet({{structname}}.ID)">
            <div class="mat-header-cell left">{{Structname}}'s fields</div>
			<div class="mat-header-cell less-right">Values</div>
			<div class="mat-header-cell far-right"><i class="material-icons">edit</i></div>
        </div>
{{NgReadableBasicHTML}}
{{BooleanFieldFormsHTML}}
{{AssociationStructFormsHTML}}
	</div>

{{ReversePointerToStructTableComponent}}
</div>

`
const ngTemplateReversePointerToStructDeclarations = `
	{{AssocStructName}}sVia{{FieldName}}FieldName = "{{FieldName}}"; // Label used to generates the table of {{AssocStructName}} that points to {{Structname}} via {{FieldName}}
	{{AssocStructName}}sVia{{FieldName}}StructName = "{{Structname}}"; // Label used to generates the table of {{AssocStructName}} that points to {{Structname}} via {{FieldName}}`

const ngTemplateReversePointerToStructTableComponentHTML = `
	<div class="sub-table-header">{{AssocStructName}}s as {{FieldName}}</div>
	<app-{{assocStructName}}s-table [ID]="{{structname}}.ID" 
		[field]="{{AssocStructName}}sVia{{FieldName}}FieldName"
		[struct]="{{AssocStructName}}sVia{{FieldName}}StructName">
	</app-{{assocStructName}}s-table>`

const ngTemplatePresentationPointerToBoolHTML = `
<div class="mat-row">
<div class="mat-cell left">{{FieldName}}</div>
<div class="mat-cell left">
	<mat-checkbox [checked]="{{structname}}.{{FieldName}}" disabled=true></mat-checkbox>
</div>
</div>`

// NgPresentationTemplateCSS ...
const NgPresentationTemplateCSS = `
.mat-form {
  min-width: 150px;
  max-width: 500px;
  width: 100%;
}

.mat-full-width {
  width: 100%;
}

.mat-table {
	display: block;
  }
  
  .mat-row,
  .mat-header-row {
	display: flex;
	border-bottom-width: 1px;
	border-bottom-style: solid;
	border-bottom-color: rgba(0,0,0,.12);
	align-items: center;
	min-height: 48px;
	padding: 0 24px;
  }
  
  .mat-cell,
  .mat-header-cell {
	flex: 1;
	overflow: hidden;
	word-wrap: break-word;
  }

  .left {
    flex-grow: 30;
}

.right {
    flex-grow: 70;
}

.less-right {
    flex-grow: 65;
}

.far-right {
    flex-grow: 5;
}

.pointer-to-struct-row:hover {
    background-color: rgba(0, 0, 0, .05);
    cursor: pointer;
}

.sub-table-header {
    width: 100%;
    font-size: 9pt;
    text-align: center;
}
`

// GenNgPresentation generates the ngpresentation on the front
func GenNgPresentation(db *gorm.DB) {

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	// generates one presentation compenent per struct
	for _, _struct := range structs {

		structName := strings.ToLower(_struct.Name)
		stringTS := NgPresentationTemplateTS
		stringHTML := NgPresentationTemplateHTML

		// get fields
		var fields models.Fields
		query := db.Model(&_struct).Related(&fields)
		if query.Error != nil {
			log.Fatal(query.Error.Error())
		}
		sort.Slice(fields[:], func(i, j int) bool {
			return fields[i].Name < fields[j].Name
		})

		genReadableBasic(db, fields, _struct, &stringTS, &stringHTML)
		genEditableReadablePointerToStruct(db, fields, _struct, &stringTS, &stringHTML, Table)

		// generate presentation of pointer to boolean field
		BooleanFieldFormsHTML := ""
		for _, field := range fields {
			// fetch the assoc struct
			var assocStruct models.Struct
			db.First(&assocStruct, field.StructID)
			assocStructName := strings.ToLower(assocStruct.Name)

			if field.Kind == reflect.Ptr && field.AssociatedStructName == "*bool" {
				BooleanFieldFormsHTML += replace2(
					ngTemplatePresentationPointerToBoolHTML,
					"{{FieldName}}", field.Name,
					"{{assocStructName}}", assocStructName)
			}
		}

		stringHTML = strings.ReplaceAll(stringHTML, "{{BooleanFieldFormsHTML}}", BooleanFieldFormsHTML)

		// REVERSE relation

		// fetch all association fields worthy of a assocation path
		// ie. where the AssociatedStructID is the struct of interest
		var fieldWhichPointsToStruct models.Field

		fieldWhichPointsToStruct.AssociatedStructID = _struct.ID
		queryAssoc := db.Where(&fieldWhichPointsToStruct).Find(&fields)
		if queryAssoc.Error != nil {
			log.Fatal(queryAssoc.Error.Error())
		}

		ReversePointerToStructDeclarations := ""
		ReversePointerToStructTableComponent := ""
		for _, field := range fields {

			// fetch the assoc struct
			var assocStruct models.Struct
			db.First(&assocStruct, field.StructID)

			fieldName := strings.ToLower(field.Name)
			assocStructName := strings.ToLower(assocStruct.Name)
			ReversePointerToStructDeclarations +=
				replace4(ngTemplateReversePointerToStructDeclarations,
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", assocStruct.Name,
					"{{assocStructName}}", assocStructName)

			ReversePointerToStructTableComponent +=
				replace4(ngTemplateReversePointerToStructTableComponentHTML,
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", assocStruct.Name,
					"{{assocStructName}}", assocStructName)

			log.Printf("field %s", field.Name)
		}
		stringTS = strings.ReplaceAll(stringTS, "{{ReversePointerToStructDeclarations}}", ReversePointerToStructDeclarations)
		stringHTML = strings.ReplaceAll(stringHTML, "{{ReversePointerToStructTableComponent}}", ReversePointerToStructTableComponent)

		stringTS = strings.ReplaceAll(stringTS, "{{Structname}}", _struct.Name)
		stringTS = strings.ReplaceAll(stringTS, "{{structname}}", structName)
		stringTS = strings.ReplaceAll(stringTS, "{{PkgName}}", PkgName)

		fileTS, fileHTML, fileCSS := createDirAndTreeFilesInNgTargetPath(_struct.Name, "-presentation")
		defer fileTS.Close()
		defer fileHTML.Close()
		defer fileCSS.Close()

		fmt.Fprintf(fileTS, "%s", stringTS)

		stringHTML = strings.ReplaceAll(stringHTML, "{{Structname}}", _struct.Name)
		stringHTML = strings.ReplaceAll(stringHTML, "{{structname}}", structName)

		fmt.Fprintf(fileHTML, "%s", stringHTML)

		stringCSS := strings.ReplaceAll(NgPresentationTemplateCSS, "{{Structname}}", _struct.Name)
		stringCSS = strings.ReplaceAll(stringCSS, "{{structname}}", structName)

		fmt.Fprintf(fileCSS, "%s", stringCSS)

	}
}
