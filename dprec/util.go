package dprec

import "math"

const (
	// Pi is the mathematical constant π.
	Pi = float64(math.Pi)

	// Tau is 2π.
	Tau = float64(math.Pi * 2.0)

	// Epsilon is the tolerance used for float equality comparisons.
	Epsilon = float64(0.000000000001)
)

// Abs returns the absolute value of the given value.
func Abs[T ~float64](value T) T {
	return T(math.Abs(float64(value)))
}

// Max returns the larger of the two values.
//
// Deprecated: Use built-in max function instead.
//
//go:fix inline
func Max[T ~float64](a, b T) T {
	return max(a, b)
}

// Min returns the smaller of the two values.
//
// Deprecated: Use built-in min function instead.
//
//go:fix inline
func Min[T ~float64](a, b T) T {
	return min(a, b)
}

// Sum returns the sum of the provided values.
func Sum[T ~float64](values ...T) T {
	var sum T
	for _, value := range values {
		sum += value
	}
	return sum
}

// Sqr returns the square of value.
func Sqr[T ~float64](value T) T {
	return value * value
}

// Floor returns the largest integer value less than or equal to value.
func Floor[T ~float64](value T) T {
	return T(math.Floor(float64(value)))
}

// Ceil returns the smallest integer value greater than or equal to value.
func Ceil[T ~float64](value T) T {
	return T(math.Ceil(float64(value)))
}

// Clamp returns value clamped to the range [lower, upper].
func Clamp[T ~float64](value, lower, upper T) T {
	return min(max(lower, value), upper)
}

// Mix performs a linear interpolation between a and b using the given amount.
// An amount of 0.0 returns a and an amount of 1.0 returns b.
func Mix[T ~float64](a, b T, amount float64) T {
	return T(float64(a) + amount*(float64(b)-float64(a)))
}

// Step returns 0.0 if value is less than edge, and 1.0 otherwise.
func Step[T ~float64](edge, value T) T {
	if value >= edge {
		return 1.0
	}
	return 0.0
}

// Smoothstep returns a smooth Hermite interpolation between 0.0 and 1.0
// for value in the range [lowerEdge, upperEdge]. Values outside the range
// are clamped to 0.0 or 1.0.
func Smoothstep[T ~float64](lowerEdge, upperEdge, value T) T {
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
func Eq(a, b float64) bool {
	return EqEps(a, b, Epsilon)
}

// EqEps returns true if the two values are within epsilon of each other.
func EqEps(a, b, epsilon float64) bool {
	return Abs(a-b) < epsilon
}

// Mod returns the floating-point remainder of a/b.
func Mod(a, b float64) float64 {
	return float64(math.Mod(a, b))
}

// Sqrt returns the square root of value.
func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

// Pow returns a raised to the power of b.
func Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

// Cos returns the cosine of the given angle.
func Cos(angle Angle) float64 {
	return math.Cos(angle.Radians())
}

// Acos returns the angle whose cosine is cs.
func Acos(cs float64) Angle {
	return Radians(math.Acos(cs))
}

// Sin returns the sine of the given angle.
func Sin(angle Angle) float64 {
	return math.Sin(angle.Radians())
}

// Asin returns the angle whose sine is sn.
func Asin(sn float64) Angle {
	return Radians(math.Asin(sn))
}

// Tan returns the tangent of the given angle.
func Tan(angle Angle) float64 {
	return math.Tan(angle.Radians())
}

// Atan2 returns the angle between the positive x-axis and the point (x, y).
func Atan2(y, x float64) Angle {
	return Radians(math.Atan2(y, x))
}

// Sign returns the sign of the value. It returns -1.0 for negative values and
// 1.0 for non-negative values. Keep in mind that zero can be either positive
// or negative, so Sign(0.0) returns 1.0,
// while Sign(math.Copysign(0.0, -1.0)) returns -1.0.
func Sign(value float64) float64 {
	if math.Signbit(value) {
		return -1.0
	}
	return 1.0
}

// IsNegative returns true if value is negative, including negative zero.
func IsNegative[T ~float64](value T) bool {
	return math.Signbit(float64(value))
}

// IsValid returns true if value is neither NaN nor Inf.
func IsValid[T ~float64](value T) bool {
	return !math.IsNaN(float64(value)) && !math.IsInf(float64(value), 0)
}

// MoveTowards advances current towards target by at most maxDelta. If current
// is already within maxDelta of target, target is returned.
func MoveTowards[T ~float64](current, target, maxDelta T) T {
	diff := target - current
	if diff > 0 {
		return current + Clamp(maxDelta, 0.0, diff)
	} else {
		return current - Clamp(maxDelta, 0.0, -diff)
	}
}
