package dprec

import "math"

func Radians(radians float64) Angle {
	return Angle(radians)
}

func Degrees(degrees float64) Angle {
	return Angle(Pi * (degrees / 180.0))
}

func NormalizeAngle(a Angle) Angle {
	radians := Mod(float64(a), Tau)
	if radians < 0.0 {
		radians += Tau
	}
	return Angle(radians)
}

type Angle float64

func (a Angle) IsNaN() bool {
	return math.IsNaN(float64(a))
}

func (a Angle) IsInf() bool {
	return math.IsInf(float64(a), 0)
}

func (a Angle) Degrees() float64 {
	return 180.0 * (float64(a) / Pi)
}

func (a Angle) Radians() float64 {
	return float64(a)
}
