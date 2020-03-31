package sprec

func Radians(radians float32) Angle {
	return Angle(radians)
}

func Degrees(degrees float32) Angle {
	return Angle(Pi * (degrees / 180.0))
}

type Angle float32

func (a Angle) Degrees() float32 {
	return 180.0 * (float32(a) / Pi)
}

func (a Angle) Radians() float32 {
	return float32(a)
}
