package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec3Coords(expectedX, expectedY, expectedZ float32) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector sprec.Vec3) bool {
			return AreEqualFloat32(vector.X, expectedX) &&
				AreEqualFloat32(vector.Y, expectedY) &&
				AreEqualFloat32(vector.Z, expectedZ)
		},
		func(vector sprec.Vec3) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ)
		},
		func(vector sprec.Vec3) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ)
		},
	)
}
