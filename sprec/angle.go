package sprec

import "math"

// Radians creates an angle from radians.
func Radians(radians float32) Angle {
	return Angle(radians)
}

// Degrees creates an angle from degrees.
func Degrees(degrees float32) Angle {
	return Angle(Pi * (degrees / 180.0))
}

// NormalizeAngle normalizes an angle to the range [-Pi..Pi].
func NormalizeAngle(a Angle) Angle {
	radians := Mod(float32(a), Tau)
	switch {
	case radians < -Pi:
		radians += Tau
	case radians > Pi:
		radians -= Tau
	}
	return Angle(radians)
}

// NormalizeAnglePos normalizes an angle to the range [0.0..Tau].
func NormalizeAnglePos(a Angle) Angle {
	radians := Mod(float32(a), Tau)
	if radians < 0.0 {
		radians += Tau
	}
	return Angle(radians)
}

// NormalizeAngleNeg normalizes an angle to the range [-Tau..0.0].
func NormalizeAngleNeg(a Angle) Angle {
	radians := Mod(float32(a), Tau)
	if radians > 0.0 {
		radians -= Tau
	}
	return Angle(radians)
}

// Angle represents an angle.
type Angle float32

// IsNaN returns true if the angle is a NaN value.
func (a Angle) IsNaN() bool {
	return math.IsNaN(float64(a))
}

// IsInf returns true if the angle is an Inf value.
func (a Angle) IsInf() bool {
	return math.IsInf(float64(a), 0)
}

// Degrees returns the angle in degrees.
func (a Angle) Degrees() float32 {
	return 180.0 * (float32(a) / Pi)
}

// Radians returns the angle in radians.
func (a Angle) Radians() float32 {
	return float32(a)
}
