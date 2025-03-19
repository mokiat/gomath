package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveQuatCoords(expectedW, expectedX, expectedY, expectedZ float64) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		quat, ok := actualValue.(dprec.Quat)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveQuatCoords matcher expects a dprec.Quat")
		}

		matches := AreEqualFloat64(quat.X, expectedX) &&
			AreEqualFloat64(quat.Y, expectedY) &&
			AreEqualFloat64(quat.Z, expectedZ) &&
			AreEqualFloat64(quat.W, expectedW)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%#v\nto have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ),
				fmt.Sprintf("Expected\n\t%#v\nnot to have coords\n\t(%f, %f, %f, %f)", quat, expectedW, expectedX, expectedY, expectedZ),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
