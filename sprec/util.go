package sprec

import "math"

const (
	Pi      = float32(math.Pi)
	Tau     = float32(math.Pi * 2.0)
	Epsilon = float32(0.000001)
)

func Abs[T ~float32](value T) T {
	return T(math.Abs(float64(value)))
}

func Max[T ~float32](a, b T) T {
	return max(a, b)
}

func Min[T ~float32](a, b T) T {
	return min(a, b)
}

func Sum[T ~float32](values ...T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

func Sqr[T ~float32](value T) T {
	return value * value
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

func Step[T ~float32](edge, value T) T {
	if value >= edge {
		return 1.0
	}
	return 0.0
}

func Smoothstep[T ~float32](lowerEdge, upperEdge, value T) T {
	if value <= lowerEdge {
		return 0.0
	}
	if value >= upperEdge {
		return 1.0
	}
	fraction := (value - lowerEdge) / (upperEdge - lowerEdge)
	return fraction * fraction * (3.0 - 2.0*fraction)
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

// Sign returns the sign of the value. It returns -1.0 for negative values and
// 1.0 for non-negative values. Keep in mind that zero can be either positive
// or negative, so Sign(0.0) returns 1.0,
// while Sign(math.Copysign(0.0, -1.0)) returns -1.0.
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

func MoveTowards[T ~float32](current, target, maxDelta T) T {
	diff := target - current
	if diff > 0 {
		return current + Clamp(maxDelta, 0.0, diff)
	} else {
		return current - Clamp(maxDelta, 0.0, -diff)
	}
}
