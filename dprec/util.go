package dprec

import "math"

const (
	Pi      = float64(math.Pi)
	Epsilon = float64(0.000000000001)
)

func Abs(value float64) float64 {
	return math.Abs(value)
}

func Eq(a, b float64) bool {
	return EqEps(a, b, Epsilon)
}

func EqEps(a, b, epsilon float64) bool {
	return Abs(a-b) < epsilon
}

func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}
