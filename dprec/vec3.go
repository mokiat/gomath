package dprec

import (
	"fmt"
	"math"
)

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func ZeroVec3() Vec3 {
	return Vec3{}
}

func BasisXVec3() Vec3 {
	return Vec3{
		X: 1.0,
		Y: 0.0,
		Z: 0.0,
	}
}

func BasisYVec3() Vec3 {
	return Vec3{
		X: 0.0,
		Y: 1.0,
		Z: 0.0,
	}
}

func BasisZVec3() Vec3 {
	return Vec3{
		X: 0.0,
		Y: 0.0,
		Z: 1.0,
	}
}

func Vec3Sum(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
	}
}

func Vec3MultiSum(first Vec3, others ...Vec3) Vec3 {
	result := first
	for _, other := range others {
		result.X += other.X
		result.Y += other.Y
		result.Z += other.Z
	}
	return result
}

func Vec3Diff(a, b Vec3) Vec3 {
	return Vec3{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
	}
}

func Vec3MultiDiff(first Vec3, others ...Vec3) Vec3 {
	result := first
	for _, other := range others {
		result.X -= other.X
		result.Y -= other.Y
		result.Z -= other.Z
	}
	return result
}

func Vec3Prod(vector Vec3, value float64) Vec3 {
	return Vec3{
		X: vector.X * value,
		Y: vector.Y * value,
		Z: vector.Z * value,
	}
}

func Vec3Quot(vector Vec3, value float64) Vec3 {
	return Vec3{
		X: vector.X / value,
		Y: vector.Y / value,
		Z: vector.Z / value,
	}
}

func Vec3Dot(a, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Vec3Cross(a, b Vec3) Vec3 {
	return Vec3{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}

func Vec3Lerp(a, b Vec3, t float64) Vec3 {
	return Vec3{
		X: (1-t)*a.X + t*b.X,
		Y: (1-t)*a.Y + t*b.Y,
		Z: (1-t)*a.Z + t*b.Z,
	}
}

func UnitVec3(vector Vec3) Vec3 {
	return Vec3Quot(vector, vector.Length())
}

func ResizedVec3(vector Vec3, newLength float64) Vec3 {
	ratio := newLength / vector.Length()
	return Vec3Prod(vector, ratio)
}

func InverseVec3(vector Vec3) Vec3 {
	return Vec3{
		X: -vector.X,
		Y: -vector.Y,
		Z: -vector.Z,
	}
}

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
	a = UnitVec3(a)
	b = UnitVec3(b)
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

func ArrayToVec3(array [3]float64) Vec3 {
	return Vec3{
		X: array[0],
		Y: array[1],
		Z: array[2],
	}
}

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vec3) IsNaN() bool {
	return math.IsNaN(v.X) || math.IsNaN(v.Y) || math.IsNaN(v.Z)
}

func (v Vec3) IsInf() bool {
	return math.IsInf(v.X, 0) || math.IsInf(v.Y, 0) || math.IsInf(v.Z, 0)
}

func (v Vec3) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0) && Eq(v.Z, 0.0)
}

func (v Vec3) SqrLength() float64 {
	return Vec3Dot(v, v)
}

func (v Vec3) Length() float64 {
	return Sqrt(Vec3Dot(v, v))
}

func (v Vec3) Array() [3]float64 {
	return [3]float64{v.X, v.Y, v.Z}
}

func (v Vec3) GoString() string {
	return fmt.Sprintf("(%f, %f, %f)", v.X, v.Y, v.Z)
}
