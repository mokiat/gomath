package stod

import (
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/sprec"
)

// Vec2 converts the specified single-precision Vec2 into a double-precision
// Vec2.
func Vec2(src sprec.Vec2) dprec.Vec2 {
	return dprec.Vec2{
		X: float64(src.X),
		Y: float64(src.Y),
	}
}

// Vec3 converts the specified single-precision Vec3 into a double-precision
// Vec3.
func Vec3(src sprec.Vec3) dprec.Vec3 {
	return dprec.Vec3{
		X: float64(src.X),
		Y: float64(src.Y),
		Z: float64(src.Z),
	}
}

// Vec4 converts the specified single-precision Vec4 into a double-precision
// Vec4.
func Vec4(src sprec.Vec4) dprec.Vec4 {
	return dprec.Vec4{
		X: float64(src.X),
		Y: float64(src.Y),
		Z: float64(src.Z),
		W: float64(src.W),
	}
}

// Quat converts the specified single-precision Quat into a double-precision
// Quat.
func Quat(src sprec.Quat) dprec.Quat {
	return dprec.Quat{
		W: float64(src.W),
		X: float64(src.X),
		Y: float64(src.Y),
		Z: float64(src.Z),
	}
}

// Mat3 converts the specified single-precision Mat3 into a double-precision
// Mat3.
func Mat3(src sprec.Mat3) dprec.Mat3 {
	return dprec.Mat3{
		M11: float64(src.M11),
		M12: float64(src.M12),
		M13: float64(src.M13),

		M21: float64(src.M21),
		M22: float64(src.M22),
		M23: float64(src.M23),

		M31: float64(src.M31),
		M32: float64(src.M32),
		M33: float64(src.M33),
	}
}

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
