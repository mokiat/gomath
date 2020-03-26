package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveVec4Coords(expectedX, expectedY, expectedZ, expectedW float64) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue interface{}) (testing.MatchStatus, error) {
		vector, ok := actualValue.(dprec.Vec4)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveVec4Coords matcher expects a dprec.Vec4")
		}

		matches := AreEqualFloat64(vector.X, expectedX) &&
			AreEqualFloat64(vector.Y, expectedY) &&
			AreEqualFloat64(vector.Z, expectedZ) &&
			AreEqualFloat64(vector.W, expectedW)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", vector, expectedX, expectedY, expectedZ, expectedW),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
