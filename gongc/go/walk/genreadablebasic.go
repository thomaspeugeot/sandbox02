package walk

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const ngReadableBasicFieldHTML = `

		<div class="mat-row">
			<div class="mat-cell left">{{FieldName}}</div>
			<div class="mat-cell right">{{{{structname}}.{{FieldName}}}}</div>
		</div>
`

const ngReadablePointerToStructFieldHTML = `

		<div *ngIf="{{FieldName}}" class="mat-row  pointer-to-struct-row" 
		(click)="setPresentationRouterOutlet( '{{assocStructName}}', {{FieldName}}.ID )">
			<div class="mat-cell left">{{FieldName}}</div>
			<div class="mat-cell right">{{{{FieldName}}.Name}}</div>
		</div>
`

func genReadableBasic(db *gorm.DB, fields models.Fields, _struct models.Struct, stringTS *string, stringHTML *string) {

	NgReadableBasicHTML := ""

	for _, field := range fields {
		if field.Kind != reflect.Ptr && field.Kind != reflect.Slice &&
			(field.Kind == reflect.Int || field.Kind == reflect.String) {

			res := strings.ReplaceAll(ngReadableBasicFieldHTML, "{{FieldName}}", field.Name)

			NgReadableBasicHTML += res
		}

		if field.Kind != reflect.Ptr && field.Kind != reflect.Slice &&
			(field.Kind == reflect.Float64) {

			// for float display, IEEE-754 makes it eay to have anoying
			// rounding effects with float like 9.000000001 or 7.9999999954
			// therefore, we have the display to systematicaly round the result.
			// TODO #109 magic gong code for setting up the precision
			res := strings.ReplaceAll(ngReadableBasicFieldHTML, ".{{FieldName}}", "."+field.Name+".toPrecision(5)")
			res = strings.ReplaceAll(res, "{{FieldName}}", field.Name)

			NgReadableBasicHTML += res
		}

		// if field.Kind == reflect.Struct && field.AssociatedStructName == "time.Time" {

		// 	res := strings.ReplaceAll(ngReadableBasicFieldHTML, "{{FieldName}}", field.Name)

		// 	NgReadableBasicHTML += res
		// }

	}

	*stringHTML = strings.ReplaceAll(*stringHTML, "{{NgReadableBasicHTML}}", NgReadableBasicHTML)

}
