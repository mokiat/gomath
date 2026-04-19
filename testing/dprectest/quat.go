package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveQuatCoords(expectedW, expectedX, expectedY, expectedZ float64) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(quat dprec.Quat) bool {
			return AreEqualFloat64(quat.X, expectedX) &&
				AreEqualFloat64(quat.Y, expectedY) &&
				AreEqualFloat64(quat.Z, expectedZ) &&
				AreEqualFloat64(quat.W, expectedW)
		},
		func(quat dprec.Quat) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ)
		},
		func(quat dprec.Quat) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ)
		},
	)
}
