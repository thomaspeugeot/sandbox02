// generated by genGOAllModelStruct.go
package models



// CreateORMMachine enables dynamic registration of a Machine instance
func CreateORMMachine(Machine *Machine) {
	AllModelStore.Machines = append(AllModelStore.Machines, Machine)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMMachine(Machine)
	}
}


// DeleteORMMachine enables dynamic registration of a Machine instance
func DeleteORMMachine(Machine *Machine) {
	for index, _Machine := range AllModelStore.Machines {
		if _Machine == Machine {
			AllModelStore.Machines[index] = AllModelStore.Machines[len(AllModelStore.Machines)-1]
			AllModelStore.Machines = AllModelStore.Machines[:len(AllModelStore.Machines)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMMachine(Machine)
	}
}

// CreateORMWasher enables dynamic registration of a Washer instance
func CreateORMWasher(Washer *Washer) {
	AllModelStore.Washers = append(AllModelStore.Washers, Washer)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMWasher(Washer)
	}
}


// DeleteORMWasher enables dynamic registration of a Washer instance
func DeleteORMWasher(Washer *Washer) {
	for index, _Washer := range AllModelStore.Washers {
		if _Washer == Washer {
			AllModelStore.Washers[index] = AllModelStore.Washers[len(AllModelStore.Washers)-1]
			AllModelStore.Washers = AllModelStore.Washers[:len(AllModelStore.Washers)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMWasher(Washer)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface {
	CreateORMMachine(Machine *Machine)
	CreateORMWasher(Washer *Washer)
}
	
type AllModelsStructDeleteInterface interface {
	DeleteORMMachine(Machine *Machine)
	DeleteORMWasher(Washer *Washer)
}
		
// swagger:ignore
type AllModelStoreStruct struct {
	Machines  []*Machine
	Washers  []*Washer
	AllModelsStructCreateCallback AllModelsStructCreateInterface
	AllModelsStructDeleteCallback AllModelsStructDeleteInterface
}

// swagger:ignore
var AllModelStore AllModelStoreStruct = AllModelStoreStruct{
	Machines:	make([]*Machine, 0),
	Washers:	make([]*Washer, 0),
}

func (allModelStoreStruct * AllModelStoreStruct) Reset() {
	allModelStoreStruct.Machines =	make([]*Machine, 0)
	allModelStoreStruct.Washers =	make([]*Washer, 0)
}

func (allModelStoreStruct * AllModelStoreStruct) Nil() {
	allModelStoreStruct.Machines = nil
	allModelStoreStruct.Washers = nil
}