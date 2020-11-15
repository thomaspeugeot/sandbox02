package walk

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
	"log"
	"path/filepath"
	"reflect"

	"github.com/thomaspeugeot/metabaron/go/models"
	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

// Walk parse structs
func Walk(RelativePkgPath string) {

	directory, err := filepath.Abs(RelativePkgPath)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

	conf := types.Config{Importer: importer.Default()}
	fset := token.NewFileSet()
	var files []*ast.File
	_, err1 := conf.Check(RelativePkgPath, fset, files, nil)
	if err1 != nil {
		log.Fatal(err)
	}

	cfg := &packages.Config{
		Dir:   directory,
		Mode:  pkgLoadMode,
		Tests: false,
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		s := fmt.Sprintf("cannot process package at path %s, err %s", RelativePkgPath, err.Error())
		log.Panic(s)
	}

	// store struct according to its name
	structMap := make(map[string]*models.Struct)

	// Traverse is in 2 steps
	// first traverse to gather all metabaron types from all Concepts
	// second taverse again to gather metabaron types fields and link them to spinsa types

	if len(pkgs) != 1 {
		log.Panicf("Expected 1 package to scope, found %d", len(pkgs))
	}
	pkg := pkgs[0]

	// fetch all pkg names ...
	scope := pkg.Types.Scope()
	log.Printf("Scoping package %s with path %s", pkg.Name, pkg.PkgPath)

	for _, name := range scope.Names() {

		obj := scope.Lookup(name)
		log.Println("obj name " + obj.Name())

		// check can be of type *type.Struct
		underlying := obj.Type().Underlying()

		// we are only interested in struct
		switch underlying.(type) {

		case *types.Struct:
			longName := obj.Type().String()
			log.Println("name : " + longName)

			// instanciate a metabaron type
			var structMB models.Struct
			structMB.Name = name

			// add it to the map (for later, when linking)
			structMap[longName] = &structMB
		default:
		}
	}

	// fetch all pkg names ...

	scope = pkg.Types.Scope()

	for _, name := range scope.Names() {

		// fetch object
		obj := scope.Lookup(name)

		switch obj.(type) {
		// case *types.Func:
		// 	log.Printf("Func")
		// case *types.Var:
		// 	log.Printf("var")
		// case *types.Const:
		// 	log.Printf("const")
		case *types.TypeName:
			log.Printf("typename")
			enumValues, _ := FindEnumValues(pkg, obj.Name())
			if len(enumValues) > 0 {

				cmt, hasComments := FindComments(pkg, obj.Name())
				if !hasComments {
					// log.Printf("no comment")
				} else {
					log.Printf("cmt \n\n%s\n", cmt.Text())
				}
				log.Print(enumValues)
			}
		case *types.Label:
			log.Printf("label")
		case *types.PkgName:
			log.Printf("PkgName      ")
		case *types.Builtin:
			log.Printf("Builtin            ")
		case *types.Nil:
			log.Printf("Nil      ")
		default:
		}
	}
	for _, name := range scope.Names() {

		// fetch object
		obj := scope.Lookup(name)

		underlying := obj.Type().Underlying()

		switch underlyingType := underlying.(type) {

		case *types.Named:
			log.Printf("named %s\n", underlyingType.String())

		case *types.Signature:
			log.Printf("signature %s\n", underlyingType.String())

		case *types.Tuple:
			log.Printf("tuple %s\n", underlyingType.String())

		case *types.Struct:
			name := obj.Type().String()
			_struct := structMap[name]
			if _struct == nil {
				log.Panic("unkown metabaron type " + name)
			}

			for fieldIndex := 0; fieldIndex < underlyingType.NumFields(); fieldIndex++ {
				log.Printf("field #%d\n", fieldIndex)

				variable := underlyingType.Field(fieldIndex)
				log.Printf("field name %s\n", variable.Name())
				log.Printf("field type name %s\n", variable.Type().String())

				// setup field to append to struct field store
				var field models.FieldModel

				field.Name = variable.Name()
				field.StructID = _struct.ID
				field.StructName = _struct.Name

				// fetch the field
				// we have to process pointers to metabaron type & array of of pointers
				fieldUnderlying := variable.Type().Underlying()
				switch t2 := fieldUnderlying.(type) {
				case *types.Basic:
					log.Printf("field is a basic pointer of type %s\n", t2.Underlying().String())
					switch t2.Kind() {
					case types.Bool:
						field.Kind = reflect.Bool
					case types.Int64:
						field.Kind = reflect.Int
					case types.Int:
						field.Kind = reflect.Int
					case types.Uint:
						field.Kind = reflect.Uint
					case types.String:
						field.Kind = reflect.String
					case types.Float64:
						field.Kind = reflect.Float64
					default:
						log.Panic("not suitable kind for metabaron")
					}
				case *types.Pointer:
					// if pointer, field should be of form *Typemetabaron
					// or a pointer on a basic type
					assocLongName := t2.Elem().String()
					log.Printf("field is a pointer of type %s\n", assocLongName)

					// fetch it in the struct store
					association := structMap[assocLongName]

					if association == nil {
						field.AssociatedStructName = variable.Type().String()
					} else {
						field.AssociatedStructID = association.ID
						field.AssociatedStructName = association.Name
					}
					field.Kind = reflect.Ptr
				case *types.Slice:
					// if slice, field should be of form []*Typemetabaron

					// first, fetch type of array, should be *Typemetabaron
					assocPtr := t2.Elem()

					// then, fetch the type after the pointer, should be Typemetabaron
					switch t3 := assocPtr.(type) {
					case *types.Pointer:
						assocLongName := t3.Elem().String()
						log.Printf("field is a slice of type %s\n", assocLongName)

						// fetch it in the struct store
						association := structMap[assocLongName]

						if association == nil {
							log.Panic("unkown association " + assocLongName)
						}

						field.AssociatedStructID = association.ID
						field.AssociatedStructName = association.Name

						field.Kind = reflect.Slice

					default:
						log.Panic("unkown type in slice")
					}
				case *types.Struct:
					// if Strut, field should be either Typemetabaron of time.Time

					assocLongName := variable.Type().String()
					log.Printf("field is a struct of type %s\n", assocLongName)

					// fetch it in the struct store
					association := structMap[assocLongName]

					if association == nil {
						field.AssociatedStructName = variable.Type().String()
					} else {
						field.AssociatedStructID = association.ID
						field.AssociatedStructName = association.Name
					}
					field.Kind = reflect.Struct
				default:
				}
				// add field to DB
				fieldDB := models.Field{}
				fieldDB.FieldModel = field
			}
		default:
		}
	}

}
