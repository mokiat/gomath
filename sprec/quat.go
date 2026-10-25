package sprec

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

// NewQuat creates a quaternion with the given W, X, Y, and Z components.
func NewQuat(w, x, y, z float32) Quat {
	return Quat{
		W: w,
		X: x,
		Y: y,
		Z: z,
	}
}

// IdentityQuat returns the identity quaternion (no rotation).
func IdentityQuat() Quat {
	return Quat{
		W: 1.0,
		X: 0.0,
		Y: 0.0,
		Z: 0.0,
	}
}

// NegativeQuat returns the negation of the given quaternion.
func NegativeQuat(q Quat) Quat {
	return Quat{
		W: -q.W,
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
	}
}

// RotationQuat returns a quaternion representing a rotation of angle around
// the given axis direction.
//
// If direction has zero length, the result has NaN components.
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

// EulerQuat returns a quaternion from the given Euler angles and rotation
// order.
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

// ConjugateQuat returns the conjugate of the given quaternion.
func ConjugateQuat(q Quat) Quat {
	return Quat{
		W: q.W,
		X: -q.X,
		Y: -q.Y,
		Z: -q.Z,
	}
}

// QuatScalarProd multiplies a quaternion by a scalar value.
func QuatScalarProd(q Quat, value float32) Quat {
	return Quat{
		W: q.W * value,
		X: q.X * value,
		Y: q.Y * value,
		Z: q.Z * value,
	}
}

// QuatScalarQuot divides a quaternion by a scalar value.
func QuatScalarQuot(q Quat, value float32) Quat {
	invValue := 1.0 / value
	return Quat{
		W: q.W * invValue,
		X: q.X * invValue,
		Y: q.Y * invValue,
		Z: q.Z * invValue,
	}
}

// QuatProd returns the Hamilton product a * b of two quaternions. The
// resulting rotation applies b's rotation first and a's rotation after it.
func QuatProd(a, b Quat) Quat {
	return Quat{
		W: a.W*b.W - a.X*b.X - a.Y*b.Y - a.Z*b.Z,
		X: a.W*b.X + a.X*b.W + a.Y*b.Z - a.Z*b.Y,
		Y: a.W*b.Y - a.X*b.Z + a.Y*b.W + a.Z*b.X,
		Z: a.W*b.Z + a.X*b.Y - a.Y*b.X + a.Z*b.W,
	}
}

// QuatDot returns the dot product of two quaternions.
func QuatDot(a, b Quat) float32 {
	return a.W*b.W + a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// QuatLerp returns the linear interpolation between first and second using t.
// A value of t=0 returns first and t=1 returns second.
//
// The components are interpolated directly, so the result is neither
// normalized nor guaranteed to follow the shorter arc. Use QuatNLerp or
// QuatSlerp to interpolate between rotations.
func QuatLerp(first, second Quat, t float32) Quat {
	return Quat{
		W: first.W + t*(second.W-first.W),
		X: first.X + t*(second.X-first.X),
		Y: first.Y + t*(second.Y-first.Y),
		Z: first.Z + t*(second.Z-first.Z),
	}
}

// QuatNLerp returns the normalized linear interpolation between first and
// second using t. The interpolation always follows the shorter arc between
// the two orientations.
func QuatNLerp(first, second Quat, t float32) Quat {
	if QuatDot(first, second) < 0.0 {
		second = NegativeQuat(second)
	}
	return UnitQuat(QuatLerp(first, second, t))
}

// QuatDiff returns the rotation difference a * b^-1, i.e. the rotation that,
// when applied after b, results in a.
//
// NOTE: This assumes that the quaternions are normalized.
func QuatDiff(a, b Quat) Quat {
	return QuatProd(a, ConjugateQuat(b))
}

// QuatPow raises the quaternion to the given power.
func QuatPow(q Quat, pow float32) Quat {
	axis := NewVec3(q.X, q.Y, q.Z)
	sn := axis.Length()
	if sn == 0.0 {
		return IdentityQuat()
	}
	// Atan2 is used instead of Acos(q.W), since the latter loses precision
	// for small angles, where q.W is close to 1.0.
	angle := (2 * Atan2(sn, q.W)) * Angle(pow)
	return RotationQuat(angle, axis)
}

// QuatSlerp returns the spherical linear interpolation between first and
// second using t. The interpolation always follows the shorter arc between
// the two orientations.
func QuatSlerp(first, second Quat, t float32) Quat {
	if QuatDot(first, second) < 0.0 {
		second = NegativeQuat(second)
	}
	delta := QuatDiff(second, first)
	fractDelta := QuatPow(delta, t)
	return QuatProd(fractDelta, first)
}

// QuatVec3Rotation rotates the vector v by the quaternion q.
func QuatVec3Rotation(q Quat, v Vec3) Vec3 {
	// Uses the direct formula: v + 2.0 * cross(q.xyz, cross(q.xyz, v) + q.w * v)
	// See https://en.wikipedia.org/wiki/Quaternions_and_spatial_rotation#Used_methods
	qAxis := NewVec3(q.X, q.Y, q.Z)
	shared := Vec3Prod(Vec3Cross(qAxis, v), 2.0)
	return Vec3Sum(
		Vec3Sum(v, Vec3Prod(shared, q.W)),
		Vec3Cross(qAxis, shared),
	)
}

// UnitQuat returns the normalized quaternion.
//
// If q has zero norm, the result has NaN components.
func UnitQuat(q Quat) Quat {
	return QuatScalarQuot(q, q.Norm())
}

// InverseQuat returns the multiplicative inverse of the given quaternion.
//
// If q has zero norm, the result has NaN components.
func InverseQuat(q Quat) Quat {
	return QuatScalarQuot(ConjugateQuat(q), q.SqrNorm())
}

// Quat is a quaternion with float32 components.
type Quat struct {
	W float32
	X float32
	Y float32
	Z float32
}

// IsNaN returns true if any component is NaN.
func (q Quat) IsNaN() bool {
	return math.IsNaN(float64(q.X)) || math.IsNaN(float64(q.Y)) || math.IsNaN(float64(q.Z)) || math.IsNaN(float64(q.W))
}

// IsInf returns true if any component is Inf.
func (q Quat) IsInf() bool {
	return math.IsInf(float64(q.X), 0) || math.IsInf(float64(q.Y), 0) || math.IsInf(float64(q.Z), 0) || math.IsInf(float64(q.W), 0)
}

// IsIdentity returns true if this is the identity quaternion.
func (q Quat) IsIdentity() bool {
	return Eq(q.X, 0.0) && Eq(q.Y, 0.0) && Eq(q.Z, 0.0) && Eq(q.W, 1.0)
}

// SqrNorm returns the squared norm of the quaternion.
func (q Quat) SqrNorm() float32 {
	return q.W*q.W + q.X*q.X + q.Y*q.Y + q.Z*q.Z
}

// Norm returns the norm (magnitude) of the quaternion.
func (q Quat) Norm() float32 {
	return Sqrt(q.SqrNorm())
}

// OrientationX returns the X basis vector of the rotation.
func (q Quat) OrientationX() Vec3 {
	return Vec3{
		X: 1.0 - 2.0*(q.Y*q.Y+q.Z*q.Z),
		Y: 2.0 * (q.X*q.Y + q.W*q.Z),
		Z: 2.0 * (q.X*q.Z - q.W*q.Y),
	}
}

// OrientationY returns the Y basis vector of the rotation.
func (q Quat) OrientationY() Vec3 {
	return Vec3{
		X: 2.0 * (q.X*q.Y - q.W*q.Z),
		Y: 1.0 - 2.0*(q.X*q.X+q.Z*q.Z),
		Z: 2.0 * (q.Y*q.Z + q.W*q.X),
	}
}

// OrientationZ returns the Z basis vector of the rotation.
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
		y = Asin(Clamp(2.0*(q.W*q.Y-q.X*q.Z), -1.0, 1.0))
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
		z = Asin(Clamp(2.0*(q.W*q.Z+q.X*q.Y), -1.0, 1.0))
	case RotationOrderGlobalYXZ:
		x = Asin(Clamp(2.0*(q.W*q.X+q.Z*q.Y), -1.0, 1.0))
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
		z = Asin(Clamp(2.0*(q.W*q.Z-q.X*q.Y), -1.0, 1.0))
	case RotationOrderGlobalZXY:
		x = Asin(Clamp(2.0*(q.W*q.X-q.Z*q.Y), -1.0, 1.0))
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
		y = Asin(Clamp(2.0*(q.W*q.Y+q.Z*q.X), -1.0, 1.0))
		z = Atan2(
			2.0*(q.W*q.Z-q.Y*q.X),
			1.0-2.0*(q.Z*q.Z+q.Y*q.Y),
		)
	}
	return
}

// String returns a string representation of the quaternion.
func (q Quat) String() string {
	return fmt.Sprintf("(%f, %f, %f, %f)", q.W, q.X, q.Y, q.Z)
}
