package dprec

import (
	"fmt"
	"math"
)

// RotationOrder specifies the order in which rotations are applied.
type RotationOrder uint8

const (
	// RotationOrderGlobalXYZ specifies that rotations are applied in the order
	// of X, Y, Z using a global gizmo.
	RotationOrderGlobalXYZ RotationOrder = iota

	// RotationOrderGlobalXZY specifies that rotations are applied in the order
	// of X, Z, Y using a global gizmo.
	RotationOrderGlobalXZY

	// RotationOrderGlobalYXZ specifies that rotations are applied in the order
	// of Y, X, Z using a global gizmo.
	RotationOrderGlobalYXZ

	// RotationOrderGlobalYZX specifies that rotations are applied in the order
	// of Y, Z, X using a global gizmo.
	RotationOrderGlobalYZX

	// RotationOrderGlobalZXY specifies that rotations are applied in the order
	// of Z, X, Y using a global gizmo.
	RotationOrderGlobalZXY

	// RotationOrderGlobalZYX specifies that rotations are applied in the order
	// of Z, Y, X using a global gizmo.
	RotationOrderGlobalZYX
)

const (
	// RotationOrderLocalXYZ specifies that rotations are applied in the order
	// of X, Y, Z using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalXYZ = RotationOrderGlobalZYX

	// RotationOrderLocalXZY specifies that rotations are applied in the order
	// of X, Z, Y using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalXZY = RotationOrderGlobalYZX

	// RotationOrderLocalYXZ specifies that rotations are applied in the order
	// of Y, X, Z using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalYXZ = RotationOrderGlobalZXY

	// RotationOrderLocalYZX specifies that rotations are applied in the order
	// of Y, Z, X using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalYZX = RotationOrderGlobalXZY

	// RotationOrderLocalZXY specifies that rotations are applied in the order
	// of Z, X, Y using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalZXY = RotationOrderGlobalYXZ

	// RotationOrderLocalZYX specifies that rotations are applied in the order
	// of Z, Y, X using a local gizmo (i.e. from the point of view of the object).
	RotationOrderLocalZYX = RotationOrderGlobalXYZ
)

func NewQuat(w, x, y, z float64) Quat {
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

func NegativeQuat(q Quat) Quat {
	return Quat{
		W: -q.W,
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
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

func EulerQuat(x, y, z Angle, order RotationOrder) Quat {
	xRot := RotationQuat(x, BasisXVec3())
	yRot := RotationQuat(y, BasisYVec3())
	zRot := RotationQuat(z, BasisZVec3())
	switch order {
	case RotationOrderGlobalXYZ:
		return QuatProd(QuatProd(zRot, yRot), xRot)
	case RotationOrderGlobalXZY:
		return QuatProd(QuatProd(yRot, zRot), xRot)
	case RotationOrderGlobalYXZ:
		return QuatProd(QuatProd(zRot, xRot), yRot)
	case RotationOrderGlobalYZX:
		return QuatProd(QuatProd(xRot, zRot), yRot)
	case RotationOrderGlobalZXY:
		return QuatProd(QuatProd(yRot, xRot), zRot)
	case RotationOrderGlobalZYX:
		return QuatProd(QuatProd(xRot, yRot), zRot)
	default:
		return IdentityQuat()
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

func QuatScalarProd(q Quat, value float64) Quat {
	return Quat{
		W: q.W * value,
		X: q.X * value,
		Y: q.Y * value,
		Z: q.Z * value,
	}
}

func QuatScalarQuot(q Quat, value float64) Quat {
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

func QuatDot(a, b Quat) float64 {
	return a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func QuatLerp(first, second Quat, t float64) Quat {
	return UnitQuat(Quat{
		W: (1-t)*first.W + t*second.W,
		X: (1-t)*first.X + t*second.X,
		Y: (1-t)*first.Y + t*second.Y,
		Z: (1-t)*first.Z + t*second.Z,
	})
}

func QuatDiff(second, first Quat, shortest bool) Quat {
	if shortest && (QuatDot(second, first) < 0) {
		second = NegativeQuat(second)
	}
	return QuatProd(second, ConjugateQuat(first))
}

func QuatPow(q Quat, pow float64) Quat {
	if q.W > 1.0 {
		return IdentityQuat()
	}
	if q.W < -1.0 {
		return NegativeQuat(IdentityQuat())
	}
	norm := NewVec3(q.X, q.Y, q.Z)
	if norm.IsZero() {
		return IdentityQuat()
	}
	angle := (2 * Acos(q.W)) * Angle(pow)
	return RotationQuat(angle, norm)
}

func QuatSlerp(first, second Quat, t float64) Quat {
	delta := QuatDiff(second, first, true)
	fractDelta := QuatPow(delta, t)
	return UnitQuat(QuatProd(fractDelta, first))
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
	W float64
	X float64
	Y float64
	Z float64
}

func (q Quat) IsNaN() bool {
	return math.IsNaN(q.X) || math.IsNaN(q.Y) || math.IsNaN(q.Z) || math.IsNaN(q.W)
}

func (q Quat) IsInf() bool {
	return math.IsInf(q.X, 0) || math.IsInf(q.Y, 0) || math.IsInf(q.Z, 0) || math.IsInf(q.W, 0)
}

func (q Quat) IsIdentity() bool {
	return Eq(q.X, 0.0) && Eq(q.Y, 0.0) && Eq(q.Z, 0.0) && Eq(q.W, 1.0)
}

func (q Quat) SqrNorm() float64 {
	return q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z
}

func (q Quat) Norm() float64 {
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

// EulerAngles returns the Euler rotation angles for the given quaternion
// and rotation order in which it was presumably created.
//
// The rotations are always returned for X, Y, Z axis in that order.
//
// NOTE: This assumes that the quaternion is normalized.
func (q Quat) EulerAngles(order RotationOrder) (x Angle, y Angle, z Angle) {
	switch order {
	case RotationOrderGlobalXYZ:
		x = Atan2(
			2.0*(q.W*q.X+q.Y*q.Z),
			1.0-2.0*(q.X*q.X+q.Y*q.Y),
		)
		y = Asin(2.0 * (q.W*q.Y - q.X*q.Z))
		z = Atan2(
			2.0*(q.W*q.Z+q.X*q.Y),
			1.0-2.0*(q.Y*q.Y+q.Z*q.Z),
		)
	case RotationOrderGlobalXZY:
		x = Atan2(
			2.0*(q.W*q.X-q.Z*q.Y),
			1.0-2.0*(q.X*q.X+q.Z*q.Z),
		)
		y = Atan2(
			2.0*(q.W*q.Y-q.X*q.Z),
			1.0-2.0*(q.Z*q.Z+q.Y*q.Y),
		)
		z = Asin(2.0 * (q.W*q.Z + q.X*q.Y))
	case RotationOrderGlobalYXZ:
		x = Asin(2.0 * (q.W*q.X + q.Z*q.Y))
		y = Atan2(
			2.0*(q.W*q.Y-q.Z*q.X),
			1.0-2.0*(q.X*q.X+q.Y*q.Y),
		)
		z = Atan2(
			2.0*(q.W*q.Z-q.X*q.Y),
			1.0-2.0*(q.Z*q.Z+q.X*q.X),
		)
	case RotationOrderGlobalYZX:
		x = Atan2(
			2.0*(q.W*q.X+q.Z*q.Y),
			1.0-2.0*(q.X*q.X+q.Z*q.Z),
		)
		y = Atan2(
			2.0*(q.W*q.Y+q.X*q.Z),
			1.0-2.0*(q.Z*q.Z+q.Y*q.Y),
		)
		z = Asin(2.0 * (q.W*q.Z - q.X*q.Y))
	case RotationOrderGlobalZXY:
		x = Asin(2.0 * (q.W*q.X - q.Z*q.Y))
		y = Atan2(
			2.0*(q.W*q.Y+q.Z*q.X),
			1.0-2.0*(q.X*q.X+q.Y*q.Y),
		)
		z = Atan2(
			2.0*(q.W*q.Z+q.X*q.Y),
			1.0-2.0*(q.Z*q.Z+q.X*q.X),
		)
	case RotationOrderGlobalZYX:
		x = Atan2(
			2.0*(q.W*q.X-q.Z*q.Y),
			1.0-2.0*(q.Y*q.Y+q.X*q.X),
		)
		y = Asin(2.0 * (q.W*q.Y + q.Z*q.X))
		z = Atan2(
			2.0*(q.W*q.Z-q.Y*q.X),
			1.0-2.0*(q.Z*q.Z+q.Y*q.Y),
		)
	}
	return
}

func (q Quat) GoString() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", q.W, q.X, q.Y, q.Z)
}
