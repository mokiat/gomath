package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec3Coords(expectedX, expectedY, expectedZ float64) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector dprec.Vec3) bool {
			return AreEqualFloat64(vector.X, expectedX) &&
				AreEqualFloat64(vector.Y, expectedY) &&
				AreEqualFloat64(vector.Z, expectedZ)
		},
		func(vector dprec.Vec3) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ)
		},
		func(vector dprec.Vec3) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ)
		},
	)
}
