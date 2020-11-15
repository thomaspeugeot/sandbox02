package walk

import (
	"fmt"
	"go/types"
	"log"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gongc/go/models"
	"golang.org/x/tools/go/packages"
)

const pkgLoadMode = packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo

// RelativePkgPath is the path to the package to be analyzed
var RelativePkgPath string

// PkgGoPath for generation
var PkgGoPath string

// PkgName is generated package name (for back) and generated project elements for front
var PkgName string

// BackendTargetPath for instance "/tmp"
var BackendTargetPath string

// OrmPkgGenPath is target path for orm package, for instance "/tmp/libraryorm"
var OrmPkgGenPath string

// ApiPkgGenPath is target path for api package
var ApiPkgGenPath string

// ControllersPkgGenPath is target path for controllers package, for instance "/tmp/librarycontrollers"
var ControllersPkgGenPath string

// // ModulesTargetPath is where ng modules are generated
// var ModulesTargetPath string

// MatTargetPath is where the ng components are generated
var MatTargetPath string

// NgWorkspacePath is the path to the Ng Workspace
var NgWorkspacePath string

// ADDR is the network address addr where the angular generated service will lookup the server
var ADDR string

// Walk parse structs
func Walk(db *gorm.DB) {

	// https://gorm.io/docs/delete.html
	// if the primary key field is blank, GORM will delete all records for the model
	blankStruct := models.Struct{}
	blankField := models.Field{}
	db.Delete(&blankStruct)
	db.Delete(&blankField)

	directory, err := filepath.Abs(RelativePkgPath)
	if err != nil {
		log.Panic("Path does not exist %s ;" + directory)
	}
	log.Println("Loading package " + directory)

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
	// first traverse to gather all spinosa types from all Concepts
	// second taverse again to gather spinosa types fields and link them to spinsa types

	if len(pkgs) != 1 {
		log.Panicf("Expected 1 package to scope, found %d", len(pkgs))
	}
	pkg := pkgs[0]

	// compute root package path name
	PkgGoPath = pkg.PkgPath
	PkgName = filepath.Base(filepath.Join(pkg.PkgPath, "../.."))

	// fetch all pkg names ...
	scope := pkg.Types.Scope()

	for _, name := range scope.Names() {

		obj := scope.Lookup(name)
		log.Printf("obj name is %s", obj.Name())
		if obj.Name() == "ENUM_VAL1" {
			fst := pkg.Fset
			pos := obj.Pos()
			file := fst.File(pos)
			offset := file.Offset(pos)
			log.Printf("%s is in file %s at pos %d ", obj.Name(), file.Name(), offset)
			// os.Exit(0)
		}

		switch obj.(type) {
		case *types.TypeName:
			log.Printf("obj is a Type declation %s", obj.Name())
		case *types.Const:

			// we process the case when the value of the Const is a string
			var ok bool
			var cst *types.Const
			if cst, ok = obj.(*types.Const); !ok {
				log.Panic("...")
			}

			// get the type
			var named *types.Named
			if named, ok = cst.Type().(*types.Named); !ok {
				// it must be some sort of other const
				continue
			}
			log.Printf("obj is a Const declation %s with enum %s", cst.Name(), named.Obj().Name())

			// fetch the enum, if it does not exist, create it
			var _enumTarget models.EnumDB
			_enumTarget.Name = named.Obj().Name()
			var _enum models.EnumDB
			if q := db.Where(&_enumTarget).First(&_enum); q.Error != nil {
				log.Printf("cannot find instance in DB : " + q.Error.Error())
			}

			if _enum.ID == 0 {
				_enum.Name = named.Obj().Name()
				if q := db.Create(&_enum); q.Error != nil {
					log.Panic("cannot find instance in DB : " + q.Error.Error())
				}
			}

			var _const models.Const
			_const.Name = cst.Name()
			_const.EnumID = _enum.ID
			_const.EnumName = _enum.Name
			_const.Value = cst.Val().String()
			if q := db.Create(&_const); q.Error != nil {
				log.Panic("cannot find instance in  DB : " + q.Error.Error())
			}

		default:
			// we are not interested
			continue
		}

		// check can be of type *type.Struct
		underlying := obj.Type().Underlying()

		switch underlying.(type) {

		// we are only interested in struct
		case *types.Struct:
			longName := obj.Type().String()
			log.Println("name : " + longName)

			cmt, hasComments := FindComments(pkg, obj.Name())
			if !hasComments {
				// log.Printf("no comment")
			} else {

				// do not generage something for struct with swwager:ignore
				if strings.Contains(cmt.Text(), "swagger:ignore") {
					log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
					continue
				}
			}

			// instanciate a metabaron type
			var structMB models.Struct
			structMB.Name = name

			// add it to the DB, permanant
			createStructDB := db.Create(&structMB)
			if createStructDB.Error != nil {
				log.Panic("impossible de créer l'instance en BD")
				return
			}

			// add it to the map (for later, when linking)
			structMap[longName] = &structMB

		case *types.Basic:
			// probably a struct
			log.Printf("Detected a typedef with basic underlying %s\n", obj.Type().String())
		default:
		}
	}

	// second traverse
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

	// third traverse
	for _, name := range scope.Names() {

		// fetch object
		obj := scope.Lookup(name)

		switch obj.(type) {
		case *types.TypeName:
			log.Printf("obj is a Type declation, therefore a Struct, hence %s", obj.Name())

		default:
			// we are not interested
			continue
		}

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

			// do not generage something for struct with swwager:ignore
			cmt, hasComments := FindComments(pkg, obj.Name())
			if hasComments && strings.Contains(cmt.Text(), "swagger:ignore") {
				log.Printf("swagger:ignore \n\n%s\n", cmt.Text())
				continue
			}

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
				// we have to process pointers to spinosa type & array of of pointers
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
						log.Panic("not suitable kind for spinosa")
					}
				case *types.Pointer:
					// if pointer, field should be of form *TypeSpinosa
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
					// if slice, field should be of form []*TypeSpinosa

					// first, fetch type of array, should be *TypeSpinosa
					assocPtr := t2.Elem()

					// then, fetch the type after the pointer, should be TypeSpinosa
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
						log.Printf("unkown type in slice" + assocPtr.String())
					}
				case *types.Struct:
					// if Strut, field should be either TypeSpinosa of time.Time

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
				createStructFieldDB := db.Create(&fieldDB)
				if createStructFieldDB.Error != nil {
					log.Panic("impossible de créer l'instance de field en BD")
					return
				}
			}
		default:
		}
	}
}
