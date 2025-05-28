package sprec

import (
	"fmt"
	"math"
)

func NewVec2(x, y float32) Vec2 {
	return Vec2{
		X: x,
		Y: y,
	}
}

func ZeroVec2() Vec2 {
	return Vec2{}
}

func BasisXVec2() Vec2 {
	return Vec2{
		X: 1.0,
		Y: 0.0,
	}
}

func BasisYVec2() Vec2 {
	return Vec2{
		X: 0.0,
		Y: 1.0,
	}
}

func Vec2Sum(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func Vec2MultiSum(first Vec2, others ...Vec2) Vec2 {
	result := first
	for _, other := range others {
		result.X += other.X
		result.Y += other.Y
	}
	return result
}

func Vec2Diff(a, b Vec2) Vec2 {
	return Vec2{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func Vec2MultiDiff(first Vec2, others ...Vec2) Vec2 {
	result := first
	for _, other := range others {
		result.X -= other.X
		result.Y -= other.Y
	}
	return result
}

func Vec2Prod(vector Vec2, value float32) Vec2 {
	return Vec2{
		X: vector.X * value,
		Y: vector.Y * value,
	}
}

func Vec2Quot(vector Vec2, value float32) Vec2 {
	return Vec2{
		X: vector.X / value,
		Y: vector.Y / value,
	}
}

func Vec2Dot(a, b Vec2) float32 {
	return a.X*b.X + a.Y*b.Y
}

func Vec2Cross(a, b Vec2) float32 {
	return a.X*b.Y - a.Y*b.X
}

func Vec2Lerp(a, b Vec2, t float32) Vec2 {
	return Vec2{
		X: (1-t)*a.X + t*b.X,
		Y: (1-t)*a.Y + t*b.Y,
	}
}

func UnitVec2(vector Vec2) Vec2 {
	return Vec2Quot(vector, vector.Length())
}

func ResizedVec2(vector Vec2, newLength float32) Vec2 {
	ratio := newLength / vector.Length()
	return Vec2Prod(vector, ratio)
}

func InverseVec2(vector Vec2) Vec2 {
	return Vec2{
		X: -vector.X,
		Y: -vector.Y,
	}
}

func NormalVec2(vector Vec2) Vec2 {
	return UnitVec2(Vec2{
		X: -vector.Y,
		Y: vector.X,
	})
}

func ArrayToVec2(array [2]float32) Vec2 {
	return Vec2{
		X: array[0],
		Y: array[1],
	}
}

type Vec2 struct {
	X float32
	Y float32
}

func (v Vec2) IsNaN() bool {
	return math.IsNaN(float64(v.X)) || math.IsNaN(float64(v.Y))
}

func (v Vec2) IsInf() bool {
	return math.IsInf(float64(v.X), 0) || math.IsInf(float64(v.Y), 0)
}

func (v Vec2) IsZero() bool {
	return Eq(v.X, 0.0) && Eq(v.Y, 0.0)
}

func (v Vec2) SqrLength() float32 {
	return Vec2Dot(v, v)
}

func (v Vec2) Length() float32 {
	return Sqrt(Vec2Dot(v, v))
}

func (v Vec2) GoString() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
