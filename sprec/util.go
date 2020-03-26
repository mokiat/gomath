package sprec

import "math"

const (
	Pi      = float32(math.Pi)
	Epsilon = float32(0.000001)
)

func Abs(value float32) float32 {
	return math.Float32frombits(math.Float32bits(value) &^ (1 << 31))
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
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

func Cos(radians float32) float32 {
	return float32(math.Cos(float64(radians)))
}

func Sin(radians float32) float32 {
	return float32(math.Sin(float64(radians)))
}
