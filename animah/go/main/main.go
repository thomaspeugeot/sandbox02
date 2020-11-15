package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/thomaspeugeot/sandbox02/animah/go/controllers"
	"github.com/thomaspeugeot/sandbox02/animah/go/models"
	"github.com/thomaspeugeot/sandbox02/animah/go/orm"
)

var (
	logDBFlag         = flag.Bool("logDB", false, "log mode for db")
	logGINFlag        = flag.Bool("logGIN", false, "log mode for gin")
	clientControlFlag = flag.Bool("client-control", false, "if true, engine waits for API calls")
)

func main() {

	log.SetPrefix("animah: ")
	log.SetFlags(0)

	flag.Parse()
	if len(flag.Args()) > 0 {
		log.Fatal("surplus arguments")
	}

	// setup GORM
	db := orm.SetupModels(*logDBFlag, ":memory:")

	if *clientControlFlag {
		models.EngineSingloton.ControlMode = models.CLIENT_CONTROL
	} else {
		models.EngineSingloton.ControlMode = models.AUTONOMOUS
	}

	// seven days of simulation
	models.EngineSingloton.StartTime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	models.EngineSingloton.CurrentTime = models.EngineSingloton.StartTime
	models.EngineSingloton.State = models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", models.EngineSingloton.StartTime)

	models.EngineSingloton.EndTime = time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC)
	log.Printf("Sim end  \t\t\t%s\n", models.EngineSingloton.EndTime)

	// register engine
	models.AllModelStore.Engines = append(models.AllModelStore.Engines, &models.EngineSingloton)
	orm.AllModelsToORM(db)

	// setup controlers
	r := gin.Default()
	r.Use(cors.Default())

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // a gin Context can have a map of variable that is set up at runtime
		c.Next()
	})

	controllers.RegisterControllers(r)

	if *clientControlFlag {
		r.Run()
	} else {
		models.EngineSingloton.Run()
	}

	// r.StaticFS("/static/", http.Dir("/Users/thomaspeugeot/go/src/github.com/thomaspeugeot/sandbox02/animah/ng/dist/ng"))

}
