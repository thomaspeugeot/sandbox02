// generated by genGOAllModelStruct.go
package models



// CreateORMClassdiagram enables dynamic registration of a Classdiagram instance
func CreateORMClassdiagram(Classdiagram *Classdiagram) {
	AllModelStore.Classdiagrams = append(AllModelStore.Classdiagrams, Classdiagram)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMClassdiagram(Classdiagram)
	}
}


// DeleteORMClassdiagram enables dynamic registration of a Classdiagram instance
func DeleteORMClassdiagram(Classdiagram *Classdiagram) {
	for index, _Classdiagram := range AllModelStore.Classdiagrams {
		if _Classdiagram == Classdiagram {
			AllModelStore.Classdiagrams[index] = AllModelStore.Classdiagrams[len(AllModelStore.Classdiagrams)-1]
			AllModelStore.Classdiagrams = AllModelStore.Classdiagrams[:len(AllModelStore.Classdiagrams)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMClassdiagram(Classdiagram)
	}
}

// CreateORMClassshape enables dynamic registration of a Classshape instance
func CreateORMClassshape(Classshape *Classshape) {
	AllModelStore.Classshapes = append(AllModelStore.Classshapes, Classshape)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMClassshape(Classshape)
	}
}


// DeleteORMClassshape enables dynamic registration of a Classshape instance
func DeleteORMClassshape(Classshape *Classshape) {
	for index, _Classshape := range AllModelStore.Classshapes {
		if _Classshape == Classshape {
			AllModelStore.Classshapes[index] = AllModelStore.Classshapes[len(AllModelStore.Classshapes)-1]
			AllModelStore.Classshapes = AllModelStore.Classshapes[:len(AllModelStore.Classshapes)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMClassshape(Classshape)
	}
}

// CreateORMField enables dynamic registration of a Field instance
func CreateORMField(Field *Field) {
	AllModelStore.Fields = append(AllModelStore.Fields, Field)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMField(Field)
	}
}


// DeleteORMField enables dynamic registration of a Field instance
func DeleteORMField(Field *Field) {
	for index, _Field := range AllModelStore.Fields {
		if _Field == Field {
			AllModelStore.Fields[index] = AllModelStore.Fields[len(AllModelStore.Fields)-1]
			AllModelStore.Fields = AllModelStore.Fields[:len(AllModelStore.Fields)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMField(Field)
	}
}

// CreateORMGorgoaction enables dynamic registration of a Gorgoaction instance
func CreateORMGorgoaction(Gorgoaction *Gorgoaction) {
	AllModelStore.Gorgoactions = append(AllModelStore.Gorgoactions, Gorgoaction)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMGorgoaction(Gorgoaction)
	}
}


// DeleteORMGorgoaction enables dynamic registration of a Gorgoaction instance
func DeleteORMGorgoaction(Gorgoaction *Gorgoaction) {
	for index, _Gorgoaction := range AllModelStore.Gorgoactions {
		if _Gorgoaction == Gorgoaction {
			AllModelStore.Gorgoactions[index] = AllModelStore.Gorgoactions[len(AllModelStore.Gorgoactions)-1]
			AllModelStore.Gorgoactions = AllModelStore.Gorgoactions[:len(AllModelStore.Gorgoactions)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMGorgoaction(Gorgoaction)
	}
}

// CreateORMLink enables dynamic registration of a Link instance
func CreateORMLink(Link *Link) {
	AllModelStore.Links = append(AllModelStore.Links, Link)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMLink(Link)
	}
}


// DeleteORMLink enables dynamic registration of a Link instance
func DeleteORMLink(Link *Link) {
	for index, _Link := range AllModelStore.Links {
		if _Link == Link {
			AllModelStore.Links[index] = AllModelStore.Links[len(AllModelStore.Links)-1]
			AllModelStore.Links = AllModelStore.Links[:len(AllModelStore.Links)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMLink(Link)
	}
}

// CreateORMPkgelt enables dynamic registration of a Pkgelt instance
func CreateORMPkgelt(Pkgelt *Pkgelt) {
	AllModelStore.Pkgelts = append(AllModelStore.Pkgelts, Pkgelt)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMPkgelt(Pkgelt)
	}
}


// DeleteORMPkgelt enables dynamic registration of a Pkgelt instance
func DeleteORMPkgelt(Pkgelt *Pkgelt) {
	for index, _Pkgelt := range AllModelStore.Pkgelts {
		if _Pkgelt == Pkgelt {
			AllModelStore.Pkgelts[index] = AllModelStore.Pkgelts[len(AllModelStore.Pkgelts)-1]
			AllModelStore.Pkgelts = AllModelStore.Pkgelts[:len(AllModelStore.Pkgelts)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMPkgelt(Pkgelt)
	}
}

// CreateORMPosition enables dynamic registration of a Position instance
func CreateORMPosition(Position *Position) {
	AllModelStore.Positions = append(AllModelStore.Positions, Position)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMPosition(Position)
	}
}


// DeleteORMPosition enables dynamic registration of a Position instance
func DeleteORMPosition(Position *Position) {
	for index, _Position := range AllModelStore.Positions {
		if _Position == Position {
			AllModelStore.Positions[index] = AllModelStore.Positions[len(AllModelStore.Positions)-1]
			AllModelStore.Positions = AllModelStore.Positions[:len(AllModelStore.Positions)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMPosition(Position)
	}
}

// CreateORMState enables dynamic registration of a State instance
func CreateORMState(State *State) {
	AllModelStore.States = append(AllModelStore.States, State)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMState(State)
	}
}


// DeleteORMState enables dynamic registration of a State instance
func DeleteORMState(State *State) {
	for index, _State := range AllModelStore.States {
		if _State == State {
			AllModelStore.States[index] = AllModelStore.States[len(AllModelStore.States)-1]
			AllModelStore.States = AllModelStore.States[:len(AllModelStore.States)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMState(State)
	}
}

// CreateORMUmlsc enables dynamic registration of a Umlsc instance
func CreateORMUmlsc(Umlsc *Umlsc) {
	AllModelStore.Umlscs = append(AllModelStore.Umlscs, Umlsc)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMUmlsc(Umlsc)
	}
}


// DeleteORMUmlsc enables dynamic registration of a Umlsc instance
func DeleteORMUmlsc(Umlsc *Umlsc) {
	for index, _Umlsc := range AllModelStore.Umlscs {
		if _Umlsc == Umlsc {
			AllModelStore.Umlscs[index] = AllModelStore.Umlscs[len(AllModelStore.Umlscs)-1]
			AllModelStore.Umlscs = AllModelStore.Umlscs[:len(AllModelStore.Umlscs)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMUmlsc(Umlsc)
	}
}

// CreateORMVertice enables dynamic registration of a Vertice instance
func CreateORMVertice(Vertice *Vertice) {
	AllModelStore.Vertices = append(AllModelStore.Vertices, Vertice)
	if AllModelStore.AllModelsStructCreateCallback != nil {
		AllModelStore.AllModelsStructCreateCallback.CreateORMVertice(Vertice)
	}
}


// DeleteORMVertice enables dynamic registration of a Vertice instance
func DeleteORMVertice(Vertice *Vertice) {
	for index, _Vertice := range AllModelStore.Vertices {
		if _Vertice == Vertice {
			AllModelStore.Vertices[index] = AllModelStore.Vertices[len(AllModelStore.Vertices)-1]
			AllModelStore.Vertices = AllModelStore.Vertices[:len(AllModelStore.Vertices)-1]
		}
	}
	if AllModelStore.AllModelsStructDeleteCallback != nil {
		AllModelStore.AllModelsStructDeleteCallback.DeleteORMVertice(Vertice)
	}
}

// swagger:ignore
type AllModelsStructCreateInterface interface {
	CreateORMClassdiagram(Classdiagram *Classdiagram)
	CreateORMClassshape(Classshape *Classshape)
	CreateORMField(Field *Field)
	CreateORMGorgoaction(Gorgoaction *Gorgoaction)
	CreateORMLink(Link *Link)
	CreateORMPkgelt(Pkgelt *Pkgelt)
	CreateORMPosition(Position *Position)
	CreateORMState(State *State)
	CreateORMUmlsc(Umlsc *Umlsc)
	CreateORMVertice(Vertice *Vertice)
}
	
type AllModelsStructDeleteInterface interface {
	DeleteORMClassdiagram(Classdiagram *Classdiagram)
	DeleteORMClassshape(Classshape *Classshape)
	DeleteORMField(Field *Field)
	DeleteORMGorgoaction(Gorgoaction *Gorgoaction)
	DeleteORMLink(Link *Link)
	DeleteORMPkgelt(Pkgelt *Pkgelt)
	DeleteORMPosition(Position *Position)
	DeleteORMState(State *State)
	DeleteORMUmlsc(Umlsc *Umlsc)
	DeleteORMVertice(Vertice *Vertice)
}
		
// swagger:ignore
type AllModelStoreStruct struct {
	Classdiagrams  []*Classdiagram
	Classshapes  []*Classshape
	Fields  []*Field
	Gorgoactions  []*Gorgoaction
	Links  []*Link
	Pkgelts  []*Pkgelt
	Positions  []*Position
	States  []*State
	Umlscs  []*Umlsc
	Vertices  []*Vertice
	AllModelsStructCreateCallback AllModelsStructCreateInterface
	AllModelsStructDeleteCallback AllModelsStructDeleteInterface
}

// swagger:ignore
var AllModelStore AllModelStoreStruct = AllModelStoreStruct{
	Classdiagrams:	make([]*Classdiagram, 0),
	Classshapes:	make([]*Classshape, 0),
	Fields:	make([]*Field, 0),
	Gorgoactions:	make([]*Gorgoaction, 0),
	Links:	make([]*Link, 0),
	Pkgelts:	make([]*Pkgelt, 0),
	Positions:	make([]*Position, 0),
	States:	make([]*State, 0),
	Umlscs:	make([]*Umlsc, 0),
	Vertices:	make([]*Vertice, 0),
}

func (allModelStoreStruct * AllModelStoreStruct) Reset() {
	allModelStoreStruct.Classdiagrams =	make([]*Classdiagram, 0)
	allModelStoreStruct.Classshapes =	make([]*Classshape, 0)
	allModelStoreStruct.Fields =	make([]*Field, 0)
	allModelStoreStruct.Gorgoactions =	make([]*Gorgoaction, 0)
	allModelStoreStruct.Links =	make([]*Link, 0)
	allModelStoreStruct.Pkgelts =	make([]*Pkgelt, 0)
	allModelStoreStruct.Positions =	make([]*Position, 0)
	allModelStoreStruct.States =	make([]*State, 0)
	allModelStoreStruct.Umlscs =	make([]*Umlsc, 0)
	allModelStoreStruct.Vertices =	make([]*Vertice, 0)
}

func (allModelStoreStruct * AllModelStoreStruct) Nil() {
	allModelStoreStruct.Classdiagrams = nil
	allModelStoreStruct.Classshapes = nil
	allModelStoreStruct.Fields = nil
	allModelStoreStruct.Gorgoactions = nil
	allModelStoreStruct.Links = nil
	allModelStoreStruct.Pkgelts = nil
	allModelStoreStruct.Positions = nil
	allModelStoreStruct.States = nil
	allModelStoreStruct.Umlscs = nil
	allModelStoreStruct.Vertices = nil
}
