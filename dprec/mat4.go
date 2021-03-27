package dprec

import "fmt"

func NewMat4(
	m11, m12, m13, m14 float64,
	m21, m22, m23, m24 float64,
	m31, m32, m33, m34 float64,
	m41, m42, m43, m44 float64,
) Mat4 {
	return Mat4{
		M11: m11, M12: m12, M13: m13, M14: m14,
		M21: m21, M22: m22, M23: m23, M24: m24,
		M31: m31, M32: m32, M33: m33, M34: m34,
		M41: m41, M42: m42, M43: m43, M44: m44,
	}
}

func ZeroMat4() Mat4 {
	return Mat4{}
}

func IdentityMat4() Mat4 {
	var result Mat4
	result.M11 = 1.0
	result.M22 = 1.0
	result.M33 = 1.0
	result.M44 = 1.0
	return result
}

func TranslationMat4(x, y, z float64) Mat4 {
	result := IdentityMat4()
	result.M14 = x
	result.M24 = y
	result.M34 = z
	return result
}

func ScaleMat4(x, y, z float64) Mat4 {
	var result Mat4
	result.M11 = x
	result.M22 = y
	result.M33 = z
	result.M44 = 1.0
	return result
}

func RotationMat4(angle Angle, x, y, z float64) Mat4 {
	vector := UnitVec3(NewVec3(x, y, z))
	return rotationMat4FromNormalizedData(Cos(angle), Sin(angle), vector)
}

func rotationMat4FromNormalizedData(cs, sn float64, vector Vec3) Mat4 {
	x, y, z := vector.X, vector.Y, vector.Z

	var result Mat4
	result.M11 = x*x*(1.0-cs) + cs
	result.M21 = x*y*(1.0-cs) + z*sn
	result.M31 = x*z*(1.0-cs) - y*sn

	result.M12 = y*x*(1.0-cs) - z*sn
	result.M22 = y*y*(1.0-cs) + cs
	result.M32 = y*z*(1.0-cs) + x*sn

	result.M13 = z*x*(1.0-cs) + y*sn
	result.M23 = z*y*(1.0-cs) - x*sn
	result.M33 = z*z*(1.0-cs) + cs

	result.M44 = 1.0
	return result
}

func OrthoMat4(left, right, top, bottom, near, far float64) Mat4 {
	var result Mat4
	result.M11 = 2.0 / (right - left)
	result.M14 = (right + left) / (left - right)

	result.M22 = 2.0 / (top - bottom)
	result.M24 = (top + bottom) / (bottom - top)

	result.M33 = 2.0 / (near - far)
	result.M34 = (far + near) / (near - far)

	result.M44 = 1.0
	return result
}

func PerspectiveMat4(left, right, bottom, top, near, far float64) Mat4 {
	var result Mat4
	result.M11 = 2.0 * near / (right - left)
	result.M13 = (right + left) / (right - left)

	result.M22 = 2.0 * near / (top - bottom)
	result.M23 = (top + bottom) / (top - bottom)

	result.M33 = (far + near) / (near - far)
	result.M34 = 2.0 * far * near / (near - far)

	result.M43 = -1.0
	return result
}

// FastInverseMat4 calculates the inverse of the matrix with a few caveats.
//
// The matrix should be a transformation one that was constructed through the multiplication
// of one or more of the following transformations: identity, translation, rotation.
//
// For all other scenarios (e.g. a scale transformation was used), the InverseMat4 method should be
// used instead, though it will be slower.
func FastInverseMat4(m Mat4) Mat4 {
	inverseTranslate := TranslationMat4(
		-m.M14, -m.M24, -m.M34,
	)
	inverseRotate := NewMat4(
		m.M11, m.M21, m.M31, 0.0,
		m.M12, m.M22, m.M32, 0.0,
		m.M13, m.M23, m.M33, 0.0,
		0.0, 0.0, 0.0, 1.0,
	)
	return Mat4Prod(inverseRotate, inverseTranslate)
}

// InverseMat4 calculates the inverse of the matrix.
//
// The behavior is undefined if the matrix is not reversible
// (i.e. has a zero determinant).
func InverseMat4(m Mat4) Mat4 {
	minor11 := m.M22*m.M33*m.M44 + m.M23*m.M34*m.M42 + m.M24*m.M32*m.M43 - m.M24*m.M33*m.M42 - m.M23*m.M32*m.M44 - m.M22*m.M34*m.M43
	minor12 := m.M21*m.M33*m.M44 + m.M23*m.M34*m.M41 + m.M24*m.M31*m.M43 - m.M24*m.M33*m.M41 - m.M23*m.M31*m.M44 - m.M21*m.M34*m.M43
	minor13 := m.M21*m.M32*m.M44 + m.M22*m.M34*m.M41 + m.M24*m.M31*m.M42 - m.M24*m.M32*m.M41 - m.M22*m.M31*m.M44 - m.M21*m.M34*m.M42
	minor14 := m.M21*m.M32*m.M43 + m.M22*m.M33*m.M41 + m.M23*m.M31*m.M42 - m.M23*m.M32*m.M41 - m.M22*m.M31*m.M43 - m.M21*m.M33*m.M42
	minor21 := m.M12*m.M33*m.M44 + m.M13*m.M34*m.M42 + m.M14*m.M32*m.M43 - m.M14*m.M33*m.M42 - m.M13*m.M32*m.M44 - m.M12*m.M34*m.M43
	minor22 := m.M11*m.M33*m.M44 + m.M13*m.M34*m.M41 + m.M14*m.M31*m.M43 - m.M14*m.M33*m.M41 - m.M13*m.M31*m.M44 - m.M11*m.M34*m.M43
	minor23 := m.M11*m.M32*m.M44 + m.M12*m.M34*m.M41 + m.M14*m.M31*m.M42 - m.M14*m.M32*m.M41 - m.M12*m.M31*m.M44 - m.M11*m.M34*m.M42
	minor24 := m.M11*m.M32*m.M43 + m.M12*m.M33*m.M41 + m.M13*m.M31*m.M42 - m.M13*m.M32*m.M41 - m.M12*m.M31*m.M43 - m.M11*m.M33*m.M42
	minor31 := m.M12*m.M23*m.M44 + m.M13*m.M24*m.M42 + m.M14*m.M22*m.M43 - m.M14*m.M23*m.M42 - m.M13*m.M22*m.M44 - m.M12*m.M24*m.M43
	minor32 := m.M11*m.M23*m.M44 + m.M13*m.M24*m.M41 + m.M14*m.M21*m.M43 - m.M14*m.M23*m.M41 - m.M13*m.M21*m.M44 - m.M11*m.M24*m.M43
	minor33 := m.M11*m.M22*m.M44 + m.M12*m.M24*m.M41 + m.M14*m.M21*m.M42 - m.M14*m.M22*m.M41 - m.M12*m.M21*m.M44 - m.M11*m.M24*m.M42
	minor34 := m.M11*m.M22*m.M43 + m.M12*m.M23*m.M41 + m.M13*m.M21*m.M42 - m.M13*m.M22*m.M41 - m.M12*m.M21*m.M43 - m.M11*m.M23*m.M42
	minor41 := m.M12*m.M23*m.M34 + m.M13*m.M24*m.M32 + m.M14*m.M22*m.M33 - m.M14*m.M23*m.M32 - m.M13*m.M22*m.M34 - m.M12*m.M24*m.M33
	minor42 := m.M11*m.M23*m.M34 + m.M13*m.M24*m.M31 + m.M14*m.M21*m.M33 - m.M14*m.M23*m.M31 - m.M13*m.M21*m.M34 - m.M11*m.M24*m.M33
	minor43 := m.M11*m.M22*m.M34 + m.M12*m.M24*m.M31 + m.M14*m.M21*m.M32 - m.M14*m.M22*m.M31 - m.M12*m.M21*m.M34 - m.M11*m.M24*m.M32
	minor44 := m.M11*m.M22*m.M33 + m.M12*m.M23*m.M31 + m.M13*m.M21*m.M32 - m.M13*m.M22*m.M31 - m.M12*m.M21*m.M33 - m.M11*m.M23*m.M32

	determinant := m.M11*minor11 - m.M12*minor12 + m.M13*minor13 - m.M14*minor14

	return NewMat4(
		+minor11/determinant, -minor21/determinant, +minor31/determinant, -minor41/determinant,
		-minor12/determinant, +minor22/determinant, -minor32/determinant, +minor42/determinant,
		+minor13/determinant, -minor23/determinant, +minor33/determinant, -minor43/determinant,
		-minor14/determinant, +minor24/determinant, -minor34/determinant, +minor44/determinant,
	)
}

func TransformationMat4(orientX, orientY, orientZ, translation Vec3) Mat4 {
	var result Mat4
	result.M11 = orientX.X
	result.M12 = orientY.X
	result.M13 = orientZ.X
	result.M14 = translation.X

	result.M21 = orientX.Y
	result.M22 = orientY.Y
	result.M23 = orientZ.Y
	result.M24 = translation.Y

	result.M31 = orientX.Z
	result.M32 = orientY.Z
	result.M33 = orientZ.Z
	result.M34 = translation.Z

	result.M44 = 1.0
	return result
}

func OrientationMat4(orientX, orientY, orientZ Vec3) Mat4 {
	var result Mat4
	result.M11 = orientX.X
	result.M12 = orientY.X
	result.M13 = orientZ.X

	result.M21 = orientX.Y
	result.M22 = orientY.Y
	result.M23 = orientZ.Y

	result.M31 = orientX.Z
	result.M32 = orientY.Z
	result.M33 = orientZ.Z

	result.M44 = 1.0
	return result
}

func RowMajorArrayMat4(values [16]float64) Mat4 {
	return Mat4{
		M11: values[0], M12: values[1], M13: values[2], M14: values[3],
		M21: values[4], M22: values[5], M23: values[6], M24: values[7],
		M31: values[8], M32: values[9], M33: values[10], M34: values[11],
		M41: values[12], M42: values[13], M43: values[14], M44: values[15],
	}
}

func ColumnMajorArrayMat4(values [16]float64) Mat4 {
	return Mat4{
		M11: values[0], M12: values[4], M13: values[8], M14: values[12],
		M21: values[1], M22: values[5], M23: values[9], M24: values[13],
		M31: values[2], M32: values[6], M33: values[10], M34: values[14],
		M41: values[3], M42: values[7], M43: values[11], M44: values[15],
	}
}

func Mat4Prod(left, right Mat4) Mat4 {
	return Mat4{
		M11: left.M11*right.M11 + left.M12*right.M21 + left.M13*right.M31 + left.M14*right.M41,
		M12: left.M11*right.M12 + left.M12*right.M22 + left.M13*right.M32 + left.M14*right.M42,
		M13: left.M11*right.M13 + left.M12*right.M23 + left.M13*right.M33 + left.M14*right.M43,
		M14: left.M11*right.M14 + left.M12*right.M24 + left.M13*right.M34 + left.M14*right.M44,

		M21: left.M21*right.M11 + left.M22*right.M21 + left.M23*right.M31 + left.M24*right.M41,
		M22: left.M21*right.M12 + left.M22*right.M22 + left.M23*right.M32 + left.M24*right.M42,
		M23: left.M21*right.M13 + left.M22*right.M23 + left.M23*right.M33 + left.M24*right.M43,
		M24: left.M21*right.M14 + left.M22*right.M24 + left.M23*right.M34 + left.M24*right.M44,

		M31: left.M31*right.M11 + left.M32*right.M21 + left.M33*right.M31 + left.M34*right.M41,
		M32: left.M31*right.M12 + left.M32*right.M22 + left.M33*right.M32 + left.M34*right.M42,
		M33: left.M31*right.M13 + left.M32*right.M23 + left.M33*right.M33 + left.M34*right.M43,
		M34: left.M31*right.M14 + left.M32*right.M24 + left.M33*right.M34 + left.M34*right.M44,

		M41: left.M41*right.M11 + left.M42*right.M21 + left.M43*right.M31 + left.M44*right.M41,
		M42: left.M41*right.M12 + left.M42*right.M22 + left.M43*right.M32 + left.M44*right.M42,
		M43: left.M41*right.M13 + left.M42*right.M23 + left.M43*right.M33 + left.M44*right.M43,
		M44: left.M41*right.M14 + left.M42*right.M24 + left.M43*right.M34 + left.M44*right.M44,
	}
}

func Mat4MultiProd(first Mat4, others ...Mat4) Mat4 {
	result := first
	for _, matrix := range others {
		result = Mat4Prod(result, matrix)
	}
	return result
}

func Mat4Vec4Prod(mat Mat4, vec Vec4) Vec4 {
	return Vec4{
		X: mat.M11*vec.X + mat.M12*vec.Y + mat.M13*vec.Z + mat.M14*vec.W,
		Y: mat.M21*vec.X + mat.M22*vec.Y + mat.M23*vec.Z + mat.M24*vec.W,
		Z: mat.M31*vec.X + mat.M32*vec.Y + mat.M33*vec.Z + mat.M34*vec.W,
		W: mat.M41*vec.X + mat.M42*vec.Y + mat.M43*vec.Z + mat.M44*vec.W,
	}
}

func Mat4Vec3Transformation(mat Mat4, vec Vec3) Vec3 {
	return Vec3{
		X: mat.M11*vec.X + mat.M12*vec.Y + mat.M13*vec.Z + mat.M14,
		Y: mat.M21*vec.X + mat.M22*vec.Y + mat.M23*vec.Z + mat.M24,
		Z: mat.M31*vec.X + mat.M32*vec.Y + mat.M33*vec.Z + mat.M34,
	}
}

type Mat4 struct {
	M11, M12, M13, M14 float64
	M21, M22, M23, M24 float64
	M31, M32, M33, M34 float64
	M41, M42, M43, M44 float64
}

func (m Mat4) OrientationX() Vec3 {
	return NewVec3(m.M11, m.M21, m.M31)
}

func (m Mat4) OrientationY() Vec3 {
	return NewVec3(m.M12, m.M22, m.M32)
}

func (m Mat4) OrientationZ() Vec3 {
	return NewVec3(m.M13, m.M23, m.M33)
}

func (m Mat4) Translation() Vec3 {
	return NewVec3(m.M14, m.M24, m.M34)
}

func (m Mat4) RowMajorArray() [16]float64 {
	return [16]float64{
		m.M11, m.M12, m.M13, m.M14,
		m.M21, m.M22, m.M23, m.M24,
		m.M31, m.M32, m.M33, m.M34,
		m.M41, m.M42, m.M43, m.M44,
	}
}

func (m Mat4) ColumnMajorArray() [16]float64 {
	return [16]float64{
		m.M11, m.M21, m.M31, m.M41,
		m.M12, m.M22, m.M32, m.M42,
		m.M13, m.M23, m.M33, m.M43,
		m.M14, m.M24, m.M34, m.M44,
	}
}

func (m Mat4) GoString() string {
	return fmt.Sprintf("((%f, %f, %f, %f), (%f, %f, %f, %f), (%f, %f, %f, %f), (%f, %f, %f, %f))",
		m.M11, m.M12, m.M13, m.M14,
		m.M21, m.M22, m.M23, m.M24,
		m.M31, m.M32, m.M33, m.M34,
		m.M41, m.M42, m.M43, m.M44,
	)
}
