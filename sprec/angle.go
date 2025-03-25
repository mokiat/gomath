package sprec

import "math"

func Radians(radians float32) Angle {
	return Angle(radians)
}

func Degrees(degrees float32) Angle {
	return Angle(Pi * (degrees / 180.0))
}

func NormalizeAngle(a Angle) Angle {
	radians := Mod(float32(a), Tau)
	if radians < 0.0 {
		radians += Tau
	}
	return Angle(radians)
}

type Angle float32

func (a Angle) IsNaN() bool {
	return math.IsNaN(float64(a))
}

func (a Angle) IsInf() bool {
	return math.IsInf(float64(a), 0)
}

func (a Angle) Degrees() float32 {
	return 180.0 * (float32(a) / Pi)
}

func (a Angle) Radians() float32 {
	return float32(a)
}
