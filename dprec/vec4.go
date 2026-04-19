package dprec

import (
	"fmt"
	"math"
)

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

func Vec4MultiSum(first Vec4, others ...Vec4) Vec4 {
	result := first
	for _, other := range others {
		result.X += other.X
		result.Y += other.Y
		result.Z += other.Z
		result.W += other.W
	}
	return result
}

func Vec4Diff(a, b Vec4) Vec4 {
	return Vec4{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func Vec4MultiDiff(first Vec4, others ...Vec4) Vec4 {
	result := first
	for _, other := range others {
		result.X -= other.X
		result.Y -= other.Y
		result.Z -= other.Z
		result.W -= other.W
	}
	return result
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
	invValue := 1.0 / value
	return Vec4{
		X: vector.X * invValue,
		Y: vector.Y * invValue,
		Z: vector.Z * invValue,
		W: vector.W * invValue,
	}
}

func Vec4Dot(a, b Vec4) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func Vec4Lerp(a, b Vec4, t float64) Vec4 {
	return Vec4{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
		Z: a.Z + t*(b.Z-a.Z),
		W: a.W + t*(b.W-a.W),
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

func ArrayToVec4(array [4]float64) Vec4 {
	return Vec4{
		X: array[0],
		Y: array[1],
		Z: array[2],
		W: array[3],
	}
}

type Vec4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v Vec4) IsNaN() bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y) || math.IsNaN(v.Z) || math.IsNaN(v.W)
}

func (v Vec4) IsInf() bool {
	return math.IsInf(v.X, 0) || math.IsInf(v.Y, 0) || math.IsInf(v.Z, 0) || math.IsInf(v.W, 0)
}

func (v Vec4) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0) && Eq(v.Z, 0.0) && Eq(v.W, 0.0)
}

func (v Vec4) VecXYZ() Vec3 {
	return NewVec3(v.X, v.Y, v.Z)
}

func (v Vec4) Array() [4]float64 {
	return [4]float64{v.X, v.Y, v.Z, v.W}
}

func (v Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
