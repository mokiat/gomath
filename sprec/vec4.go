package sprec

import (
	"fmt"
	"math"
)

// NewVec4 creates a Vec4 with the given X, Y, Z, and W components.
func NewVec4(x, y, z, w float32) Vec4 {
	return Vec4{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// ZeroVec4 returns the zero Vec4.
func ZeroVec4() Vec4 {
	return Vec4{}
}

// Vec4Sum returns the sum of two vectors.
func Vec4Sum(a, b Vec4) Vec4 {
	return Vec4{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

// Vec4MultiSum returns the sum of multiple vectors.
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

// Vec4Diff returns the difference of two vectors (a - b).
func Vec4Diff(a, b Vec4) Vec4 {
	return Vec4{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

// Vec4MultiDiff subtracts each subsequent vector from the first.
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

// Vec4Prod multiplies a vector by a scalar value.
func Vec4Prod(vector Vec4, value float32) Vec4 {
	return Vec4{
		X: vector.X * value,
		Y: vector.Y * value,
		Z: vector.Z * value,
		W: vector.W * value,
	}
}

// Vec4Quot divides a vector by a scalar value.
func Vec4Quot(vector Vec4, value float32) Vec4 {
	invValue := 1.0 / value
	return Vec4{
		X: vector.X * invValue,
		Y: vector.Y * invValue,
		Z: vector.Z * invValue,
		W: vector.W * invValue,
	}
}

// Vec4Dot returns the dot product of two vectors.
func Vec4Dot(a, b Vec4) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// Vec4Lerp returns the linear interpolation between a and b using t.
// A value of t=0 returns a and t=1 returns b.
func Vec4Lerp(a, b Vec4, t float32) Vec4 {
	return Vec4{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
		Z: a.Z + t*(b.Z-a.Z),
		W: a.W + t*(b.W-a.W),
	}
}

// InverseVec4 returns the negation of the given vector.
func InverseVec4(vector Vec4) Vec4 {
	return Vec4{
		X: -vector.X,
		Y: -vector.Y,
		Z: -vector.Z,
		W: -vector.W,
	}
}

// ArrayToVec4 creates a Vec4 from a four-element array.
func ArrayToVec4(array [4]float32) Vec4 {
	return Vec4{
		X: array[0],
		Y: array[1],
		Z: array[2],
		W: array[3],
	}
}

// Vec4 is a four-dimensional vector with float32 components.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

// IsNaN returns true if any component is NaN.
func (v Vec4) IsNaN() bool {
	return math.IsNaN(float64(v.X)) || math.IsNaN(float64(v.Y)) || math.IsNaN(float64(v.Z)) || math.IsNaN(float64(v.W))
}

// IsInf returns true if any component is Inf.
func (v Vec4) IsInf() bool {
	return math.IsInf(float64(v.X), 0) || math.IsInf(float64(v.Y), 0) || math.IsInf(float64(v.Z), 0) || math.IsInf(float64(v.W), 0)
}

// IsZero returns true if all components are within Epsilon of zero.
func (v Vec4) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0) && Eq(v.Z, 0.0) && Eq(v.W, 0.0)
}

// VecXYZ returns the X, Y, and Z components as a Vec3.
func (v Vec4) VecXYZ() Vec3 {
	return NewVec3(v.X, v.Y, v.Z)
}

// Array returns the vector components as an array.
func (v Vec4) Array() [4]float32 {
	return [4]float32{v.X, v.Y, v.Z, v.W}
}

// String returns a string representation of the vector.
func (v Vec4) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", v.X, v.Y, v.Z, v.W)
}
