package sprec

func NewVec3(x, y, z float32) Vec3 {
	return Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}
