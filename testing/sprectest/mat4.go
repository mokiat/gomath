package sprectest

import (
	"fmt"

	"github.com/mokiat/gomath/sprec"
	"github.com/mokiat/gomath/testing"
	"github.com/onsi/gomega/types"
)

func HaveMat4Elements(
	e11, e12, e13, e14 float32,
	e21, e22, e23, e24 float32,
	e31, e32, e33, e34 float32,
	e41, e42, e43, e44 float32,
) types.GomegaMatcher {
	return testing.SimpleMatcher(func(actualValue any) (testing.MatchStatus, error) {
		matrix, ok := actualValue.(sprec.Mat4)
		if !ok {
			return testing.MatchStatus{}, fmt.Errorf("HaveMat4Elements matcher expects a sprec.Mat4")
		}

		matches := AreEqualFloat32(matrix.M11, e11) &&
			AreEqualFloat32(matrix.M12, e12) &&
			AreEqualFloat32(matrix.M13, e13) &&
			AreEqualFloat32(matrix.M14, e14) &&
			AreEqualFloat32(matrix.M21, e21) &&
			AreEqualFloat32(matrix.M22, e22) &&
			AreEqualFloat32(matrix.M23, e23) &&
			AreEqualFloat32(matrix.M24, e24) &&
			AreEqualFloat32(matrix.M31, e31) &&
			AreEqualFloat32(matrix.M32, e32) &&
			AreEqualFloat32(matrix.M33, e33) &&
			AreEqualFloat32(matrix.M34, e34) &&
			AreEqualFloat32(matrix.M41, e41) &&
			AreEqualFloat32(matrix.M42, e42) &&
			AreEqualFloat32(matrix.M43, e43) &&
			AreEqualFloat32(matrix.M44, e44)
		if !matches {
			return testing.FailureMatchStatus(
				fmt.Sprintf(`Expected
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)					
to have elements
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)`,
					matrix.M11, matrix.M12, matrix.M13, matrix.M14,
					matrix.M21, matrix.M22, matrix.M23, matrix.M24,
					matrix.M31, matrix.M32, matrix.M33, matrix.M34,
					matrix.M41, matrix.M42, matrix.M43, matrix.M44,
					e11, e12, e13, e14,
					e21, e22, e23, e24,
					e31, e32, e33, e34,
					e41, e42, e43, e44,
				),
				fmt.Sprintf(`Expected
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)					
not to have elements
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)
	(%f, %f, %f, %f)`,
					matrix.M11, matrix.M12, matrix.M13, matrix.M14,
					matrix.M21, matrix.M22, matrix.M23, matrix.M24,
					matrix.M31, matrix.M32, matrix.M33, matrix.M34,
					matrix.M41, matrix.M42, matrix.M43, matrix.M44,
					e11, e12, e13, e14,
					e21, e22, e23, e24,
					e31, e32, e33, e34,
					e41, e42, e43, e44,
				),
			), nil
		}
		return testing.SuccessMatchStatus(), nil
	})
}
