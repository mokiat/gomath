package dprec

import "math"

const (
	Pi      = float64(math.Pi)
	Epsilon = float64(0.000000000001)
)

func Abs[T ~float64](value T) T {
	return T(math.Abs(float64(value)))
}

func Max[T ~float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T ~float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Clamp[T ~float64](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
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

func Sqrt(value float64) float64 {
	return math.Sqrt(value)
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

func Sign(value float64) float64 {
	if math.Signbit(value) {
		return -1.0
	}
	return 1.0
}
