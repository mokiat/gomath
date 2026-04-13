package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec4Coords(expectedX, expectedY, expectedZ, expectedW float32) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector sprec.Vec4) bool {
			return AreEqualFloat32(vector.X, expectedX) &&
				AreEqualFloat32(vector.Y, expectedY) &&
				AreEqualFloat32(vector.Z, expectedZ) &&
				AreEqualFloat32(vector.W, expectedW)
		},
		func(vector sprec.Vec4) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW)
		},
		func(vector sprec.Vec4) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW)
		},
	)
}
