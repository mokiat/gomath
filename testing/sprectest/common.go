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
	return testing.GenericMatcher(
		func(floatValue float32) bool {
			return AreEqualFloat32(floatValue, expectedValue)
		},
		func(floatValue float32) string {
			return fmt.Sprintf("Expected\n\t%f\nto equal\n\t%f", floatValue, expectedValue)
		},
		func(floatValue float32) string {
			return fmt.Sprintf("Expected\n\t%f\nnot to equal\n\t%f", floatValue, expectedValue)
		},
	)
}
