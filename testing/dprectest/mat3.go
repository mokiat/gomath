package dprectest

import (
	"fmt"

	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveMat3Elements(
	e11, e12, e13 float64,
	e21, e22, e23 float64,
	e31, e32, e33 float64,
) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue interface{}) (testing.MatchStatus, error) {
		matrix, ok := actualValue.(dprec.Mat3)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveMat3Elements matcher expects a dprec.Mat3")
		}

		matches := AreEqualFloat64(matrix.M11, e11) &&
			AreEqualFloat64(matrix.M12, e12) &&
			AreEqualFloat64(matrix.M13, e13) &&
			AreEqualFloat64(matrix.M21, e21) &&
			AreEqualFloat64(matrix.M22, e22) &&
			AreEqualFloat64(matrix.M23, e23) &&
			AreEqualFloat64(matrix.M31, e31) &&
			AreEqualFloat64(matrix.M32, e32) &&
			AreEqualFloat64(matrix.M33, e33)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf(`Expected
	(%f, %f, %f)
	(%f, %f, %f)
	(%f, %f, %f)
to have elements
	(%f, %f, %f)
	(%f, %f, %f)
	(%f, %f, %f)`,
					matrix.M11, matrix.M12, matrix.M13,
					matrix.M21, matrix.M22, matrix.M23,
					matrix.M31, matrix.M32, matrix.M33,
					e11, e12, e13,
					e21, e22, e23,
					e31, e32, e33,
				),
				fmt.Sprintf(`Expected
	(%f, %f, %f)
	(%f, %f, %f)
	(%f, %f, %f)
not to have elements
	(%f, %f, %f)
	(%f, %f, %f)
	(%f, %f, %f)`,
					matrix.M11, matrix.M12, matrix.M13,
					matrix.M21, matrix.M22, matrix.M23,
					matrix.M31, matrix.M32, matrix.M33,
					e11, e12, e13,
					e21, e22, e23,
					e31, e32, e33,
				),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
