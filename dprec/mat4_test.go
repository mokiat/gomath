package dprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/dprec"
	. "github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Mat4", func() {
	var matrix Mat4
	var otherMatrix Mat4
	var vector Vec4

	BeforeEach(func() {
		matrix = NewMat4(
			0.1, 0.2, 0.3, 0.4,
			0.5, 0.6, 0.7, 0.8,
			0.9, 1.0, 1.1, 1.2,
			1.3, 1.4, 1.5, 1.6,
		)
		otherMatrix = NewMat4(
			0.5, 0.3, 0.2, 0.0,
			0.2, 0.8, 0.7, 0.4,
			0.1, 0.2, 0.9, 0.8,
			0.6, 0.6, 0.3, 0.1,
		)
		vector = NewVec4(2.5, 1.5, 3.0, 1.0)
	})

	Specify("NewMat4", func() {
		Expect(matrix).To(HaveMat4Elements(
			0.1, 0.2, 0.3, 0.4,
			0.5, 0.6, 0.7, 0.8,
			0.9, 1.0, 1.1, 1.2,
			1.3, 1.4, 1.5, 1.6,
		))
	})

	Specify("ZeroMat4", func() {
		Expect(ZeroMat4()).To(HaveMat4Elements(
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0,
			0.0, 0.0, 0.0, 0.0,
		))
	})

	Specify("IdentityMat4", func() {
		identityMatrix := IdentityMat4()
		transformedVector := Mat4Vec4Prod(identityMatrix, vector)
		Expect(transformedVector).To(HaveVec4Coords(vector.X, vector.Y, vector.Z, vector.W))
	})

	Specify("TransposedMat4", func() {
		result := TransposedMat4(matrix)
		Expect(result).To(HaveMat4Elements(
			0.1, 0.5, 0.9, 1.3,
			0.2, 0.6, 1.0, 1.4,
			0.3, 0.7, 1.1, 1.5,
			0.4, 0.8, 1.2, 1.6,
		))
	})

	Specify("TranslationMat4", func() {
		translationMatrix := TranslationMat4(2.0, -3.0, 4.0)
		transformedVector := Mat4Vec4Prod(translationMatrix, vector)
		Expect(transformedVector).To(HaveVec4Coords(4.5, -1.5, 7.0, 1.0))
	})

	Specify("ScaleMat4", func() {
		scaleMatrix := ScaleMat4(2.0, -3.0, 4.0)
		transformedVector := Mat4Vec4Prod(scaleMatrix, vector)
		Expect(transformedVector).To(HaveVec4Coords(5.0, -4.5, 12.0, 1.0))
	})

	Specify("RotationMat4", func() {
		vector := NewVec4(1.0, 0.0, 0.0, 1.0)
		rotationMatrix := RotationMat4(Degrees(120.0), 1.0, 1.0, 1.0)
		transformedVector := Mat4Vec4Prod(rotationMatrix, vector)
		Expect(transformedVector).To(HaveVec4Coords(0.0, 1.0, 0.0, 1.0))
	})

	Specify("TRSMat4", func() {
		translation := NewVec3(5.5, 4.4, 3.3)
		rotation := RotationQuat(Degrees(35), UnitVec3(NewVec3(1.0, 0.5, 0.25)))
		scale := NewVec3(1.2, 1.3, 1.4)
		slowTRS := Mat4MultiProd(
			TranslationMat4(translation.X, translation.Y, translation.Z),
			OrientationMat4(rotation.OrientationX(), rotation.OrientationY(), rotation.OrientationZ()),
			ScaleMat4(scale.X, scale.Y, scale.Z),
		)

		trs := TRSMat4(translation, rotation, scale)
		Expect(trs).To(HaveMat4Elements(
			slowTRS.M11, slowTRS.M12, slowTRS.M13, slowTRS.M14,
			slowTRS.M21, slowTRS.M22, slowTRS.M23, slowTRS.M24,
			slowTRS.M31, slowTRS.M32, slowTRS.M33, slowTRS.M34,
			slowTRS.M41, slowTRS.M42, slowTRS.M43, slowTRS.M44,
		))
	})

	Specify("OrthoMat4", func() {
		orthoMatrix := OrthoMat4(-1.1, 2.1, 1.5, -3.4, 1.7, 3.8)

		// test negative boundary vector projection
		nearCorner := NewVec4(-1.1, -3.4, -1.7, 1.0)
		projectedNearCorner := Mat4Vec4Prod(orthoMatrix, nearCorner)
		projectedNearCorner = Vec4Quot(projectedNearCorner, projectedNearCorner.W)
		Expect(projectedNearCorner).To(HaveVec4Coords(-1.0, -1.0, -1.0, 1.0))

		// test positive boundary vector projection
		farCorner := NewVec4(2.1, 1.5, -3.8, 1.0)
		projectedFarCorner := Mat4Vec4Prod(orthoMatrix, farCorner)
		projectedFarCorner = Vec4Quot(projectedFarCorner, projectedFarCorner.W)
		Expect(projectedFarCorner).To(HaveVec4Coords(1.0, 1.0, 1.0, 1.0))
	})

	Specify("PerspectiveMat4", func() {
		perspectiveMatrix := PerspectiveMat4(-1.1, 2.1, -3.4, 1.5, 1.7, 3.8)

		// test negative boundary vector projection
		nearCorner := NewVec4(-1.1, -3.4, -1.7, 1.0)
		projectedNearCorner := Mat4Vec4Prod(perspectiveMatrix, nearCorner)
		projectedNearCorner = Vec4Quot(projectedNearCorner, projectedNearCorner.W)
		Expect(projectedNearCorner).To(HaveVec4Coords(-1.0, -1.0, -1.0, 1.0))

		// test positive boundary vector projection
		farCorner := NewVec4(4.694117647058, 3.352941176470, -3.8, 1.0)
		projectedFarCorner := Mat4Vec4Prod(perspectiveMatrix, farCorner)
		projectedFarCorner = Vec4Quot(projectedFarCorner, projectedFarCorner.W)
		Expect(projectedFarCorner).To(HaveVec4Coords(1.0, 1.0, 1.0, 1.0))
	})

	Specify("FastInverseMat4", func() {
		matrix = IdentityMat4()
		matrix = Mat4Prod(matrix, TranslationMat4(1.5, 2.3, 3.7))
		matrix = Mat4Prod(matrix, RotationMat4(Degrees(45.0), 0.5, 0.3, 0.2))

		inverseMatrix := FastInverseMat4(matrix)
		productMatrix := Mat4Prod(inverseMatrix, matrix)
		Expect(productMatrix).To(HaveMat4Elements(
			1.0, 0.0, 0.0, 0.0,
			0.0, 1.0, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		))
	})

	Specify("InverseMat4", func() {
		matrix := NewMat4(
			4.0, 3.0, 2.0, 1.0,
			1.1, 4.1, 3.1, 2.1,
			2.2, 3.2, 4.2, 1.2,
			3.3, 2.3, 1.3, 4.3,
		)
		inverseMatrix := InverseMat4(matrix)
		productMatrix := Mat4Prod(inverseMatrix, matrix)
		Expect(productMatrix).To(HaveMat4Elements(
			1.0, 0.0, 0.0, 0.0,
			0.0, 1.0, 0.0, 0.0,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		))
	})

	Specify("TransformationMat4", func() {
		matrix := TransformationMat4(
			NewVec3(1.0, 2.0, 3.0),
			NewVec3(5.0, 6.0, 7.0),
			NewVec3(9.0, 10.0, 11.0),
			NewVec3(13.0, 14.0, 15.0),
		)
		Expect(matrix).To(HaveMat4Elements(
			1.0, 5.0, 9.0, 13.0,
			2.0, 6.0, 10.0, 14.0,
			3.0, 7.0, 11.0, 15.0,
			0.0, 0.0, 0.0, 1.0,
		))
	})

	Specify("OrientationMat4", func() {
		matrix := OrientationMat4(
			NewVec3(1.0, 2.0, 3.0),
			NewVec3(5.0, 6.0, 7.0),
			NewVec3(9.0, 10.0, 11.0),
		)
		Expect(matrix).To(HaveMat4Elements(
			1.0, 5.0, 9.0, 0.0,
			2.0, 6.0, 10.0, 0.0,
			3.0, 7.0, 11.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		))
	})

	Specify("RowMajorArrayToMat4", func() {
		matrix := RowMajorArrayToMat4([16]float64{
			1.0, 2.0, 3.0, 4.0,
			5.0, 6.0, 7.0, 8.0,
			9.0, 10.0, 11.0, 12.0,
			13.0, 14.0, 15.0, 16.0,
		})
		Expect(matrix).To(HaveMat4Elements(
			1.0, 2.0, 3.0, 4.0,
			5.0, 6.0, 7.0, 8.0,
			9.0, 10.0, 11.0, 12.0,
			13.0, 14.0, 15.0, 16.0,
		))
	})

	Specify("ColumnMajorArrayToMat4", func() {
		matrix := ColumnMajorArrayToMat4([16]float64{
			1.0, 5.0, 9.0, 13.0,
			2.0, 6.0, 10.0, 14.0,
			3.0, 7.0, 11.0, 15.0,
			4.0, 8.0, 12.0, 16.0,
		})
		Expect(matrix).To(HaveMat4Elements(
			1.0, 2.0, 3.0, 4.0,
			5.0, 6.0, 7.0, 8.0,
			9.0, 10.0, 11.0, 12.0,
			13.0, 14.0, 15.0, 16.0,
		))
	})

	Specify("Mat4Prod", func() {
		result := Mat4Prod(matrix, otherMatrix)
		Expect(result).To(HaveMat4Elements(
			0.36, 0.49, 0.55, 0.36,
			0.92, 1.25, 1.39, 0.88,
			1.48, 2.01, 2.23, 1.40,
			2.04, 2.77, 3.07, 1.92,
		))
	})

	Specify("Mat4MultiProd", func() {
		matrix = Mat4MultiProd(
			TranslationMat4(2.0, 3.0, 5.0),
			ScaleMat4(2.0, 4.0, 8.0),
		)
		vector := Mat4Vec4Prod(matrix, NewVec4(1.0, 1.0, 1.0, 1.0))
		Expect(vector).To(HaveVec4Coords(4.0, 7.0, 13.0, 1.0))
	})

	Specify("Mat4Vec4Prod", func() {
		result := Mat4Vec4Prod(matrix, vector)
		Expect(result).To(HaveVec4Coords(1.85, 5.05, 8.25, 11.45))
	})

	Specify("Mat4Vec3Transformation", func() {
		result := Mat4Vec3Transformation(matrix, NewVec3(2.5, 1.5, 3.0))
		Expect(result).To(HaveVec3Coords(1.85, 5.05, 8.25))
	})

	DescribeTable("#IsNaN",
		func(mat Mat4, expected bool) {
			Expect(mat.IsNaN()).To(Equal(expected))
		},

		Entry("standard floats", NewMat4(
			1.0, 2.0, 3.0, 4.0,
			5.0, 6.0, 7.0, 8.0,
			9.0, 10.0, 11.0, 12.0,
			13.0, 14.0, 15.0, 16.0,
		), false),

		Entry("M11 is +inf", Mat4{M11: math.Inf(1)}, false),
		Entry("M12 is +inf", Mat4{M12: math.Inf(1)}, false),
		Entry("M13 is +inf", Mat4{M13: math.Inf(1)}, false),
		Entry("M14 is +inf", Mat4{M14: math.Inf(1)}, false),
		Entry("M21 is +inf", Mat4{M21: math.Inf(1)}, false),
		Entry("M22 is +inf", Mat4{M22: math.Inf(1)}, false),
		Entry("M23 is +inf", Mat4{M23: math.Inf(1)}, false),
		Entry("M24 is +inf", Mat4{M24: math.Inf(1)}, false),
		Entry("M31 is +inf", Mat4{M31: math.Inf(1)}, false),
		Entry("M32 is +inf", Mat4{M32: math.Inf(1)}, false),
		Entry("M33 is +inf", Mat4{M33: math.Inf(1)}, false),
		Entry("M34 is +inf", Mat4{M34: math.Inf(1)}, false),
		Entry("M41 is +inf", Mat4{M41: math.Inf(1)}, false),
		Entry("M42 is +inf", Mat4{M42: math.Inf(1)}, false),
		Entry("M43 is +inf", Mat4{M43: math.Inf(1)}, false),
		Entry("M44 is +inf", Mat4{M44: math.Inf(1)}, false),

		Entry("M11 is -inf", Mat4{M11: math.Inf(-1)}, false),
		Entry("M12 is -inf", Mat4{M12: math.Inf(-1)}, false),
		Entry("M13 is -inf", Mat4{M13: math.Inf(-1)}, false),
		Entry("M14 is -inf", Mat4{M14: math.Inf(-1)}, false),
		Entry("M21 is -inf", Mat4{M21: math.Inf(-1)}, false),
		Entry("M22 is -inf", Mat4{M22: math.Inf(-1)}, false),
		Entry("M23 is -inf", Mat4{M23: math.Inf(-1)}, false),
		Entry("M24 is -inf", Mat4{M24: math.Inf(-1)}, false),
		Entry("M31 is -inf", Mat4{M31: math.Inf(-1)}, false),
		Entry("M32 is -inf", Mat4{M32: math.Inf(-1)}, false),
		Entry("M33 is -inf", Mat4{M33: math.Inf(-1)}, false),
		Entry("M34 is -inf", Mat4{M34: math.Inf(-1)}, false),
		Entry("M41 is -inf", Mat4{M41: math.Inf(-1)}, false),
		Entry("M42 is -inf", Mat4{M42: math.Inf(-1)}, false),
		Entry("M43 is -inf", Mat4{M43: math.Inf(-1)}, false),
		Entry("M44 is -inf", Mat4{M44: math.Inf(-1)}, false),

		Entry("M11 is NaN", Mat4{M11: math.NaN()}, true),
		Entry("M12 is NaN", Mat4{M12: math.NaN()}, true),
		Entry("M13 is NaN", Mat4{M13: math.NaN()}, true),
		Entry("M14 is NaN", Mat4{M14: math.NaN()}, true),
		Entry("M21 is NaN", Mat4{M21: math.NaN()}, true),
		Entry("M22 is NaN", Mat4{M22: math.NaN()}, true),
		Entry("M23 is NaN", Mat4{M23: math.NaN()}, true),
		Entry("M24 is NaN", Mat4{M24: math.NaN()}, true),
		Entry("M31 is NaN", Mat4{M31: math.NaN()}, true),
		Entry("M32 is NaN", Mat4{M32: math.NaN()}, true),
		Entry("M33 is NaN", Mat4{M33: math.NaN()}, true),
		Entry("M34 is NaN", Mat4{M34: math.NaN()}, true),
		Entry("M41 is NaN", Mat4{M41: math.NaN()}, true),
		Entry("M42 is NaN", Mat4{M42: math.NaN()}, true),
		Entry("M43 is NaN", Mat4{M43: math.NaN()}, true),
		Entry("M44 is NaN", Mat4{M44: math.NaN()}, true),
	)

	DescribeTable("#IsInf",
		func(mat Mat4, expected bool) {
			Expect(mat.IsInf()).To(Equal(expected))
		},

		Entry("standard floats", NewMat4(
			1.0, 2.0, 3.0, 4.0,
			5.0, 6.0, 7.0, 8.0,
			9.0, 10.0, 11.0, 12.0,
			13.0, 14.0, 15.0, 16.0,
		), false),

		Entry("M11 is +inf", Mat4{M11: math.Inf(1)}, true),
		Entry("M12 is +inf", Mat4{M12: math.Inf(1)}, true),
		Entry("M13 is +inf", Mat4{M13: math.Inf(1)}, true),
		Entry("M14 is +inf", Mat4{M14: math.Inf(1)}, true),
		Entry("M21 is +inf", Mat4{M21: math.Inf(1)}, true),
		Entry("M22 is +inf", Mat4{M22: math.Inf(1)}, true),
		Entry("M23 is +inf", Mat4{M23: math.Inf(1)}, true),
		Entry("M24 is +inf", Mat4{M24: math.Inf(1)}, true),
		Entry("M31 is +inf", Mat4{M31: math.Inf(1)}, true),
		Entry("M32 is +inf", Mat4{M32: math.Inf(1)}, true),
		Entry("M33 is +inf", Mat4{M33: math.Inf(1)}, true),
		Entry("M34 is +inf", Mat4{M34: math.Inf(1)}, true),
		Entry("M41 is +inf", Mat4{M41: math.Inf(1)}, true),
		Entry("M42 is +inf", Mat4{M42: math.Inf(1)}, true),
		Entry("M43 is +inf", Mat4{M43: math.Inf(1)}, true),
		Entry("M44 is +inf", Mat4{M44: math.Inf(1)}, true),

		Entry("M11 is -inf", Mat4{M11: math.Inf(-1)}, true),
		Entry("M12 is -inf", Mat4{M12: math.Inf(-1)}, true),
		Entry("M13 is -inf", Mat4{M13: math.Inf(-1)}, true),
		Entry("M14 is -inf", Mat4{M14: math.Inf(-1)}, true),
		Entry("M21 is -inf", Mat4{M21: math.Inf(-1)}, true),
		Entry("M22 is -inf", Mat4{M22: math.Inf(-1)}, true),
		Entry("M23 is -inf", Mat4{M23: math.Inf(-1)}, true),
		Entry("M24 is -inf", Mat4{M24: math.Inf(-1)}, true),
		Entry("M31 is -inf", Mat4{M31: math.Inf(-1)}, true),
		Entry("M32 is -inf", Mat4{M32: math.Inf(-1)}, true),
		Entry("M33 is -inf", Mat4{M33: math.Inf(-1)}, true),
		Entry("M34 is -inf", Mat4{M34: math.Inf(-1)}, true),
		Entry("M41 is -inf", Mat4{M41: math.Inf(-1)}, true),
		Entry("M42 is -inf", Mat4{M42: math.Inf(-1)}, true),
		Entry("M43 is -inf", Mat4{M43: math.Inf(-1)}, true),
		Entry("M44 is -inf", Mat4{M44: math.Inf(-1)}, true),

		Entry("M11 is NaN", Mat4{M11: math.NaN()}, false),
		Entry("M12 is NaN", Mat4{M12: math.NaN()}, false),
		Entry("M13 is NaN", Mat4{M13: math.NaN()}, false),
		Entry("M14 is NaN", Mat4{M14: math.NaN()}, false),
		Entry("M21 is NaN", Mat4{M21: math.NaN()}, false),
		Entry("M22 is NaN", Mat4{M22: math.NaN()}, false),
		Entry("M23 is NaN", Mat4{M23: math.NaN()}, false),
		Entry("M24 is NaN", Mat4{M24: math.NaN()}, false),
		Entry("M31 is NaN", Mat4{M31: math.NaN()}, false),
		Entry("M32 is NaN", Mat4{M32: math.NaN()}, false),
		Entry("M33 is NaN", Mat4{M33: math.NaN()}, false),
		Entry("M34 is NaN", Mat4{M34: math.NaN()}, false),
		Entry("M41 is NaN", Mat4{M41: math.NaN()}, false),
		Entry("M42 is NaN", Mat4{M42: math.NaN()}, false),
		Entry("M43 is NaN", Mat4{M43: math.NaN()}, false),
		Entry("M44 is NaN", Mat4{M44: math.NaN()}, false),
	)

	Specify("#Row1", func() {
		vector := matrix.Row1()
		Expect(vector).To(HaveVec4Coords(0.1, 0.2, 0.3, 0.4))
	})

	Specify("#Row2", func() {
		vector := matrix.Row2()
		Expect(vector).To(HaveVec4Coords(0.5, 0.6, 0.7, 0.8))
	})

	Specify("#Row3", func() {
		vector := matrix.Row3()
		Expect(vector).To(HaveVec4Coords(0.9, 1.0, 1.1, 1.2))
	})

	Specify("#Row4", func() {
		vector := matrix.Row4()
		Expect(vector).To(HaveVec4Coords(1.3, 1.4, 1.5, 1.6))
	})

	Specify("#Column1", func() {
		vector := matrix.Column1()
		Expect(vector).To(HaveVec4Coords(0.1, 0.5, 0.9, 1.3))
	})

	Specify("#Column2", func() {
		vector := matrix.Column2()
		Expect(vector).To(HaveVec4Coords(0.2, 0.6, 1.0, 1.4))
	})

	Specify("#Column3", func() {
		vector := matrix.Column3()
		Expect(vector).To(HaveVec4Coords(0.3, 0.7, 1.1, 1.5))
	})

	Specify("#Column4", func() {
		vector := matrix.Column4()
		Expect(vector).To(HaveVec4Coords(0.4, 0.8, 1.2, 1.6))
	})

	Specify("#OrientationX", func() {
		vector := matrix.OrientationX()
		Expect(vector).To(HaveVec3Coords(0.1, 0.5, 0.9))
	})

	Specify("#OrientationY", func() {
		vector := matrix.OrientationY()
		Expect(vector).To(HaveVec3Coords(0.2, 0.6, 1.0))
	})

	Specify("#OrientationZ", func() {
		vector := matrix.OrientationZ()
		Expect(vector).To(HaveVec3Coords(0.3, 0.7, 1.1))
	})

	Specify("#Translation", func() {
		vector := matrix.Translation()
		Expect(vector).To(HaveVec3Coords(0.4, 0.8, 1.2))
	})

	Specify("#Scale", func() {
		Expect(IdentityMat4().Scale()).To(HaveVec3Coords(1.0, 1.0, 1.0))
		Expect(matrix.Scale()).To(HaveVec3Coords(
			1.034408043278860045433020786732,
			1.183215956619923181492026742490,
			1.337908816025965119678176051821,
		))
	})

	Specify("#Rotation", func() {
		matrix = IdentityMat4()
		quat := matrix.Rotation()
		Expect(quat).To(HaveQuatCoords(1.0, 0.0, 0.0, 0.0))

		matrix = RotationMat4(Degrees(30), 0.0, 1.0, 0.0)
		quat = matrix.Rotation()
		rotatedVector := QuatVec3Rotation(quat, NewVec3(1.0, 0.0, 0.0))
		Expect(rotatedVector).To(HaveVec3Coords(0.86602540378443870761, 0.0, -0.5))
	})

	Specify("#TRS", func() {
		translation := NewVec3(15.0, 5.0, -3.0)
		rotation := RotationQuat(Degrees(30), BasisXVec3())
		scale := NewVec3(0.1, 0.5, 0.3)
		matrix := TRSMat4(translation, rotation, scale)

		t, r, s := matrix.TRS()
		Expect(t).To(HaveVec3Coords(translation.X, translation.Y, translation.Z))
		Expect(r).To(HaveQuatCoords(rotation.W, rotation.X, rotation.Y, rotation.Z))
		Expect(s).To(HaveVec3Coords(scale.X, scale.Y, scale.Z))
	})

	Specify("#RowMajorArray", func() {
		array := matrix.RowMajorArray()
		Expect(array[0]).To(EqualFloat64(0.1))
		Expect(array[1]).To(EqualFloat64(0.2))
		Expect(array[2]).To(EqualFloat64(0.3))
		Expect(array[3]).To(EqualFloat64(0.4))
		Expect(array[4]).To(EqualFloat64(0.5))
		Expect(array[5]).To(EqualFloat64(0.6))
		Expect(array[6]).To(EqualFloat64(0.7))
		Expect(array[7]).To(EqualFloat64(0.8))
		Expect(array[8]).To(EqualFloat64(0.9))
		Expect(array[9]).To(EqualFloat64(1.0))
		Expect(array[10]).To(EqualFloat64(1.1))
		Expect(array[11]).To(EqualFloat64(1.2))
		Expect(array[12]).To(EqualFloat64(1.3))
		Expect(array[13]).To(EqualFloat64(1.4))
		Expect(array[14]).To(EqualFloat64(1.5))
		Expect(array[15]).To(EqualFloat64(1.6))
	})

	Specify("#ColumnMajorArray", func() {
		array := matrix.ColumnMajorArray()
		Expect(array[0]).To(EqualFloat64(0.1))
		Expect(array[1]).To(EqualFloat64(0.5))
		Expect(array[2]).To(EqualFloat64(0.9))
		Expect(array[3]).To(EqualFloat64(1.3))
		Expect(array[4]).To(EqualFloat64(0.2))
		Expect(array[5]).To(EqualFloat64(0.6))
		Expect(array[6]).To(EqualFloat64(1.0))
		Expect(array[7]).To(EqualFloat64(1.4))
		Expect(array[8]).To(EqualFloat64(0.3))
		Expect(array[9]).To(EqualFloat64(0.7))
		Expect(array[10]).To(EqualFloat64(1.1))
		Expect(array[11]).To(EqualFloat64(1.5))
		Expect(array[12]).To(EqualFloat64(0.4))
		Expect(array[13]).To(EqualFloat64(0.8))
		Expect(array[14]).To(EqualFloat64(1.2))
		Expect(array[15]).To(EqualFloat64(1.6))
	})

	Specify("#GoString", func() {
		result := matrix.GoString()
		Expect(result).Should(Equal("(" +
			"(0.100000, 0.200000, 0.300000, 0.400000), " +
			"(0.500000, 0.600000, 0.700000, 0.800000), " +
			"(0.900000, 1.000000, 1.100000, 1.200000), " +
			"(1.300000, 1.400000, 1.500000, 1.600000))",
		))
	})
})
