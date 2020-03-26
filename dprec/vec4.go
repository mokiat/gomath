package dprec

import "fmt"

func NewVec4(x, y, z, w float64) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

func ZeroVec4() Vec4 {
	return Vec4{}
}

func Vec4Sum(a, b Vec4) Vec4 {
	return Vec4{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func Vec4Diff(a, b Vec4) Vec4 {
	return Vec4{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func Vec4Prod(vector Vec4, value float64) Vec4 {
	return Vec4{
		X: vector.X * value,
		Y: vector.Y * value,
		Z: vector.Z * value,
		W: vector.W * value,
	}
}

func Vec4Quot(vector Vec4, value float64) Vec4 {
	return Vec4{
		X: vector.X / value,
		Y: vector.Y / value,
		Z: vector.Z / value,
		W: vector.W / value,
	}
}

func InverseVec4(vector Vec4) Vec4 {
	return Vec4{
		X: -vector.X,
		Y: -vector.Y,
		Z: -vector.Z,
		W: -vector.W,
	}
}

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v Vec4) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0) && Eq(v.Z, 0.0) && Eq(v.W, 0.0)
}

func (v Vec4) VecXYZ() Vec3 {
	return NewVec3(v.X, v.Y, v.Z)
}

func (v Vec4) GoString() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
