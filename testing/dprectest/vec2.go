package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec2Coords(expectedX, expectedY float64) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector dprec.Vec2) bool {
			return AreEqualFloat64(vector.X, expectedX) && AreEqualFloat64(vector.Y, expectedY)
		},
		func(vector dprec.Vec2) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f)", vector, expectedX, expectedY)
		},
		func(vector dprec.Vec2) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f)", vector, expectedX, expectedY)
		},
	)
}
