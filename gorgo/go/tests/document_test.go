package tests

import (
	"testing"

	"github.com/thomaspeugeot/sandbox02/gorgo/go/models"
)

func TestUmarshallDocument(t *testing.T) {

	var pkgelt models.Pkgelt

	// unmarshall diagrams of the document
	pkgelt.Unmarshall("geometry/diagrams")

}
