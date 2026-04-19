package sprec

import (
	"fmt"
	"math"
)

// NewVec2 creates a Vec2 with the given X and Y components.
func NewVec2(x, y float32) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

// ZeroVec2 returns the zero Vec2.
func ZeroVec2() Vec2 {
	return Vec2{}
}

// BasisXVec2 returns the unit vector along the X axis.
func BasisXVec2() Vec2 {
	return Vec2{
		X: 1.0,
		Y: 0.0,
	}
}

// BasisYVec2 returns the unit vector along the Y axis.
func BasisYVec2() Vec2 {
	return Vec2{
		X: 0.0,
		Y: 1.0,
	}
}

// Vec2Sum returns the sum of two vectors.
func Vec2Sum(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

// Vec2MultiSum returns the sum of multiple vectors.
func Vec2MultiSum(first Vec2, others ...Vec2) Vec2 {
	result := first
	for _, other := range others {
		result.X += other.X
		result.Y += other.Y
	}
	return result
}

// Vec2Diff returns the difference of two vectors (a - b).
func Vec2Diff(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

// Vec2MultiDiff subtracts each subsequent vector from the first.
func Vec2MultiDiff(first Vec2, others ...Vec2) Vec2 {
	result := first
	for _, other := range others {
		result.X -= other.X
		result.Y -= other.Y
	}
	return result
}

// Vec2Prod multiplies a vector by a scalar value.
func Vec2Prod(vector Vec2, value float32) Vec2 {
	return Vec2{
		X: vector.X * value,
		Y: vector.Y * value,
	}
}

// Vec2Quot divides a vector by a scalar value.
func Vec2Quot(vector Vec2, value float32) Vec2 {
	invValue := 1.0 / value
	return Vec2{
		X: vector.X * invValue,
		Y: vector.Y * invValue,
	}
}

// Vec2Dot returns the dot product of two vectors.
func Vec2Dot(a, b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

// Vec2Cross returns the 2D cross product (scalar) of two vectors.
func Vec2Cross(a, b Vec2) float32 {
	return a.X*b.Y - a.Y*b.X
}

// Vec2Lerp returns the linear interpolation between a and b using t.
// A value of t=0 returns a and t=1 returns b.
func Vec2Lerp(a, b Vec2, t float32) Vec2 {
	return Vec2{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
	}
}

// UnitVec2 returns the unit (normalized) vector in the direction of vector.
func UnitVec2(vector Vec2) Vec2 {
	return Vec2Quot(vector, vector.Length())
}

// ResizedVec2 returns a vector in the same direction as vector but with
// the given length.
func ResizedVec2(vector Vec2, newLength float32) Vec2 {
	ratio := newLength / vector.Length()
	return Vec2Prod(vector, ratio)
}

// InverseVec2 returns the negation of the given vector.
func InverseVec2(vector Vec2) Vec2 {
	return Vec2{
		X: -vector.X,
		Y: -vector.Y,
	}
}

// NormalVec2 returns a unit vector perpendicular to the given vector.
func NormalVec2(vector Vec2) Vec2 {
	return UnitVec2(Vec2{
		X: -vector.Y,
		Y: vector.X,
	})
}

// ArrayToVec2 creates a Vec2 from a two-element array.
func ArrayToVec2(array [2]float32) Vec2 {
	return Vec2{
		X: array[0],
		Y: array[1],
	}
}

// Vec2 is a two-dimensional vector with float32 components.
type Vec2 struct {
	X float32
	Y float32
}

// IsNaN returns true if any component is NaN.
func (v Vec2) IsNaN() bool {
	return math.IsNaN(float64(v.X)) || math.IsNaN(float64(v.Y))
}

// IsInf returns true if any component is Inf.
func (v Vec2) IsInf() bool {
	return math.IsInf(float64(v.X), 0) || math.IsInf(float64(v.Y), 0)
}

// IsZero returns true if all components are within Epsilon of zero.
func (v Vec2) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0)
}

// SqrLength returns the squared length of the vector.
func (v Vec2) SqrLength() float32 {
	return Vec2Dot(v, v)
}

// Length returns the length of the vector.
func (v Vec2) Length() float32 {
	return Sqrt(Vec2Dot(v, v))
}

// String returns a string representation of the vector.
func (v Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
