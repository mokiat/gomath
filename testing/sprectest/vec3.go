package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec3Coords(expectedX, expectedY, expectedZ float32) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue interface{}) (testing.MatchStatus, error) {
		vector, ok := actualValue.(sprec.Vec3)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec3Coords matcher expects a sprec.Vec3")
		}

		matches := AreEqualFloat32(vector.X, expectedX) &&
			AreEqualFloat32(vector.Y, expectedY) &&
			AreEqualFloat32(vector.Z, expectedZ)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
