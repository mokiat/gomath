package stod

import (
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/sprec"
)

// Mat4 converts the specified single-precision Mat4 into a double-precision
// Mat4.
func Mat4(src sprec.Mat4) dprec.Mat4 {
	return dprec.Mat4{
		M11: float64(src.M11),
		M12: float64(src.M12),
		M13: float64(src.M13),
		M14: float64(src.M14),

		M21: float64(src.M21),
		M22: float64(src.M22),
		M23: float64(src.M23),
		M24: float64(src.M24),

		M31: float64(src.M31),
		M32: float64(src.M32),
		M33: float64(src.M33),
		M34: float64(src.M34),

		M41: float64(src.M41),
		M42: float64(src.M42),
		M43: float64(src.M43),
		M44: float64(src.M44),
	}
}
