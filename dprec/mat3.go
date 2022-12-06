package dprec

import (
	"fmt"
	"math"
)

func NewMat3(
	m11, m12, m13 float64,
	m21, m22, m23 float64,
	m31, m32, m33 float64,
) Mat3 {
	return Mat3{
		M11: m11, M12: m12, M13: m13,
		M21: m21, M22: m22, M23: m23,
		M31: m31, M32: m32, M33: m33,
	}
}

func ZeroMat3() Mat3 {
	return Mat3{}
}

func IdentityMat3() Mat3 {
	var result Mat3
	result.M11 = 1.0
	result.M22 = 1.0
	result.M33 = 1.0
	return result
}

func TransposedMat3(m Mat3) Mat3 {
	return NewMat3(
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
	)
}

func TranslationMat3(x, y float64) Mat3 {
	result := IdentityMat3()
	result.M13 = x
	result.M23 = y
	return result
}

func ScaleMat3(x, y float64) Mat3 {
	var result Mat3
	result.M11 = x
	result.M22 = y
	result.M33 = 1.0
	return result
}

func RotationMat3(angle Angle) Mat3 {
	cs := Cos(angle)
	sn := Sin(angle)

	var result Mat3
	result.M11 = cs
	result.M12 = -sn
	result.M21 = sn
	result.M22 = cs
	result.M33 = 1.0
	return result
}

func OrthoMat3(left, right, top, bottom float64) Mat3 {
	var result Mat3
	result.M11 = 2.0 / (right - left)
	result.M13 = (right + left) / (left - right)

	result.M22 = 2.0 / (top - bottom)
	result.M23 = (top + bottom) / (bottom - top)

	result.M33 = 1.0
	return result
}

// FastInverseMat3 calculates the inverse of the matrix with a few caveats.
//
// The matrix should be a transformation one that was constructed through the multiplication
// of one or more of the following transformations: identity, translation, rotation.
//
// For all other scenarios (e.g. a scale transformation was used), the InverseMat3 method should be
// used instead, though it will be slower.
func FastInverseMat3(m Mat3) Mat3 {
	inverseTranslate := TranslationMat3(-m.M13, -m.M23)
	inverseRotate := NewMat3(
		m.M11, m.M21, 0.0,
		m.M12, m.M22, 0.0,
		0.0, 0.0, 1.0,
	)
	return Mat3Prod(inverseRotate, inverseTranslate)
}

// InverseMat3 calculates the inverse of the matrix.
//
// The behavior is undefined if the matrix is not reversible
// (i.e. has a zero determinant).
func InverseMat3(m Mat3) Mat3 {
	minor11 := m.M22*m.M33 - m.M23*m.M32
	minor12 := m.M21*m.M33 - m.M23*m.M31
	minor13 := m.M21*m.M32 - m.M22*m.M31

	minor21 := m.M12*m.M33 - m.M13*m.M32
	minor22 := m.M11*m.M33 - m.M13*m.M31
	minor23 := m.M11*m.M32 - m.M12*m.M31

	minor31 := m.M12*m.M23 - m.M13*m.M22
	minor32 := m.M11*m.M23 - m.M13*m.M21
	minor33 := m.M11*m.M22 - m.M12*m.M21

	determinant := m.M11*minor11 - m.M12*minor12 + m.M13*minor13

	return NewMat3(
		+minor11/determinant, -minor21/determinant, +minor31/determinant,
		-minor12/determinant, +minor22/determinant, -minor32/determinant,
		+minor13/determinant, -minor23/determinant, +minor33/determinant,
	)
}

func TransformationMat3(orientX, orientY, translation Vec2) Mat3 {
	var result Mat3
	result.M11 = orientX.X
	result.M12 = orientY.X
	result.M13 = translation.X

	result.M21 = orientX.Y
	result.M22 = orientY.Y
	result.M23 = translation.Y

	result.M33 = 1.0
	return result
}

func RowMajorArrayToMat3(values [9]float64) Mat3 {
	return Mat3{
		M11: values[0], M12: values[1], M13: values[2],
		M21: values[3], M22: values[4], M23: values[5],
		M31: values[6], M32: values[7], M33: values[8],
	}
}

func ColumnMajorArrayToMat3(values [9]float64) Mat3 {
	return Mat3{
		M11: values[0], M12: values[3], M13: values[6],
		M21: values[1], M22: values[4], M23: values[7],
		M31: values[2], M32: values[5], M33: values[8],
	}
}

func Mat3Prod(left, right Mat3) Mat3 {
	return Mat3{
		M11: left.M11*right.M11 + left.M12*right.M21 + left.M13*right.M31,
		M12: left.M11*right.M12 + left.M12*right.M22 + left.M13*right.M32,
		M13: left.M11*right.M13 + left.M12*right.M23 + left.M13*right.M33,

		M21: left.M21*right.M11 + left.M22*right.M21 + left.M23*right.M31,
		M22: left.M21*right.M12 + left.M22*right.M22 + left.M23*right.M32,
		M23: left.M21*right.M13 + left.M22*right.M23 + left.M23*right.M33,

		M31: left.M31*right.M11 + left.M32*right.M21 + left.M33*right.M31,
		M32: left.M31*right.M12 + left.M32*right.M22 + left.M33*right.M32,
		M33: left.M31*right.M13 + left.M32*right.M23 + left.M33*right.M33,
	}
}

func Mat3MultiProd(first Mat3, others ...Mat3) Mat3 {
	result := first
	for _, matrix := range others {
		result = Mat3Prod(result, matrix)
	}
	return result
}

func Mat3Vec3Prod(mat Mat3, vec Vec3) Vec3 {
	return Vec3{
		X: mat.M11*vec.X + mat.M12*vec.Y + mat.M13*vec.Z,
		Y: mat.M21*vec.X + mat.M22*vec.Y + mat.M23*vec.Z,
		Z: mat.M31*vec.X + mat.M32*vec.Y + mat.M33*vec.Z,
	}
}

type Mat3 struct {
	M11, M12, M13 float64
	M21, M22, M23 float64
	M31, M32, M33 float64
}

func (m Mat3) IsNaN() bool {
	return math.IsNaN(m.M11) || math.IsNaN(m.M12) || math.IsNaN(m.M13) ||
		math.IsNaN(m.M21) || math.IsNaN(m.M22) || math.IsNaN(m.M23) ||
		math.IsNaN(m.M31) || math.IsNaN(m.M32) || math.IsNaN(m.M33)
}

func (m Mat3) IsInf() bool {
	return math.IsInf(m.M11, 0) || math.IsInf(m.M12, 0) || math.IsInf(m.M13, 0) ||
		math.IsInf(m.M21, 0) || math.IsInf(m.M22, 0) || math.IsInf(m.M23, 0) ||
		math.IsInf(m.M31, 0) || math.IsInf(m.M32, 0) || math.IsInf(m.M33, 0)
}

func (m Mat3) Row1() Vec3 {
	return NewVec3(m.M11, m.M12, m.M13)
}

func (m Mat3) Row2() Vec3 {
	return NewVec3(m.M21, m.M22, m.M23)
}

func (m Mat3) Row3() Vec3 {
	return NewVec3(m.M31, m.M32, m.M33)
}

func (m Mat3) Column1() Vec3 {
	return NewVec3(m.M11, m.M21, m.M31)
}

func (m Mat3) Column2() Vec3 {
	return NewVec3(m.M12, m.M22, m.M32)
}

func (m Mat3) Column3() Vec3 {
	return NewVec3(m.M13, m.M23, m.M33)
}

func (m Mat3) OrientationX() Vec2 {
	return NewVec2(m.M11, m.M21)
}

func (m Mat3) OrientationY() Vec2 {
	return NewVec2(m.M12, m.M22)
}

func (m Mat3) Translation() Vec2 {
	return NewVec2(m.M13, m.M23)
}

func (m Mat3) RowMajorArray() [9]float64 {
	return [9]float64{
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	}
}

func (m Mat3) ColumnMajorArray() [9]float64 {
	return [9]float64{
		m.M11, m.M21, m.M31,
		m.M12, m.M22, m.M32,
		m.M13, m.M23, m.M33,
	}
}

func (m Mat3) GoString() string {
	return fmt.Sprintf("((%f, %f, %f), (%f, %f, %f), (%f, %f, %f))",
		m.M11, m.M12, m.M13,
		m.M21, m.M22, m.M23,
		m.M31, m.M32, m.M33,
	)
}
