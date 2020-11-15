package walk

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
)

const goTemplateJSONMarshallIndent = `

func (m {{Structname}}Map) MarshalJSON() (res []byte, err error) {

	// sort {{structname}}s
	{{structname}}sSorted := make([]string, 0, len(m))
	for k, _ := range m {
		{{structname}}sSorted = append({{structname}}sSorted, k)
	}
	sort.Strings({{structname}}sSorted)
	res = append(res, '[')
	idx := 0
	for _, value := range {{structname}}sSorted {
		var b []byte
		b, err = json.Marshal(m[value])
		if err != nil {
			return
		}
		res = append(res, b...)

		if idx < (len(m) - 1) {
			res = append(res, ',')
		}
		idx++
	}
	res = append(res, ']')
	return
}

`

// GenJSONMarshallers generates the setup file for the gorm
// return the registration calls
func GenJSONMarshallers(db *gorm.DB) string {

	jsonmarshallersCodeReg := ""

	// create the list of structs
	var structs []models.Struct
	db.Find(&structs)

	filename := filepath.Join("/tmp", fmt.Sprintf("marshallingjson.go"))

	// we should use go generate
	log.Println("generating jsonmarshaller file : " + filename)

	f, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(f, "package %s", PkgName)

	defer f.Close()

	for _, _struct := range structs {

		lowerCaseStructName := strings.ToLower(_struct.Name)

		res := strings.ReplaceAll(goTemplateJSONMarshallIndent, "{{Structname}}", _struct.Name)
		res = strings.ReplaceAll(res, "{{structname}}", lowerCaseStructName)

		fmt.Fprintf(f, "%s", res)

	}
	return jsonmarshallersCodeReg
}
