package dprectest

import (
	"fmt"
	"math"

	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

const float64Margin = 0.000000000001

func AreEqualFloat64(a, b float64) bool {
	return math.Abs(a-b) < float64Margin
}

func EqualFloat64(expectedValue float64) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue interface{}) (testing.MatchStatus, error) {
		floatValue, ok := actualValue.(float64)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("EqualFloat64 matcher expects a float64")
		}

		if !AreEqualFloat64(floatValue, expectedValue) {
			return testing.FailureMatchStatus(
				fmt.Sprintf("Expected\n\t%f\nto equal\n\t%f", floatValue, expectedValue),
				fmt.Sprintf("Expected\n\t%f\nnot to equal\n\t%f", floatValue, expectedValue),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
