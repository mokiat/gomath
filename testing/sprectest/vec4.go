package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec4Coords(expectedX, expectedY, expectedZ, expectedW float32) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		vector, ok := actualValue.(sprec.Vec4)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec4Coords matcher expects a sprec.Vec4")
		}

		matches := AreEqualFloat32(vector.X, expectedX) &&
			AreEqualFloat32(vector.Y, expectedY) &&
			AreEqualFloat32(vector.Z, expectedZ) &&
			AreEqualFloat32(vector.W, expectedW)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
