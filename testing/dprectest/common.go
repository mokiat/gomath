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
	return testing.GenericMatcher(
		func(floatValue float64) bool {
			return AreEqualFloat64(floatValue, expectedValue)
		},
		func(floatValue float64) string {
			return fmt.Sprintf("Expected\n\t%f\nto equal\n\t%f", floatValue, expectedValue)
		},
		func(floatValue float64) string {
			return fmt.Sprintf("Expected\n\t%f\nnot to equal\n\t%f", floatValue, expectedValue)
		},
	)
}
