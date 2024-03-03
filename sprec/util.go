package sprec

import "math"

const (
	Pi      = float32(math.Pi)
	Epsilon = float32(0.000001)
)

func Abs[T ~float32](value T) T {
	return T(math.Float32frombits(math.Float32bits(float32(value)) &^ (1 << 31)))
}

func Max[T ~float32](a, b T) T {
	return max(a, b)
}

func Min[T ~float32](a, b T) T {
	return min(a, b)
}

func Floor[T ~float32](value T) T {
	return T(math.Floor(float64(value)))
}

func Ceil[T ~float32](value T) T {
	return T(math.Ceil(float64(value)))
}

func Clamp[T ~float32](value, lower, upper T) T {
	return min(max(lower, value), upper)
}

func Mix[T ~float32](a, b T, amount float32) T {
	return T(float32(a)*(1.0-amount) + float32(b)*amount)
}

func Eq(a, b float32) bool {
	return EqEps(a, b, Epsilon)
}

func EqEps(a, b, epsilon float32) bool {
	return Abs(a-b) < epsilon
}

func Mod(a, b float32) float32 {
	return float32(math.Mod(float64(a), float64(b)))
}

func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

func Pow(a, b float32) float32 {
	return float32(math.Pow(float64(a), float64(b)))
}

func Cos(angle Angle) float32 {
	return float32(math.Cos(float64(angle.Radians())))
}

func Acos(cs float32) Angle {
	return Radians(float32(math.Acos(float64(cs))))
}

func Sin(angle Angle) float32 {
	return float32(math.Sin(float64(angle.Radians())))
}

func Asin(sn float32) Angle {
	return Radians(float32(math.Asin(float64(sn))))
}

func Tan(angle Angle) float32 {
	return float32(math.Tan(float64(angle.Radians())))
}

func Atan2(y, x float32) Angle {
	return Radians(float32(math.Atan2(float64(y), float64(x))))
}

func Sign(value float32) float32 {
	if math.Signbit(float64(value)) {
		return -1.0
	}
	return 1.0
}

func IsNegative[T ~float32](value T) bool {
	return math.Signbit(float64(value))
}

func IsValid[T ~float32](value T) bool {
	return !math.IsNaN(float64(value)) && !math.IsInf(float64(value), 0)
}
