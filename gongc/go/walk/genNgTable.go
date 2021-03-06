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

// NgTableTemplateTS ...
const NgTableTemplateTS = `// generated by genNgTable.go
import { Component, OnInit, OnChanges, Input, Output, EventEmitter } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { Router, RouterState } from '@angular/router';
import { {{Structname}}DB } from '../{{structname}}-db'
import { {{Structname}}Service } from '../{{structname}}.service'
{{AssociationStructImportsTS}}

// generated table component
@Component({
  selector: 'app-{{structname}}s-table',
  templateUrl: './{{structname}}s-table.component.html',
  styleUrls: ['./{{structname}}s-table.component.css']
})
export class {{Structname}}sTableComponent implements OnInit {

  // the data source for the table
  {{structname}}s: {{Structname}}DB[];

  @Input() ID : number; // ID of the caller when component called from struct in reverse relation
  @Input() struct : string; // struct with pointer to {{Structname}}
  @Input() field : string; // field to display

  displayedColumns: string[] = ['ID', {{FieldColmunsTS}}'Edit', 'Delete'];

  constructor(
    private {{structname}}Service: {{Structname}}Service,
{{AssociationStructDeclarationServicesTS}}
    private router: Router,
  ) {
    // observable for changes in structs
    this.{{structname}}Service.{{Structname}}ServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.get{{Structname}}s()
        }
      }
    )
  }

  ngOnInit(): void {
    this.get{{Structname}}s()
  }

  get{{Structname}}s(): void {
    if (this.ID == null) {
      this.{{structname}}Service.get{{Structname}}s().subscribe(
        {{Structname}}s => {
          this.{{structname}}s = {{Structname}}s;
        }
      )
    }
  {{ReversePointerToStructGetTS}}
  }

  // new{{Structname}} initiate a new {{structname}}
  // create a new {{Structname}} objet
  new{{Structname}}() {
  }

  delete{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {
    // la liste des {{structname}}s est amputée du {{structname}} avant le delete afin
    // de mettre à jour l'IHM
    this.{{structname}}s = this.{{structname}}s.filter(h => h !== {{structname}});

    this.{{structname}}Service.delete{{Structname}}({{structname}}ID).subscribe();
  }

  edit{{Structname}}({{structname}}ID: number, {{structname}}: {{Structname}}DB) {

  }

  // display {{structname}} in router
  display{{Structname}}InRouter({{structname}}ID: number) {
    this.router.navigate( ["{{structname}}-display", {{structname}}ID])
  }

  // set editor outlet
  setEditorRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        editor: ["{{structname}}-detail", {{structname}}ID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet({{structname}}ID: number) {
    this.router.navigate([{
      outlets: {
        presentation: ["{{structname}}-presentation", {{structname}}ID]
      }
    }]);
  }
}
`

// NgTableTemplateHTML is used for the generation
const NgTableTemplateHTML = `<table mat-table [dataSource]="{{structname}}s">

    <ng-container matColumnDef="ID">
        <th mat-header-cell *matHeaderCellDef> ID. </th>
        <td mat-cell *matCellDef="let {{Structname}}"> {{{{Structname}}.ID}} </td>
    </ng-container>
{{FieldColmunsHTML}}
    <ng-container matColumnDef="Edit">
    <th mat-header-cell *matHeaderCellDef> Edit </th>
    <td mat-cell *matCellDef="let {{structname}};  let j = index;">
        <i class="material-icons" [ngStyle]="{'color':'rgba(0,0,0,.50)'}" (click)="setEditorRouterOutlet({{structname}}.ID)">edit</i>
    </td>
    </ng-container>

    <ng-container matColumnDef="Delete">
        <th mat-header-cell *matHeaderCellDef> Delete </th>
        <td mat-cell *matCellDef="let {{structname}};  let j = index;">
            <i class="material-icons" [ngStyle]="{'color':'rgba(0,0,0,.50)'}" (click)="delete{{Structname}}({{structname}}.ID, {{structname}})">delete</i>
        </td>
    </ng-container>

    <tr mat-header-row *matHeaderRowDef="displayedColumns"></tr>

    <tr mat-row *matRowDef="
    let row; 
    columns: displayedColumns;
    " 
    (click)="setPresentationRouterOutlet( row.ID ) "
    class="row-link">
    </tr>
</table>
`

// NgColumnTemplateHTML ...
const NgColumnTemplateHTML = `
    <ng-container matColumnDef="{{FieldName}}">
        <th mat-header-cell *matHeaderCellDef> {{FieldNameOrStructName}} </th>
        <td mat-cell *matCellDef="let {{Structname}}">
            {{{{Structname}}.{{FieldName}}}}
        </td>
    </ng-container>
`

// NgFloat64ColumnTemplateHTML ...
const NgFloat64ColumnTemplateHTML = `
    <ng-container matColumnDef="{{FieldName}}">
        <th mat-header-cell *matHeaderCellDef> {{FieldNameOrStructName}} </th>
        <td mat-cell *matCellDef="let {{Structname}}">
            {{{{Structname}}.{{FieldName}}.toPrecision(5)}}
        </td>
    </ng-container>
`

const ngTemplateFieldPointerHTML = `
    <ng-container matColumnDef="{{FieldName}}">
        <th mat-header-cell *matHeaderCellDef> {{FieldName}} </th>
        <td mat-cell *matCellDef="let {{Structname}}">
            {{{{Structname}}.{{FieldName}}Name}}
        </td>
    </ng-container>
`

// NgBooleanColumnTemplateHTML ...
const NgBooleanColumnTemplateHTML = `
    <ng-container matColumnDef="{{FieldName}}">
        <th mat-header-cell *matHeaderCellDef> {{FieldName}} </th>
        <td mat-cell *matCellDef="let {{Structname}}">
            <mat-checkbox [checked]="{{Structname}}.{{FieldName}}" disabled=true></mat-checkbox>
        </td>
    </ng-container>
`

// NgTableTemplateCSS ...
const NgTableTemplateCSS = `
table {
    width: 100%;
}

section {
  display: table;
  margin: 8px;
}

button {
  background-color: #eee;
  border: none;
  padding: 5px 10px;
  border-radius: 4px;
  cursor: pointer;
  cursor: hand;
  font-family: Arial;
}

button:hover {
  background-color: #cfd8dc;
}

button.delete {
  position: relative;
  /* left: 194px; */
  left: 510px;
  top: -32px;
  background-color: gray !important;
  color: white;
}

.mat-row-link{
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;          
     
}

.row-link:hover {
  background-color: rgba(0, 0, 0, .05);
  cursor: pointer;
}

.material-icons {
  font-family: 'Material Icons';
  font-weight: normal;
  font-style: normal;
  font-size: 24px;  /* Preferred icon size */
  display: inline-block;
  line-height: 1;
  text-transform: none;
  letter-spacing: normal;
  word-wrap: normal;
  white-space: nowrap;
  direction: ltr;

  /* Support for all WebKit browsers. */
  -webkit-font-smoothing: antialiased;
  /* Support for Safari and Chrome. */
  text-rendering: optimizeLegibility;

  /* Support for Firefox. */
  -moz-osx-font-smoothing: grayscale;

  /* Support for IE. */
  font-feature-settings: 'liga';
}
`

// GenNgTable generates the ngtable on the front
func GenNgTable(db *gorm.DB) {

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	// generates one table compenent per struct
	for _, _struct := range structs {

		// compute {{FieldColmunsTS}}
		FieldColmunsTS := ""
		FieldColmunsHTML := ""

		lowerCaseStructName := strings.ToLower(_struct.Name)

		// get fields
		var fields models.Fields
		query := db.Model(&_struct).Related(&fields)
		if query.Error != nil {
			log.Fatal(query.Error.Error())
		}
		sort.Slice(fields[:], func(i, j int) bool {
			return fields[i].Name < fields[j].Name
		})

		for _, field := range fields {
			if field.Kind != reflect.Ptr && field.Kind != reflect.Slice &&
				(field.Kind == reflect.Int || field.Kind == reflect.String) {
				FieldColmunsTS += fmt.Sprintf("'%s', ", field.Name)

				res := strings.ReplaceAll(NgColumnTemplateHTML, "{{Structname}}", _struct.Name)
				res = strings.ReplaceAll(res, "{{FieldName}}", field.Name)
				if field.Name == "Name" {
					res = strings.ReplaceAll(res, "{{FieldNameOrStructName}}", _struct.Name)
				} else {
					res = strings.ReplaceAll(res, "{{FieldNameOrStructName}}", field.Name)
				}

				FieldColmunsHTML += res
			}
			if field.Kind != reflect.Ptr && field.Kind != reflect.Slice &&
				(field.Kind == reflect.Float64) {
				FieldColmunsTS += fmt.Sprintf("'%s', ", field.Name)

				res := strings.ReplaceAll(NgFloat64ColumnTemplateHTML, "{{Structname}}", _struct.Name)
				res = strings.ReplaceAll(res, "{{FieldName}}", field.Name)
				if field.Name == "Name" {
					res = strings.ReplaceAll(res, "{{FieldNameOrStructName}}", _struct.Name)
				} else {
					res = strings.ReplaceAll(res, "{{FieldNameOrStructName}}", field.Name)
				}

				FieldColmunsHTML += res
			}
			if field.Kind == reflect.Ptr && field.AssociatedStructName == "*bool" {
				FieldColmunsTS += fmt.Sprintf("'%s', ", field.Name)

				res := strings.ReplaceAll(NgBooleanColumnTemplateHTML, "{{Structname}}", _struct.Name)
				res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
				res = strings.ReplaceAll(res, "{{FieldName}}", field.Name)

				FieldColmunsHTML += res
			}

			if field.Kind == reflect.Ptr && field.AssociatedStructID != 0 {
				FieldColmunsTS += fmt.Sprintf("'%s', ", field.Name)

				res := strings.ReplaceAll(ngTemplateFieldPointerHTML, "{{Structname}}", _struct.Name)
				res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)
				res = strings.ReplaceAll(res, "{{FieldName}}", field.Name)

				FieldColmunsHTML += res

				// the compoent can be initialised with some input when it is a reverse relation
			}
		}

		fileTS, fileHTML, fileCSS := createDirAndTreeFilesInNgTargetPath(_struct.Name, "s-table")
		defer fileTS.Close()
		defer fileHTML.Close()
		defer fileCSS.Close()

		stringTS := NgTableTemplateTS
		stringHTML := NgTableTemplateHTML

		// for pointer to struct, some pattern from the detail components are usefull
		genEditableReadablePointerToStruct(db, fields, _struct, &stringTS, &stringHTML, Table)

		stringTS = strings.ReplaceAll(stringTS, "{{Structname}}", _struct.Name)
		stringTS = strings.ReplaceAll(stringTS, "{{structname}}", lowerCaseStructName)
		stringTS = strings.ReplaceAll(stringTS, "{{PkgName}}", PkgName)
		stringTS = strings.ReplaceAll(stringTS, "{{FieldColmunsTS}}", FieldColmunsTS)

		fmt.Fprintf(fileTS, "%s", stringTS)

		stringHTML = strings.ReplaceAll(stringHTML, "{{Structname}}", _struct.Name)
		stringHTML = strings.ReplaceAll(stringHTML, "{{structname}}", lowerCaseStructName)
		stringHTML = strings.ReplaceAll(stringHTML, "{{structname}}", lowerCaseStructName)
		stringHTML = strings.ReplaceAll(stringHTML, "{{FieldColmunsHTML}}", FieldColmunsHTML)

		fmt.Fprintf(fileHTML, "%s", stringHTML)

		stringCSS := strings.ReplaceAll(NgTableTemplateCSS, "{{Structname}}", _struct.Name)
		stringCSS = strings.ReplaceAll(stringCSS, "{{structname}}", lowerCaseStructName)

		fmt.Fprintf(fileCSS, "%s", stringCSS)

	}
}
