package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveQuatCoords(expectedW, expectedX, expectedY, expectedZ float32) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(quat sprec.Quat) bool {
			return AreEqualFloat32(quat.X, expectedX) &&
				AreEqualFloat32(quat.Y, expectedY) &&
				AreEqualFloat32(quat.Z, expectedZ) &&
				AreEqualFloat32(quat.W, expectedW)
		},
		func(quat sprec.Quat) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ)
		},
		func(quat sprec.Quat) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ)
		},
	)
}
