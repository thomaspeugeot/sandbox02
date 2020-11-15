package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	target_engine "github.com/thomaspeugeot/metabaron/examples/laundromat/go/engine"

	target_controllers "github.com/thomaspeugeot/metabaron/examples/laundromat/go/controllers"
	target_orm "github.com/thomaspeugeot/metabaron/examples/laundromat/go/orm"

	gorgo_controllers "github.com/thomaspeugeot/metabaron/libs/gorgo/go/controllers"
	gorgo_orm "github.com/thomaspeugeot/metabaron/libs/gorgo/go/orm"

	animah_controllers "github.com/thomaspeugeot/metabaron/libs/animah/go/controllers"
	animah_models "github.com/thomaspeugeot/metabaron/libs/animah/go/models"
	animah_orm "github.com/thomaspeugeot/metabaron/libs/animah/go/orm"
)

var (
	logDBFlag         = flag.Bool("logDB", false, "log mode for db")
	logGINFlag        = flag.Bool("logGIN", false, "log mode for gin")
	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

var db *gorm.DB

//
// generic code
//
// specific code is in target_engine
func main() {

	log.SetPrefix("laundromat: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// Animah
	if *clientControlFlag {
		animah_models.EngineSingloton.ControlMode = animah_models.CLIENT_CONTROL
	} else {
		animah_models.EngineSingloton.ControlMode = animah_models.AUTONOMOUS
	}

	// setup GORM
	db = target_orm.SetupModels(*logDBFlag, ":memory:")
	db.DB().SetMaxOpenConns(1)

	// add gorgo database
	gorgo_orm.AutoMigrate(db)

	// add animah database
	animah_orm.AutoMigrate(db)

	// attach specific engine callback to the laundromat model
	engineSpecificSingloton := target_engine.NewEngineSpecific(db)
	animah_models.AllModelStore.Engines = append(animah_models.AllModelStore.Engines, &animah_models.EngineSingloton)
	animah_models.EngineSingloton.EngineSpecificInteface = engineSpecificSingloton

	//
	//  setup controlers
	//
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	target_controllers.RegisterControllers(r)
	animah_controllers.RegisterControllers(r)
	gorgo_controllers.RegisterControllers(r)

	// put all to database
	if err := animah_orm.AllModelsToORM(db); err != nil {
		log.Fatal("Unable to put models to DB " + err.Error())
	}
	if err := target_orm.AllModelsToORM(db); err != nil {
		log.Fatal("Unable to put models to DB " + err.Error())
	}
	if err := gorgo_orm.AllModelsToORM(db); err != nil {
		log.Fatal("Unable to put models to DB " + err.Error())
	}

	r.Run()
	os.Exit(0)

}
