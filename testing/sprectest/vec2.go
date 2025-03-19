package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec2Coords(expectedX, expectedY float32) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		vector, ok := actualValue.(sprec.Vec2)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec2Coords matcher expects a sprec.Vec2")
		}

		matches := AreEqualFloat32(vector.X, expectedX) && AreEqualFloat32(vector.Y, expectedY)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f)", vector, expectedX, expectedY),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f)", vector, expectedX, expectedY),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
