package sprec

import "fmt"

func NewQuat(w, x, y, z float32) Quat {
	return Quat{
		W: w,
		X: x,
		Y: y,
		Z: z,
	}
}

func IdentityQuat() Quat {
	return Quat{
		W: 1.0,
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}
}

func RotationQuat(angle Angle, direction Vec3) Quat {
	cs := Cos(angle / 2.0)
	sn := Sin(angle / 2.0)
	normalizedDirection := UnitVec3(direction)
	return Quat{
		W: cs,
		X: sn * normalizedDirection.X,
		Y: sn * normalizedDirection.Y,
		Z: sn * normalizedDirection.Z,
	}
}

func ConjugateQuat(q Quat) Quat {
	return Quat{
		W: q.W,
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
	}
}

func QuatScalarProd(q Quat, value float32) Quat {
	return Quat{
		W: q.W * value,
		X: q.X * value,
		Y: q.Y * value,
		Z: q.Z * value,
	}
}

func QuatScalarQuot(q Quat, value float32) Quat {
	return Quat{
		W: q.W / value,
		X: q.X / value,
		Y: q.Y / value,
		Z: q.Z / value,
	}
}

func QuatProd(first, second Quat) Quat {
	return Quat{
		W: first.W*second.W - first.X*second.X - first.Y*second.Y - first.Z*second.Z,
		X: first.W*second.X + first.X*second.W + first.Y*second.Z - first.Z*second.Y,
		Y: first.W*second.Y - first.X*second.Z + first.Y*second.W + first.Z*second.X,
		Z: first.W*second.Z + first.X*second.Y - first.Y*second.X + first.Z*second.W,
	}
}

func QuatVec3Rotation(q Quat, v Vec3) Vec3 {
	vectorQuat := Quat{
		W: 0.0,
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
	res := QuatProd(QuatProd(q, vectorQuat), ConjugateQuat(q))
	return Vec3{
		X: res.X,
		Y: res.Y,
		Z: res.Z,
	}
}

func UnitQuat(q Quat) Quat {
	return QuatScalarQuot(q, q.Norm())
}

func InverseQuat(q Quat) Quat {
	return QuatScalarQuot(ConjugateQuat(q), q.SqrNorm())
}

type Quat struct {
	W float32
	X float32
	Y float32
	Z float32
}

func (q Quat) SqrNorm() float32 {
	return q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z
}

func (q Quat) Norm() float32 {
	return Sqrt(q.SqrNorm())
}

func (q Quat) OrientationX() Vec3 {
	return Vec3{
		X: 1.0 - 2.0*(q.Y*q.Y+q.Z*q.Z),
		Y: 2.0 * (q.X*q.Y + q.W*q.Z),
		Z: 2.0 * (q.X*q.Z - q.W*q.Y),
	}
}

func (q Quat) OrientationY() Vec3 {
	return Vec3{
		X: 2.0 * (q.X*q.Y - q.W*q.Z),
		Y: 1.0 - 2.0*(q.X*q.X+q.Z*q.Z),
		Z: 2.0 * (q.Y*q.Z + q.W*q.X),
	}
}

func (q Quat) OrientationZ() Vec3 {
	return Vec3{
		X: 2.0 * (q.X*q.Z + q.W*q.Y),
		Y: 2.0 * (q.Y*q.Z - q.W*q.X),
		Z: 1.0 - 2.0*(q.X*q.X+q.Y*q.Y),
	}
}

func (q Quat) GoString() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", q.W, q.X, q.Y, q.Z)
}
