package engine

import (
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/thomaspeugeot/sandbox02/laundromat/go/events"
	target_models "github.com/thomaspeugeot/sandbox02/laundromat/go/models"
	target_orm "github.com/thomaspeugeot/sandbox02/laundromat/go/orm"

	gorgo_models "github.com/thomaspeugeot/sandbox02/gorgo/go/models"
	gorgo_orm "github.com/thomaspeugeot/sandbox02/gorgo/go/orm"

	animah_controllers "github.com/thomaspeugeot/sandbox02/animah/go/controllers"
	animah_models "github.com/thomaspeugeot/sandbox02/animah/go/models"
	animah_orm "github.com/thomaspeugeot/sandbox02/animah/go/orm"
)

// EngineSpecific is the callback support for
// events that happens on the generic engine
type EngineSpecific struct {
	db *gorm.DB

	machine target_models.Machine
	washer  target_models.Washer

	machineState target_models.MachineStateEnum
	washerState  target_models.WasherStateEnum

	Etats_Machine *gorgo_models.Umlsc
	Etats_Washer  *gorgo_models.Umlsc
}

func (engineSpecific *EngineSpecific) setDB(db *gorm.DB) {
	engineSpecific.db = db
}

// EventFired is called from the Action
func (engineSpecific *EngineSpecific) EventFired(engine *animah_models.Engine) {

	// update DB with new values
	// convert time.Duration in minutes
	engineSpecific.machine.RemainingTimeMinutes = int(engineSpecific.machine.RemainingTime.Minutes())
	if err := target_orm.AllModelsUpdateToORM(engineSpecific.db); err != nil {
		log.Fatal("Unable to update database with specfic models objects " + err.Error())
	}

	// update DB with new values
	if err := animah_orm.AllModelsUpdateToORM(engineSpecific.db); err != nil {
		log.Fatal("Unable to update database with generic simulation models objects " + err.Error())
	}
}

// NewEngineSpecific ...
func NewEngineSpecific(db *gorm.DB) (engineSpecific *EngineSpecific) {

	engineSpecific = &EngineSpecific{}
	engineSpecific.setDB(db)

	// seven days of simulation for animah
	animah_models.EngineSingloton.StartTime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	animah_models.EngineSingloton.CurrentTime = animah_models.EngineSingloton.StartTime
	animah_models.EngineSingloton.State = animah_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", animah_models.EngineSingloton.StartTime)

	animah_models.EngineSingloton.EndTime = time.Date(2020, time.January, 10, 0, 0, 0, 0, time.UTC)
	log.Printf("Sim end  \t\t\t%s\n", animah_models.EngineSingloton.EndTime)

	// seven days of simulation
	animah_models.EngineSingloton.StartTime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	animah_models.EngineSingloton.CurrentTime = animah_models.EngineSingloton.StartTime
	animah_models.EngineSingloton.State = animah_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", animah_models.EngineSingloton.StartTime)

	animah_models.EngineSingloton.EndTime = time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC)
	log.Printf("Sim end  \t\t\t%s\n", animah_models.EngineSingloton.EndTime)

	animah_models.EngineSingloton.Speed = 1.0

	// append Machine & Washer
	engineSpecific.machine.State = target_models.MACHINE_DOOR_CLOSED_IDLE
	engineSpecific.machine.TechName = "machine-1"
	engineSpecific.machine.Name = "machine-1"
	cleaned := true
	engineSpecific.machine.Cleanedlaundry = &cleaned
	animah_models.EngineSingloton.AppendAgent(&engineSpecific.machine)

	engineSpecific.washer.State = target_models.WASHER_IDLE
	engineSpecific.washer.Machine = &engineSpecific.machine
	engineSpecific.washer.TechName = "washer-1"
	engineSpecific.washer.Name = "washer-1"
	animah_models.EngineSingloton.AppendAgent(&engineSpecific.washer)

	// attach generic engine callback to Action controller
	// to funnel client action to the engine
	actionSingloton := animah_models.ActionSingloton{}
	animah_controllers.ActionSinglotonID.Callback = &actionSingloton

	engineSpecific.CreateInitialEvents()

	engineSpecific.machineState = engineSpecific.machine.State
	engineSpecific.washerState = engineSpecific.washer.State

	target_models.AllModelStore.Machines = append(target_models.AllModelStore.Machines, &engineSpecific.machine)
	target_models.AllModelStore.Washers = append(target_models.AllModelStore.Washers, &engineSpecific.washer)

	var pkgelt gorgo_models.Pkgelt
	gorgo_models.AllModelStore.Pkgelts = append(gorgo_models.AllModelStore.Pkgelts, &pkgelt)

	// classdiagram can only be fully in memory when they are Unmarshalled
	// for instance, the Name of diagrams or the Name of the Link
	pkgelt.Unmarshall("../diagrams")

	// Umarshall docmuments (after diagrams, because they might reference diagrams)
	pkgelt.Unmarshall("../docs")

	for _, umlsc := range gorgo_models.AllModelStore.Umlscs {
		if strings.Contains(umlsc.Name, "Machine") {
			engineSpecific.Etats_Machine = umlsc
		}
		if strings.Contains(umlsc.Name, "Washer") {
			engineSpecific.Etats_Washer = umlsc
		}
	}

	return
}

// CreateInitialEvents ...
func (engineSpecific *EngineSpecific) CreateInitialEvents() {

	// washer checks its state every 30''
	var washerUpdateStateEvent animah_models.UpdateState
	washerUpdateStateEvent.FireTime = animah_models.EngineSingloton.StartTime
	washerUpdateStateEvent.Period = 30 * time.Second
	washerUpdateStateEvent.Name = "w update"

	engineSpecific.washer.QueueEvent(&washerUpdateStateEvent)

	// washer add laundry periodic event (every 12 hours)
	var addLaundry events.NewDirtyLaundry
	addLaundry.FireTime = animah_models.EngineSingloton.StartTime
	addLaundry.Period = events.LAUNDRY_LOAD_PERIOD
	addLaundry.Name = "addLaundry"

	engineSpecific.washer.QueueEvent(&addLaundry)

	// machine checks state every minute
	var machineUpdateStateEvent animah_models.UpdateState
	machineUpdateStateEvent.FireTime = animah_models.EngineSingloton.StartTime
	machineUpdateStateEvent.Period = time.Minute
	machineUpdateStateEvent.Name = "m update"

	engineSpecific.machine.QueueEvent(&machineUpdateStateEvent)
}

// HasAnyStateChanged ...
func (engineSpecific *EngineSpecific) HasAnyStateChanged(engine *animah_models.Engine) bool {

	if engineSpecific.washer.State != engineSpecific.washerState ||
		engineSpecific.machine.State != engineSpecific.machineState {
		engineSpecific.washerState = engineSpecific.washer.State
		engineSpecific.machineState = engineSpecific.machine.State

		engineSpecific.Etats_Machine.Activestate = string(engineSpecific.machineState)
		engineSpecific.Etats_Washer.Activestate = string(engineSpecific.washerState)

		if err := gorgo_orm.AllModelsUpdateToORM(engineSpecific.db); err != nil {
			log.Fatal("Unable to update database with generic simulation models objects " + err.Error())
		}

		log.Printf("EngineSpecific: HasAnyStateChanged %s", engine.CurrentTime.String())
		return true
	}

	return false
}
