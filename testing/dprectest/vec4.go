package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec4Coords(expectedX, expectedY, expectedZ, expectedW float64) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector dprec.Vec4) bool {
			return AreEqualFloat64(vector.X, expectedX) &&
				AreEqualFloat64(vector.Y, expectedY) &&
				AreEqualFloat64(vector.Z, expectedZ) &&
				AreEqualFloat64(vector.W, expectedW)
		},
		func(vector dprec.Vec4) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW)
		},
		func(vector dprec.Vec4) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW)
		},
	)
}
