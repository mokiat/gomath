package dtos

import (
	"github.com/mokiat/gomath/dprec"
	"github.com/mokiat/gomath/sprec"
)

// Vec2 converts the specified double-precision Vec2 into a single-precision
// Vec2.
func Vec2(src dprec.Vec2) sprec.Vec2 {
	return sprec.Vec2{
		X: float32(src.X),
		Y: float32(src.Y),
	}
}

// Vec3 converts the specified double-precision Vec3 into a single-precision
// Vec3.
func Vec3(src dprec.Vec3) sprec.Vec3 {
	return sprec.Vec3{
		X: float32(src.X),
		Y: float32(src.Y),
		Z: float32(src.Z),
	}
}

// Vec4 converts the specified double-precision Vec4 into a single-precision
// Vec4.
func Vec4(src dprec.Vec4) sprec.Vec4 {
	return sprec.Vec4{
		X: float32(src.X),
		Y: float32(src.Y),
		Z: float32(src.Z),
		W: float32(src.W),
	}
}

// Quat converts the specified double-precision Quat into a single-precision
// Quat.
func Quat(src dprec.Quat) sprec.Quat {
	return sprec.Quat{
		W: float32(src.W),
		X: float32(src.X),
		Y: float32(src.Y),
		Z: float32(src.Z),
	}
}

// Mat3 converts the specified double-precision Mat3 into a single-precision
// Mat3.
func Mat3(src dprec.Mat3) sprec.Mat3 {
	return sprec.Mat3{
		M11: float32(src.M11),
		M12: float32(src.M12),
		M13: float32(src.M13),

		M21: float32(src.M21),
		M22: float32(src.M22),
		M23: float32(src.M23),

		M31: float32(src.M31),
		M32: float32(src.M32),
		M33: float32(src.M33),
	}
}

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
