package sprec

import "math"

const (
	Pi      = float32(math.Pi)
	Epsilon = float32(0.000001)
)

func Abs(value float32) float32 {
	return math.Float32frombits(math.Float32bits(value) &^ (1 << 31))
}

func Eq(a, b float32) bool {
	return EqEps(a, b, Epsilon)
}

func EqEps(a, b, epsilon float32) bool {
	return Abs(a-b) < epsilon
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}
