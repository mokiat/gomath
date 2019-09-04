package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec2Coords(expectedX, expectedY float64) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue interface{}) (testing.MatchStatus, error) {
		vector, ok := actualValue.(dprec.Vec2)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec2Coords matcher expects a dprec.Vec2")
		}

		matches := AreEqualFloat64(vector.X, expectedX) && AreEqualFloat64(vector.Y, expectedY)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f)", vector, expectedX, expectedY),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f)", vector, expectedX, expectedY),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
