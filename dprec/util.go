package dprec

import "math"

const (
	Pi      = float64(math.Pi)
	Tau     = float64(math.Pi * 2.0)
	Epsilon = float64(0.000000000001)
)

func Abs[T ~float64](value T) T {
	return T(math.Abs(float64(value)))
}

func Max[T ~float64](a, b T) T {
	return max(a, b)
}

func Min[T ~float64](a, b T) T {
	return min(a, b)
}

func Floor[T ~float64](value T) T {
	return T(math.Floor(float64(value)))
}

func Ceil[T ~float64](value T) T {
	return T(math.Ceil(float64(value)))
}

func Clamp[T ~float64](value, lower, upper T) T {
	return min(max(lower, value), upper)
}

func Mix[T ~float64](a, b T, amount float64) T {
	return T(float64(a)*(1.0-amount) + float64(b)*amount)
}

func Eq(a, b float64) bool {
	return EqEps(a, b, Epsilon)
}

func EqEps(a, b, epsilon float64) bool {
	return Abs(a-b) < epsilon
}

func Mod(a, b float64) float64 {
	return float64(math.Mod(a, b))
}

func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func Cos(angle Angle) float64 {
	return math.Cos(angle.Radians())
}

func Acos(cs float64) Angle {
	return Radians(math.Acos(cs))
}

func Sin(angle Angle) float64 {
	return math.Sin(angle.Radians())
}

func Asin(sn float64) Angle {
	return Radians(math.Asin(sn))
}

func Tan(angle Angle) float64 {
	return math.Tan(angle.Radians())
}

func Atan2(y, x float64) Angle {
	return Radians(math.Atan2(y, x))
}

func Sign(value float64) float64 {
	if math.Signbit(value) {
		return -1.0
	}
	return 1.0
}

func IsNegative[T ~float64](value T) bool {
	return math.Signbit(float64(value))
}

func IsValid[T ~float64](value T) bool {
	return !math.IsNaN(float64(value)) && !math.IsInf(float64(value), 0)
}
