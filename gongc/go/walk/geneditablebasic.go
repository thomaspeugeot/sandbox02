package walk

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const ngTemplateDetailFormFieldHTML = `
  <form>
    <mat-form-field>
      <mat-label>{{FieldName}}</mat-label>
      <input {{TypeInput}} matInput [(ngModel)]="{{structname}}.{{FieldName}}">
    </mat-form-field>
  </form>
`

func genEditableBasic(db *gorm.DB, fields models.Fields, _struct models.Struct, stringTS *string, stringHTML *string) {

	NgEditableBasicInputFormHTML := ""

	for _, field := range fields {
		if field.Kind != reflect.Ptr && field.Kind != reflect.Slice &&
			(field.Kind == reflect.Int || field.Kind == reflect.Float64 || field.Kind == reflect.String) {

			// conversion form go type to ts type
			TypeInput := "name=\"\" [ngModelOptions]=\"{standalone: true}\" 	"
			if field.Kind == reflect.Int || field.Kind == reflect.Float64 {
				TypeInput = "type=\"number\" [ngModelOptions]=\"{standalone: true}\" "
			}

			// for html
			// prepare for type="number"
			res := strings.ReplaceAll(ngTemplateDetailFormFieldHTML, "{{FieldName}}", field.Name)
			res = strings.ReplaceAll(res, "{{TypeInput}}", TypeInput)
			NgEditableBasicInputFormHTML += res
		}
	}

	*stringHTML = strings.ReplaceAll(*stringHTML, "{{NgEditableBasicInputFormHTML}}", NgEditableBasicInputFormHTML)

}
