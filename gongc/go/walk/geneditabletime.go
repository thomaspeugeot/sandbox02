package walk

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const ngTemplateTimeHTML = `
  <form>
    <mat-form-field>
      <mat-label>{{FieldName}}</mat-label>
      <input {{TypeInput}} matInput [(ngModel)]="{{structname}}.{{FieldName}}">
    </mat-form-field>
  </form>
`

func genEditableTime(db *gorm.DB, fields models.Fields, _struct models.Struct, stringTS *string, stringHTML *string) {

	NgEditableTimeInputFormHTML := ""

	for _, field := range fields {
		if field.Kind == reflect.Struct && field.AssociatedStructName == "time.Time" {

			TypeInput := "name=\"\" [ngModelOptions]=\"{standalone: true}\" 	"

			// for html
			// prepare for type="number"
			res := strings.ReplaceAll(ngTemplateTimeHTML, "{{FieldName}}", field.Name)
			res = strings.ReplaceAll(res, "{{TypeInput}}", TypeInput)
			NgEditableTimeInputFormHTML += res
		}
	}

	*stringHTML = strings.ReplaceAll(*stringHTML, "{{NgEditableTimeInputFormHTML}}", NgEditableTimeInputFormHTML)

}
