package dprec

import "math"

const (
	Pi      = float64(math.Pi)
	Epsilon = float64(0.000000000001)
)

func Abs(value float64) float64 {
	return math.Abs(value)
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Mix(a, b, amount float64) float64 {
	return a*(1.0-amount) + b*amount
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

func Cos(angle Angle) float64 {
	return math.Cos(angle.Radians())
}

func Sin(angle Angle) float64 {
	return math.Sin(angle.Radians())
}

func Tan(angle Angle) float64 {
	return math.Tan(angle.Radians())
}

func Sign(value float64) float64 {
	if math.Signbit(value) {
		return -1.0
	}
	return 1.0
}
