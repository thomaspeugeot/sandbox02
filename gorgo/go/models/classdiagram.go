package models

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"
	"os"
	"sort"
)

// Classdiagram diagram struct store a class diagram
// temporary here
// swagger:model Classdiagram
type Classdiagram struct {
	Name string

	// this is the memory model (and not the "memory motel" of the Rolling Stones)
	// it is not ignored by swagger because it is used by the angular model
	Classshapes []*Classshape
}

const DiagramMarginX = 10.0
const DiagramMarginY = 10.0

// Extent compute max X and max Y
func (classdiagram *Classdiagram) Extent() (x, y, maxClassshapeHeigth float64) {

	for _, classshape := range classdiagram.Classshapes {

		if maxClassshapeHeigth < classshape.Heigth {
			maxClassshapeHeigth = classshape.Heigth
		}

		if classshape.Position.X+classshape.Width > x {
			x = classshape.Position.X + classshape.Width
		}
		if classshape.Position.Y+classshape.Heigth > y {
			y = classshape.Position.Y + classshape.Heigth
		}

		for _, link := range classshape.Links {
			if link.Middlevertice.X > x {
				x = link.Middlevertice.X
			}
			if link.Middlevertice.Y > y {
				y = link.Middlevertice.Y
			}
		}
	}
	// margin
	x += DiagramMarginX
	y += DiagramMarginY
	return
}

// MarshallAsVariable transform the class diagram into a var
func (classdiagram *Classdiagram) MarshallAsVariable(file *os.File) error {

	fmt.Fprintf(file, "var %s uml.Classdiagram = uml.Classdiagram{\n", classdiagram.Name)

	fmt.Fprintf(file, "\tClassshapes: []*uml.Classshape{\n")

	if len(classdiagram.Classshapes) > 0 {
		// sort Classshapes
		sort.Slice(classdiagram.Classshapes[:], func(i, j int) bool {
			return classdiagram.Classshapes[i].Structname < classdiagram.Classshapes[j].Structname
		})
		for _, classshape := range classdiagram.Classshapes {
			classshape.Marshall(file, 2)
			fmt.Fprintf(file, ",\n")
		}
		fmt.Fprintf(file, "\t},\n")
	}

	fmt.Fprintf(file, "}\n\n")
	return nil
}

// ClassdiagramMap is a Map of all Classdiagram via their Name
type ClassdiagramMap map[string]*Classdiagram

// ClassdiagramStore is a handy ClassdiagramMap
var ClassdiagramStore ClassdiagramMap = make(map[string]*Classdiagram, 0)

// Unmarshall updates a classdiagram values from an ast.Epr
// and appends it to the ClassdiagramStore
func (classdiagram *Classdiagram) Unmarshall(expr ast.Expr, fset *token.FileSet) {

	// expression should be a composite literal expression
	// "uml.Classdiagram{
	//   	Classshapes: []*uml.Classshape{ ...
	//		Links: []*uml.Links{ ...
	if cl, ok := expr.(*ast.CompositeLit); ok {

		// parse all KeyValues of the Classdiagram
		for _, elt := range cl.Elts {
			if structvaluekeyexpr, ok := elt.(*ast.KeyValueExpr); !ok {
				log.Panic("Expression should be a struct key" +
					fset.Position(structvaluekeyexpr.Pos()).String())
			} else {

				if key, ok := structvaluekeyexpr.Key.(*ast.Ident); !ok {
					log.Panic("Key shoud be an ident" +
						fset.Position(key.Pos()).String())
				} else {
					switch key.Name {
					case "Classshapes":
						var cl *ast.CompositeLit
						var ok bool
						if cl, ok = structvaluekeyexpr.Value.(*ast.CompositeLit); !ok {
							log.Panic("Value shoud be a composite lit" +
								fset.Position(structvaluekeyexpr.Pos()).String())
						}
						// get the array of shapes (either as definition or as reference)
						for _, expr := range cl.Elts {

							var classshape *Classshape
							switch exp := expr.(type) {
							case *ast.UnaryExpr: // this is a reference to a variable
								if ident, ok := exp.X.(*ast.Ident); !ok {
									log.Panic("" + fset.Position(exp.Pos()).String())
								} else {
									log.Printf("found %s", ident.Name)

									for _, _classshape := range AllModelStore.Classshapes {
										if _classshape.Name == ident.Name {
											classshape = _classshape
											continue
										}
									}
									if classshape == nil {
										log.Panic("Unable to find the shape " + ident.Name + " " +
											fset.Position(cl.Pos()).String())
									}
								}
							case *ast.CompositeLit: // this is a definition
								var _classshape Classshape
								classshape = &_classshape
								classshape.Unmarshall(exp, fset)
							default:
								log.Panic("Value shoud be a composite lit or a unary" +
									fset.Position(structvaluekeyexpr.Pos()).String())
							}

							classdiagram.Classshapes = append(classdiagram.Classshapes, classshape)
						}

					case "Name":
						// already initialized
					default:
						log.Panic("Key shoud be Classshapes, Field or Link" +
							fset.Position(key.Pos()).String())
					}
				}

			}

		}
	}
	AllModelStore.Classdiagrams = append(AllModelStore.Classdiagrams, classdiagram)
}

// serialize the package and its elements to the AllModelStore
// this is used if one Umlsc is dynamicaly created
func (classdiagram *Classdiagram) SerializeToAllModelStore() {

	AllModelStore.Classdiagrams = append(AllModelStore.Classdiagrams, classdiagram)

	for _, classshape := range classdiagram.Classshapes {
		classshape.SerializeToAllModelStore()
	}
}
