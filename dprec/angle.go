package dprec

func Radians(radians float64) Angle {
	return Angle(radians)
}

func Degrees(degrees float64) Angle {
	return Angle(Pi * (degrees / 180.0))
}

type Angle float64

func (a Angle) Degrees() float64 {
	return 180.0 * (float64(a) / Pi)
}

func (a Angle) Radians() float64 {
	return float64(a)
}
