package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveMat3Elements(
	e11, e12, e13 float32,
	e21, e22, e23 float32,
	e31, e32, e33 float32,
) types.GomegaMatcher {
	return testing.GenericMatcher(
		func(matrix sprec.Mat3) bool {
			return AreEqualFloat32(matrix.M11, e11) &&
				AreEqualFloat32(matrix.M12, e12) &&
				AreEqualFloat32(matrix.M13, e13) &&
				AreEqualFloat32(matrix.M21, e21) &&
				AreEqualFloat32(matrix.M22, e22) &&
				AreEqualFloat32(matrix.M23, e23) &&
				AreEqualFloat32(matrix.M31, e31) &&
				AreEqualFloat32(matrix.M32, e32) &&
				AreEqualFloat32(matrix.M33, e33)
		},
		func(matrix sprec.Mat3) string {
			return fmt.Sprintf(`Expected
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
			)
		},
		func(matrix sprec.Mat3) string {
			return fmt.Sprintf(`Expected
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
			)
		},
	)
}
