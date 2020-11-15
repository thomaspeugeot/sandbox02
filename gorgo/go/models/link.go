package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"reflect"
	"strings"
)

// Link represent the UML Link in any diagram
// uni-directional https://en.wikipedia.org/wiki/Association_(object-oriented_programming)
// More specificaly It is a 0..1 ---> 0..1
// swagger:model Link
type Link struct {
	Name string

	// swagger:ignore
	Field         interface{} `gorm:"-"` // field that is diagrammed
	Fieldname     string
	Structname    string
	Fieldtypename string

	// Vertices at the middle
	// swagger:ignore
	Middlevertice *Vertice
}

// Unmarshall
func (link *Link) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	var cl *ast.CompositeLit
	var ok bool
	if cl, ok = expr.(*ast.CompositeLit); !ok {
		log.Panic("Expecting a composite litteral like 	Field: models.Line{}.Start, " +
			fset.Position(expr.Pos()).String())
	}

	{
		var kve *ast.KeyValueExpr
		if kve, ok = cl.Elts[0].(*ast.KeyValueExpr); !ok {
			log.Panic("Expecting 1 key value Field: ... " + fset.Position(cl.Pos()).String())
		}

		// get the key "Field"
		var ident *ast.Ident
		if ident, ok = kve.Key.(*ast.Ident); !ok {
			log.Panic("Expecting 1 ident " + fset.Position(kve.Pos()).String())
		}
		// check Link Field is "Field", ah ah ah !!! happy meta programmer
		if ident.Name != "Field" {
			log.Panic("Expecting 1 Field Field " + fset.Position(ident.Pos()).String())
		}

		// parsing models.Line{}.Start, which is
		// ast.SelectorExpr
		// 	X: ast.CompositeLit
		// 		Type: ast.SelectorExpr
		// 			X: ast.Ident
		// 				Name: "models"
		// 			Sel: ast.Ident
		// 				Name: "Line"
		// 	Sel: ast.Ident
		// 		Name: "Start"
		var se *ast.SelectorExpr
		if se, ok = kve.Value.(*ast.SelectorExpr); !ok {
			log.Panic("Expecting 1 selector " + fset.Position(kve.Pos()).String())
		}

		var cl2 *ast.CompositeLit
		if cl2, ok = se.X.(*ast.CompositeLit); !ok {
			log.Panic("Expecting 1 composite lit " + fset.Position(se.Pos()).String())
		}

		var se2 *ast.SelectorExpr
		if se2, ok = cl2.Type.(*ast.SelectorExpr); !ok {
			log.Panic("Expecting 1 selector " + fset.Position(cl.Pos()).String())
		}

		var ident2 *ast.Ident
		if ident2, ok = se2.X.(*ast.Ident); !ok {
			log.Panic("Expecting 1 ident " + fset.Position(se2.Pos()).String())
		}

		structnameWithX := ident2.Name + "." + se2.Sel.Name
		link.Structname = se2.Sel.Name
		link.Fieldname = se.Sel.Name

		// now, let's find the link target !!!
		fieldname := fmt.Sprintf("%s{}.%s", structnameWithX, link.Fieldname)

		fieldtypename := MapExpToType[fieldname]

		// extract only the selector
		words := strings.Split(fieldtypename, ".")
		link.Fieldtypename = words[len(words)-1]
	}
	// extract vertice
	if len(cl.Elts) > 1 {
		var kve *ast.KeyValueExpr
		if kve, ok = cl.Elts[1].(*ast.KeyValueExpr); !ok {
			log.Panic("Expecting 1 key value Field: ... " + fset.Position(cl.Pos()).String())
		}

		// get the key "Field"
		var ident *ast.Ident
		if ident, ok = kve.Key.(*ast.Ident); !ok {
			log.Panic("Expecting 1 ident " + fset.Position(kve.Pos()).String())
		}
		// check Link Field is
		if ident.Name != "Middlevertice" {
			log.Panic("Expecting 1 Field Middlevertice " + fset.Position(ident.Pos()).String())
		}

		var vertice Vertice
		link.Middlevertice = &vertice
		link.Middlevertice.Unmarshall(kve.Value, fset)
	}

	AllModelStore.Links = append(AllModelStore.Links, link)
}

// Marshall provides the element of link as declaration
func (link *Link) Marshall(file *os.File, nbIndentation int) error {
	indent(file, nbIndentation)
	fmt.Fprintf(file, "{\n")

	indent(file, nbIndentation)
	fmt.Fprintf(file, "\tField: models.%s{}.%s,\n", link.Structname, link.Fieldname)

	if link.Middlevertice != nil {
		indent(file, nbIndentation+1)
		fmt.Fprintf(file, "Middlevertice: ")
		link.Middlevertice.Marshall(file, nbIndentation+1)
	}

	indent(file, nbIndentation)
	fmt.Fprintf(file, "}")

	return nil
}

// serialize the package and its elements to the AllModelStore
// this is used if one Umlsc is dynamicaly created
func (link *Link) SerializeToAllModelStore() {

	AllModelStore.Links = append(AllModelStore.Links, link)

	// update name if not done
	if link.Name == "" {
		if link.Field != nil {
			typeofstruct := reflect.TypeOf(link.Field).Elem().String()
			link.Name = strings.Split(typeofstruct, ".")[1]
			link.Structname = link.Name
		}
	}
}
