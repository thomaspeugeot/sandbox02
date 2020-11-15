package walk

import (
	"log"
	"reflect"
	"sort"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NamesUpdatesOfAssociatedStructTemaplateTS
const NamesUpdatesOfAssociatedStructTemaplateTS = `
`

const ngTemplateAssociationStructInitPointerTS = `
		if (editor.ID == this.book.{{FieldName}}ID) {
			this.{{FieldName}} = editor
		}
`

const ngTemplateAssociationStructImportTS = `
import { {{AssocStructName}}API} from '../{{assocStructName}}-api'
import { {{AssocStructName}}DB} from '../{{assocStructName}}-db'
import { {{AssocStructName}}Service} from '../{{assocStructName}}.service'
`

const ngTemplateAssociationStructDeclarationSliceTS = `

	// generated by genEditableReadablePointerToStruct.go
	{{assocStructName}}s: {{AssocStructName}}DB[];`

const ngTemplateAssociationStructDeclarationServiceTS = `
		private {{assocStructName}}Service: {{AssocStructName}}Service,`

const ngTemplateAssociationStructGetCallTS = `
    	this.get{{AssocStructName}}s();
`
const ngTemplateAssociationStructGetDefinitionTS = `
	// generated by genEditableReadablePointerToStruct.go
	get{{AssocStructName}}s(): void {
		this.{{assocStructName}}Service.get{{AssocStructName}}s().subscribe(
			{{assocStructName}}s => {
				this.{{assocStructName}}s = {{assocStructName}}s;

				// init variable for each pointer
				this.{{assocStructName}}s.forEach({{assocStructName}} => {
					{{InitsForFieldPointerToAssStructTS}}
				});
      		}
    	)
	}`

const ngTemplateDeclForFieldPointerToAssStructTS = `
	{{FieldName}} = {} as {{AssocStructName}}DB; // storing values of the field {{FieldName}} of type {{AssocStructName}}`

const ngTemplateInitForFieldPointerToAssStructTS = `if ({{assocStructName}}.ID == this.{{structname}}.{{FieldName}}ID) {
						this.{{FieldName}} = {{assocStructName}}
					}`

const ngTemplateUpdateForFieldPointerToAssStructTS = `
	this.{{structname}}.{{FieldName}}ID = this.{{FieldName}}.ID;
	this.{{structname}}.{{FieldName}}Name = this.{{FieldName}}.Name;`

const ngTemplateEditablePointerToStructFormHTML = `
    <mat-form-field>
        <mat-label>{{FieldName}}</mat-label>
        <mat-select [(ngModel)]="{{FieldName}}">
            <mat-option *ngFor="let {{assocStructName}} of {{assocStructName}}s" [value]="{{assocStructName}}">
                {{{{assocStructName}}.Name}}
            </mat-option>
        </mat-select>
    </mat-form-field>
`

const ngTemplateBasicPointerToStructFormHTML = `
    <mat-form-field>
        <mat-label>{{FieldName}}</mat-label>
        <mat-select [(ngModel)]="{{FieldName}}">
            <mat-option *ngFor="let {{assocStructName}} of {{assocStructName}}s" [value]="{{assocStructName}}">
                {{{{assocStructName}}.Name}}
            </mat-option>
        </mat-select>
    </mat-form-field>
`

const ngTemplatReversePointerToStructGetTS = `
		if (this.ID != 0 && this.field == "{{FieldName}}" && this.struct == "{{AssocStructName}}") {
			this.{{assocStructName}}Service.get{{AssocStructName}}{{Structname}}sVia{{FieldName}}(this.ID).subscribe(
				{{structname}}s => {
					this.{{structname}}s = {{structname}}s;
				}
			)
		}`

type PointerToStructFiledGenerationType int

const (
	Editable = iota
	Readable
	Table
)

func genEditableReadablePointerToStruct(db *gorm.DB, fields models.Fields, _struct models.Struct,
	stringTS *string, stringHTML *string,
	generationType PointerToStructFiledGenerationType) {

	AssociationStructImportsTS := ""
	AssociationStructDeclarationSlicesTS := ""
	AssociationStructDeclarationServicesTS := ""
	AssociationStructGetCallsTS := ""
	AssociationStructFormsHTML := ""

	NamesUpdatesOfAssociatedStructTS := ""

	//  generated code for association structs
	// an "association struct" is a struct that is the type of one or more field of the _struct
	DeclsForFieldPointerToAssStructTS := "" // generated code for the declaration of the field

	// generated code for a getter function of all instances of the association struct in the database
	// when the function receive the values, each field is instanciated with its corresponding value
	// NOTE : this part has to be optimized with direct getStructID() functions
	AssociationStructGetFunctionsTS := ""

	// generated code for update functions
	UpdatesForFieldPointerToAssStructTS := ""

	ReversePointerToStructGetTS := ""

	// TODO #51 when multiple field have the same association struct
	assocStructMatch := make(map[uint]bool)

	// map of getter function of associated struct
	assocStructInitPerStructStrings := make(map[string]string)

	// map of code that initiate a field value within the getter function
	assocStructInitPerFieldStrings := make(map[string]string)

	for _, field := range fields {
		if field.Kind == reflect.Ptr && field.AssociatedStructID != 0 && field.AssociatedStructID != _struct.ID {
			// fetch the assoc struct
			var assocStruct models.Struct
			db.First(&assocStruct, field.AssociatedStructID)

			fieldName := strings.ToLower(field.Name)
			assocStructName := strings.ToLower(assocStruct.Name)

			_, ok := assocStructMatch[field.AssociatedStructID]
			if !ok {

				assocStructMatch[field.AssociatedStructID] = true

				// case when the is an association with itself
				if assocStruct.Name != _struct.Name {
					AssociationStructImportsTS += replace2(ngTemplateAssociationStructImportTS,
						"{{AssocStructName}}", assocStruct.Name, "{{assocStructName}}", assocStructName)

					AssociationStructDeclarationServicesTS += replace2(ngTemplateAssociationStructDeclarationServiceTS,
						"{{AssocStructName}}", assocStruct.Name, "{{assocStructName}}", assocStructName)
				}

				AssociationStructDeclarationSlicesTS += replace2(ngTemplateAssociationStructDeclarationSliceTS,
					"{{AssocStructName}}", assocStruct.Name, "{{assocStructName}}", assocStructName)

				AssociationStructGetCallsTS += replace2(ngTemplateAssociationStructGetCallTS,
					"{{AssocStructName}}", assocStruct.Name, "{{assocStructName}}", assocStructName)

				assocStructInitPerStructStrings[field.AssociatedStructName] = replace2(ngTemplateAssociationStructGetDefinitionTS,
					"{{AssocStructName}}", assocStruct.Name, "{{assocStructName}}", assocStructName)

				assocStructInitPerFieldStrings[field.AssociatedStructName] = ""

			}
			ReversePointerToStructGetTS += replace4(ngTemplatReversePointerToStructGetTS,
				"{{fieldName}}", fieldName,
				"{{FieldName}}", field.Name,
				"{{AssocStructName}}", assocStruct.Name,
				"{{assocStructName}}", assocStructName)

			switch generationType {
			case Editable:
				AssociationStructFormsHTML += replace2(ngTemplateEditablePointerToStructFormHTML,
					"{{FieldName}}", field.Name, "{{assocStructName}}", assocStructName)
			case Readable:
				AssociationStructFormsHTML += replace4(ngReadablePointerToStructFieldHTML,
					"{{fieldName}}", fieldName,
					"{{FieldName}}", field.Name,
					"{{AssocStructName}}", assocStruct.Name,
					"{{assocStructName}}", assocStructName)
			case Table:
			default:
				log.Panic("Unknown generation type")
			}

			// declarations of fields that match each pointer to a metabaron type Struct
			DeclsForFieldPointerToAssStructTS += replace2(ngTemplateDeclForFieldPointerToAssStructTS,
				"{{FieldName}}", field.Name, "{{AssocStructName}}", assocStruct.Name)

			assocStructInitPerFieldStrings[field.AssociatedStructName] += replace4(ngTemplateInitForFieldPointerToAssStructTS,
				"{{fieldName}}", fieldName,
				"{{FieldName}}", field.Name,
				"{{AssocStructName}}", assocStruct.Name,
				"{{assocStructName}}", assocStructName)

			UpdatesForFieldPointerToAssStructTS += replace4(ngTemplateUpdateForFieldPointerToAssStructTS,
				"{{fieldName}}", fieldName,
				"{{FieldName}}", field.Name,
				"{{AssocStructName}}", assocStruct.Name,
				"{{assocStructName}}", assocStructName)

		}
	}
	*stringTS = strings.ReplaceAll(*stringTS, "{{AssociationStructImportsTS}}", AssociationStructImportsTS)
	*stringTS = strings.ReplaceAll(*stringTS, "{{AssociationStructDeclarationSlicesTS}}", AssociationStructDeclarationSlicesTS)
	*stringTS = strings.ReplaceAll(*stringTS, "{{AssociationStructDeclarationServicesTS}}", AssociationStructDeclarationServicesTS)
	*stringTS = strings.ReplaceAll(*stringTS, "{{AssociationStructGetCallsTS}}", AssociationStructGetCallsTS)

	*stringTS = strings.ReplaceAll(*stringTS, "{{NamesUpdatesOfAssociatedStructTS}}", NamesUpdatesOfAssociatedStructTS)

	*stringTS = strings.ReplaceAll(*stringTS, "{{DeclsForFieldPointerToAssStructTS}}", DeclsForFieldPointerToAssStructTS)

	keys := make([]string, 0, len(assocStructInitPerStructStrings))
	for k := range assocStructInitPerStructStrings {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, assocStructName := range keys {
		assocGetStructTS := assocStructInitPerStructStrings[assocStructName]
		AssociationStructGetFunctionsTS += strings.ReplaceAll(assocGetStructTS,
			"{{InitsForFieldPointerToAssStructTS}}",
			assocStructInitPerFieldStrings[assocStructName])
	}
	*stringTS = strings.ReplaceAll(*stringTS, "{{AssociationStructGetDefinitionsTS}}", AssociationStructGetFunctionsTS)

	*stringTS = strings.ReplaceAll(*stringTS, "{{UpdatesForFieldPointerToAssStructTS}}", UpdatesForFieldPointerToAssStructTS)

	*stringHTML = strings.ReplaceAll(*stringHTML, "{{AssociationStructFormsHTML}}", AssociationStructFormsHTML)

	*stringTS = strings.ReplaceAll(*stringTS, "{{ReversePointerToStructGetTS}}", ReversePointerToStructGetTS)

}
