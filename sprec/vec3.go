package sprec

import (
	"fmt"
	"math"
)

// NewVec3 creates a Vec3 with the given X, Y, and Z components.
func NewVec3(x, y, z float32) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

// ZeroVec3 returns the zero Vec3.
func ZeroVec3() Vec3 {
	return Vec3{}
}

// BasisXVec3 returns the unit vector along the X axis.
func BasisXVec3() Vec3 {
	return Vec3{
		X: 1.0,
		Y: 0.0,
		Z: 0.0,
	}
}

// BasisYVec3 returns the unit vector along the Y axis.
func BasisYVec3() Vec3 {
	return Vec3{
		X: 0.0,
		Y: 1.0,
		Z: 0.0,
	}
}

// BasisZVec3 returns the unit vector along the Z axis.
func BasisZVec3() Vec3 {
	return Vec3{
		X: 0.0,
		Y: 0.0,
		Z: 1.0,
	}
}

// Vec3Sum returns the sum of two vectors.
func Vec3Sum(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

// Vec3MultiSum returns the sum of multiple vectors.
func Vec3MultiSum(first Vec3, others ...Vec3) Vec3 {
	result := first
	for _, other := range others {
		result.X += other.X
		result.Y += other.Y
		result.Z += other.Z
	}
	return result
}

// Vec3Diff returns the difference of two vectors (a - b).
func Vec3Diff(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

// Vec3MultiDiff subtracts each subsequent vector from the first.
func Vec3MultiDiff(first Vec3, others ...Vec3) Vec3 {
	result := first
	for _, other := range others {
		result.X -= other.X
		result.Y -= other.Y
		result.Z -= other.Z
	}
	return result
}

// Vec3Prod multiplies a vector by a scalar value.
func Vec3Prod(vector Vec3, value float32) Vec3 {
	return Vec3{
		X: vector.X * value,
		Y: vector.Y * value,
		Z: vector.Z * value,
	}
}

// Vec3Quot divides a vector by a scalar value.
func Vec3Quot(vector Vec3, value float32) Vec3 {
	invValue := 1.0 / value
	return Vec3{
		X: vector.X * invValue,
		Y: vector.Y * invValue,
		Z: vector.Z * invValue,
	}
}

// Vec3Dot returns the dot product of two vectors.
func Vec3Dot(a, b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Vec3Cross returns the cross product of two vectors.
func Vec3Cross(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

// Vec3Lerp returns the linear interpolation between a and b using t.
// A value of t=0 returns a and t=1 returns b.
func Vec3Lerp(a, b Vec3, t float32) Vec3 {
	return Vec3{
		X: a.X + t*(b.X-a.X),
		Y: a.Y + t*(b.Y-a.Y),
		Z: a.Z + t*(b.Z-a.Z),
	}
}

// UnitVec3 returns the unit (normalized) vector in the direction of vector.
func UnitVec3(vector Vec3) Vec3 {
	return Vec3Quot(vector, vector.Length())
}

// ResizedVec3 returns a vector in the same direction as vector but with
// the given length.
func ResizedVec3(vector Vec3, newLength float32) Vec3 {
	ratio := newLength / vector.Length()
	return Vec3Prod(vector, ratio)
}

// InverseVec3 returns the negation of the given vector.
func InverseVec3(vector Vec3) Vec3 {
	return Vec3{
		X: -vector.X,
		Y: -vector.Y,
		Z: -vector.Z,
	}
}

// NormalVec3 returns a unit vector perpendicular to the given vector.
func NormalVec3(vector Vec3) Vec3 {
	sqrX := vector.X * vector.X
	sqrY := vector.Y * vector.Y
	sqrZ := vector.Z * vector.Z
	if (sqrZ > sqrX) && (sqrZ > sqrY) {
		return UnitVec3(Vec3{
			X: 1.0,
			Y: 1.0,
			Z: -(vector.X + vector.Y) / vector.Z,
		})
	} else {
		if sqrX > sqrY {
			return UnitVec3(Vec3{
				X: -(vector.Y + vector.Z) / vector.X,
				Y: 1.0,
				Z: 1.0,
			})
		} else {
			return UnitVec3(Vec3{
				X: 1.0,
				Y: -(vector.X + vector.Z) / vector.Y,
				Z: 1.0,
			})
		}
	}
}

// Vec3Angle returns the shortest angle between two vectors. It always
// returns a positive angle.
func Vec3Angle(a, b Vec3) Angle {
	dot := Vec3Dot(a, b)
	cross := Vec3Cross(a, b)
	return Atan2(cross.Length(), dot)
}

// Vec3Projection returns the specified vector flattened along the specified
// normal. The normal must be a unit vector. The result is the projection of
// the vector onto the plane defined by the normal.
func Vec3Projection(vector Vec3, normal Vec3) Vec3 {
	dot := Vec3Dot(vector, normal)
	return Vec3Diff(vector, Vec3Prod(normal, dot))
}

// Vec3ProjectionAngle returns the angle between two vectors projected onto
// a plane defined by a normal vector. Unlike Vec3Angle, this function
// returns a signed angle and the ordering of the vectors matters.
func Vec3ProjectionAngle(a, b, normal Vec3) Angle {
	flatA := UnitVec3(Vec3Projection(a, normal))
	flatB := UnitVec3(Vec3Projection(b, normal))
	dot := Vec3Dot(flatA, flatB)
	cross := Vec3Cross(flatA, flatB)
	return Atan2(Vec3Dot(cross, normal), dot)
}

// ArrayToVec3 creates a Vec3 from a three-element array.
func ArrayToVec3(array [3]float32) Vec3 {
	return Vec3{
		X: array[0],
		Y: array[1],
		Z: array[2],
	}
}

// Vec3 is a three-dimensional vector with float32 components.
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

// IsNaN returns true if any component is NaN.
func (v Vec3) IsNaN() bool {
	return math.IsNaN(float64(v.X)) || math.IsNaN(float64(v.Y)) || math.IsNaN(float64(v.Z))
}

// IsInf returns true if any component is Inf.
func (v Vec3) IsInf() bool {
	return math.IsInf(float64(v.X), 0) || math.IsInf(float64(v.Y), 0) || math.IsInf(float64(v.Z), 0)
}

// IsZero returns true if all components are within Epsilon of zero.
func (v Vec3) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0) && Eq(v.Z, 0.0)
}

// SqrLength returns the squared length of the vector.
func (v Vec3) SqrLength() float32 {
	return Vec3Dot(v, v)
}

// Length returns the length of the vector.
func (v Vec3) Length() float32 {
	return Sqrt(Vec3Dot(v, v))
}

// Array returns the vector components as an array.
func (v Vec3) Array() [3]float32 {
	return [3]float32{v.X, v.Y, v.Z}
}

// String returns a string representation of the vector.
func (v Vec3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}
