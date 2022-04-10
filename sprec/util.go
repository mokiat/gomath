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

func Clamp(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Mix(a, b, amount float32) float32 {
	return a*(1.0-amount) + b*amount
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

func Cos(angle Angle) float32 {
	return float32(math.Cos(float64(angle.Radians())))
}

func Sin(angle Angle) float32 {
	return float32(math.Sin(float64(angle.Radians())))
}

func Tan(angle Angle) float32 {
	return float32(math.Tan(float64(angle.Radians())))
}

func Sign(value float32) float32 {
	if math.Signbit(float64(value)) {
		return -1.0
	}
	return 1.0
}
