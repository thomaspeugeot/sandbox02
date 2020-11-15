package walk

import (
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

// NgDetailTemplateBooleanFieldDeclarationTS ...
const NgDetailTemplateBooleanFieldDeclarationTS = `
	{{FieldName}}FormControl = new FormControl(false);`

// NgDetailTemplateBooleanFieldInitialisationTS ...
const NgDetailTemplateBooleanFieldInitialisationTS = `
					this.{{FieldName}}FormControl.setValue( this.{{structname}}.{{FieldName}})`

// NgDetailTemplateBooleanFieldRecoveryTS ...
const NgDetailTemplateBooleanFieldRecoveryTS = `
	this.{{structname}}.{{FieldName}} = this.{{FieldName}}FormControl.value`

const NgTemplateEditablePointerToBoolFormFieldHTML = `
	<form>
		<mat-checkbox [formControl]="{{FieldName}}FormControl">{{FieldName}}</mat-checkbox>
	</form>
`

func genEditablePointerToBool(db *gorm.DB, fields models.Fields, _struct models.Struct, stringTS *string, stringHTML *string) {
	BooleanFieldFormControlDeclarationsTS := ""
	BooleanFieldFormControlInitialisationsTS := ""
	BooleanFieldFormRecoveriesTS := ""
	BooleanFieldFormsHTML := ""
	for _, field := range fields {

		if field.Kind == reflect.Ptr && field.AssociatedStructName == "*bool" {

			res := strings.ReplaceAll(NgDetailTemplateBooleanFieldDeclarationTS, "{{FieldName}}", field.Name)
			BooleanFieldFormControlDeclarationsTS += res

			res = strings.ReplaceAll(NgDetailTemplateBooleanFieldInitialisationTS, "{{FieldName}}", field.Name)
			BooleanFieldFormControlInitialisationsTS += res

			res = strings.ReplaceAll(NgDetailTemplateBooleanFieldRecoveryTS, "{{FieldName}}", field.Name)
			BooleanFieldFormRecoveriesTS += res

			res = strings.ReplaceAll(NgTemplateEditablePointerToBoolFormFieldHTML, "{{FieldName}}", field.Name)
			BooleanFieldFormsHTML += res
		}
	}
	*stringTS = strings.ReplaceAll(*stringTS, "{{BooleanFieldFormControlDeclarations}}", BooleanFieldFormControlDeclarationsTS)
	*stringTS = strings.ReplaceAll(*stringTS, "{{BooleanFieldFormControlInitialisations}}", BooleanFieldFormControlInitialisationsTS)
	*stringTS = strings.ReplaceAll(*stringTS, "{{BooleanFieldFormRecoveries}}", BooleanFieldFormRecoveriesTS)

	*stringHTML = strings.ReplaceAll(*stringHTML, "{{BooleanFieldFormsHTML}}", BooleanFieldFormsHTML)
}
