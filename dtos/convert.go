package dtos

import (
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/sprec"
)

// Mat4 converts the specified double-precision Mat4 into a single-precision
// Mat4.
func Mat4(src dprec.Mat4) sprec.Mat4 {
	return sprec.Mat4{
		M11: float32(src.M11),
		M12: float32(src.M12),
		M13: float32(src.M13),
		M14: float32(src.M14),

		M21: float32(src.M21),
		M22: float32(src.M22),
		M23: float32(src.M23),
		M24: float32(src.M24),

		M31: float32(src.M31),
		M32: float32(src.M32),
		M33: float32(src.M33),
		M34: float32(src.M34),

		M41: float32(src.M41),
		M42: float32(src.M42),
		M43: float32(src.M43),
		M44: float32(src.M44),
	}
}
