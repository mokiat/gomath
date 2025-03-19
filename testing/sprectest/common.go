package sprectest

import (
	"fmt"
	"math"

	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

const float32Margin = 0.000001

func AreEqualFloat32(a, b float32) bool {
	return math.Abs(float64(a-b)) < float64(float32Margin)
}

func EqualFloat32(expectedValue float32) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		floatValue, ok := actualValue.(float32)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("EqualFloat32 matcher expects a float32")
		}

		if !AreEqualFloat32(floatValue, expectedValue) {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%f\nto equal\n\t%f", floatValue, expectedValue),
				fmt.Sprintf("Expected\n\t%f\nnot to equal\n\t%f", floatValue, expectedValue),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
