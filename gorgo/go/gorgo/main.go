package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/controllers"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	"github.com/thomaspeugeot/sandbox02/gorgo/go/orm"

	"github.com/thomaspeugeot/sandbox02/gorgo/go/tests/geometry/diagrams"
)

var (
	diagramPkgPath = flag.String("pkgPath", "../../../metabaron/examples/bookstore/models", "path to package for analysis")

	logBBFlag = flag.Bool("logDB", false, "log mode for db")
)

func main() {

	log.SetPrefix("gorgo: ")
	log.SetFlags(0)
	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	log.Printf("Diagram 1 : %s", diagrams.Diagram1.Name)

	// setup GORM
	db := orm.SetupModels(*logBBFlag, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	controllers.RegisterControllers(r)

	// r.StaticFS("/static/", http.Dir("/Users/thomaspeugeot/go/src/github.com/thomaspeugeot/sandbox02/bookstore/ng/dist/ng"))

	var pkgelt models.Pkgelt
	// parse the diagram package
	pkgelt.Unmarshall(*diagramPkgPath)

	// parse the "docs" package
	docsPkgPath := filepath.Join(*diagramPkgPath, "../docs")
	pkgelt.Unmarshall(docsPkgPath)

	// append package to store
	models.AllModelStore.Pkgelts = append(models.AllModelStore.Pkgelts, &pkgelt)

	if err := orm.AllModelsToORM(db); err != nil {
		log.Fatal("Unable to put models to DB " + err.Error())
	}

	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))

	// attach callback to Action controller
	marshaller := Marshaller{}
	marshaller.db = db
	controllers.GorgoactionSinglotonID.Callback = &marshaller

	r.Run(":8080")
}

// Marshaller catch callback
type Marshaller struct {
	db *gorm.DB
}

// PostGorgoaction is called from the Controller
func (marshaller *Marshaller) PostGorgoaction(action *models.Gorgoaction) {

	log.Printf("Post action called with %s", action.Name)

	switch action.Name {
	case string(models.MARSHALL_ALL_DIAGRAMS):
		// put into models store
		models.AllModelStore.Reset()
		orm.AllORMToModels(marshaller.db)

		// marshall
		diagramsPath := filepath.Join(*diagramPkgPath, "../diagrams")

		// fetch the only classdiagram collection
		for _, pkgelt := range models.AllModelStore.Pkgelts {
			pkgelt.Marshall(diagramsPath)
		}
	case string(models.UNMARSHALL_ALL_DIAGRAMS):

		// close & reopen database
		marshaller.db.Unscoped().Delete(&orm.GorgoactionDB{})
		marshaller.db.Unscoped().Delete(&orm.ClassdiagramDB{})
		marshaller.db.Unscoped().Delete(&orm.ClassshapeDB{})
		marshaller.db.Unscoped().Delete(&orm.LinkDB{})
		marshaller.db.Unscoped().Delete(&orm.PkgeltDB{})
		marshaller.db.Unscoped().Delete(&orm.PositionDB{})
		marshaller.db.Unscoped().Delete(&orm.StateDB{})
		marshaller.db.Unscoped().Delete(&orm.UmlscDB{})
		marshaller.db.Unscoped().Delete(&orm.VerticeDB{})
		marshaller.db.Unscoped().Delete(&orm.FieldDB{})

		// reset the tree maps and the AllModelStore
		orm.AllORMToModels(marshaller.db)

		models.AllModelStore.Reset()
		var pkgelt models.Pkgelt
		pkgelt.Unmarshall(*diagramPkgPath)

		if err := orm.AllModelsToORM(marshaller.db); err != nil {
			log.Fatal("Unable to put models to DB " + err.Error())
		}
	case string(models.PRINT_ALL_DOCUMENTS):
		// fetch the only classdiagram collection

	default:
		log.Panic("unkwn action " + action.Name)
	}
}
