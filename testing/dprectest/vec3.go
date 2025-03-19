package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec3Coords(expectedX, expectedY, expectedZ float64) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		vector, ok := actualValue.(dprec.Vec3)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec3Coords matcher expects a dprec.Vec3")
		}

		matches := AreEqualFloat64(vector.X, expectedX) &&
			AreEqualFloat64(vector.Y, expectedY) &&
			AreEqualFloat64(vector.Z, expectedZ)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f)", vector, expectedX, expectedY, expectedZ),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
