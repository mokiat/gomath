package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec2Coords(expectedX, expectedY float32) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(vector sprec.Vec2) bool {
			return AreEqualFloat32(vector.X, expectedX) && AreEqualFloat32(vector.Y, expectedY)
		},
		func(vector sprec.Vec2) string {
			return fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f)", vector, expectedX, expectedY)
		},
		func(vector sprec.Vec2) string {
			return fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f)", vector, expectedX, expectedY)
		},
	)
}
