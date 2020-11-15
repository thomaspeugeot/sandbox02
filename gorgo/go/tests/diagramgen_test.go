package tests

import (
	"testing"

	"github.com/thomaspeugeot/metabaron/libs/gorgo/go/models"
)

func TestGenerateSVG(t *testing.T) {

	var pkgelt models.Pkgelt

	pkgelt.Unmarshall("geometry/diagrams")

}
