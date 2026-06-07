package sprec

import "math"

const (
	// Pi is the mathematical constant π.
	Pi = float32(math.Pi)

	// Tau is 2π.
	Tau = float32(math.Pi * 2.0)

	// Epsilon is the tolerance used for float equality comparisons.
	Epsilon = float32(0.000001)
)

// Abs returns the absolute value of the given value.
func Abs[T ~float32](value T) T {
	return T(math.Abs(float64(value)))
}

// Max returns the larger of the two values.
//
// Deprecated: Use built-in max function instead.
//
//go:fix inline
func Max[T ~float32](a, b T) T {
	return max(a, b)
}

// Min returns the smaller of the two values.
//
// Deprecated: Use built-in min function instead.
//
//go:fix inline
func Min[T ~float32](a, b T) T {
	return min(a, b)
}

// Sum returns the sum of the provided values.
func Sum[T ~float32](values ...T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

// Sqr returns the square of value.
func Sqr[T ~float32](value T) T {
	return value * value
}

// Floor returns the largest integer value less than or equal to value.
func Floor[T ~float32](value T) T {
	return T(math.Floor(float64(value)))
}

// Ceil returns the smallest integer value greater than or equal to value.
func Ceil[T ~float32](value T) T {
	return T(math.Ceil(float64(value)))
}

// Clamp returns value clamped to the range [lower, upper].
func Clamp[T ~float32](value, lower, upper T) T {
	return min(max(lower, value), upper)
}

// Mix performs a linear interpolation between a and b using the given amount.
// An amount of 0.0 returns a and an amount of 1.0 returns b.
func Mix[T ~float32](a, b T, amount float32) T {
	return T(float32(a) + amount*(float32(b)-float32(a)))
}

// Step returns 0.0 if value is less than edge, and 1.0 otherwise.
func Step[T ~float32](edge, value T) T {
	if value >= edge {
		return 1.0
	}
	return 0.0
}

// Smoothstep returns a smooth Hermite interpolation between 0.0 and 1.0
// for value in the range [lowerEdge, upperEdge]. Values outside the range
// are clamped to 0.0 or 1.0.
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

// Eq returns true if the two values are within Epsilon of each other.
func Eq(a, b float32) bool {
	return EqEps(a, b, Epsilon)
}

// EqEps returns true if the two values are within epsilon of each other.
func EqEps(a, b, epsilon float32) bool {
	return Abs(a-b) < epsilon
}

// Mod returns the floating-point remainder of a/b.
func Mod(a, b float32) float32 {
	return float32(math.Mod(float64(a), float64(b)))
}

// Sqrt returns the square root of value.
func Sqrt(value float32) float32 {
	return float32(math.Sqrt(float64(value)))
}

// Pow returns a raised to the power of b.
func Pow(a, b float32) float32 {
	return float32(math.Pow(float64(a), float64(b)))
}

// Cos returns the cosine of the given angle.
func Cos(angle Angle) float32 {
	return float32(math.Cos(float64(angle.Radians())))
}

// Acos returns the angle whose cosine is cs.
func Acos(cs float32) Angle {
	return Radians(float32(math.Acos(float64(cs))))
}

// Sin returns the sine of the given angle.
func Sin(angle Angle) float32 {
	return float32(math.Sin(float64(angle.Radians())))
}

// Asin returns the angle whose sine is sn.
func Asin(sn float32) Angle {
	return Radians(float32(math.Asin(float64(sn))))
}

// Tan returns the tangent of the given angle.
func Tan(angle Angle) float32 {
	return float32(math.Tan(float64(angle.Radians())))
}

// Atan2 returns the angle between the positive x-axis and the point (x, y).
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

// IsNegative returns true if value is negative, including negative zero.
func IsNegative[T ~float32](value T) bool {
	return math.Signbit(float64(value))
}

// IsValid returns true if value is neither NaN nor Inf.
func IsValid[T ~float32](value T) bool {
	return !math.IsNaN(float64(value)) && !math.IsInf(float64(value), 0)
}

// MoveTowards advances current towards target by at most maxDelta. If current
// is already within maxDelta of target, target is returned.
func MoveTowards[T ~float32](current, target, maxDelta T) T {
	diff := target - current
	if diff > 0 {
		return current + Clamp(maxDelta, 0.0, diff)
	} else {
		return current - Clamp(maxDelta, 0.0, -diff)
	}
}
