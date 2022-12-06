package dprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/dprec"
	. "github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Mat3", func() {
	var matrix Mat3
	var secondMatrix Mat3
	var thirdMatrix Mat3
	var vector Vec3

	BeforeEach(func() {
		matrix = NewMat3(
			0.1, 0.2, 0.3,
			0.4, 0.5, 0.6,
			0.7, 0.8, 0.9,
		)
		secondMatrix = NewMat3(
			1.1, 1.2, 1.3,
			1.4, 1.5, 1.6,
			1.7, 1.8, 1.9,
		)
		thirdMatrix = NewMat3(
			2.1, 2.2, 2.3,
			2.4, 2.5, 2.6,
			2.7, 2.8, 2.9,
		)
		vector = NewVec3(2.5, 3.5, 1.0)
	})

	Specify("NewMat3", func() {
		Expect(matrix).To(HaveMat3Elements(
			0.1, 0.2, 0.3,
			0.4, 0.5, 0.6,
			0.7, 0.8, 0.9,
		))
	})

	Specify("ZeroMat3", func() {
		Expect(ZeroMat3()).To(HaveMat3Elements(
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
			0.0, 0.0, 0.0,
		))
	})

	Specify("IdentityMat3", func() {
		identityMatrix := IdentityMat3()
		transformedVector := Mat3Vec3Prod(identityMatrix, vector)
		Expect(transformedVector).To(HaveVec3Coords(vector.X, vector.Y, vector.Z))
	})

	Specify("TransposedMat3", func() {
		result := TransposedMat3(matrix)
		Expect(result).To(HaveMat3Elements(
			0.1, 0.4, 0.7,
			0.2, 0.5, 0.8,
			0.3, 0.6, 0.9,
		))
	})

	Specify("TranslationMat3", func() {
		translationMatrix := TranslationMat3(2.0, -3.0)
		transformedVector := Mat3Vec3Prod(translationMatrix, vector)
		Expect(transformedVector).To(HaveVec3Coords(4.5, 0.5, 1.0))
	})

	Specify("ScaleMat3", func() {
		scaleMatrix := ScaleMat3(2.0, -3.0)
		transformedVector := Mat3Vec3Prod(scaleMatrix, vector)
		Expect(transformedVector).To(HaveVec3Coords(5.0, -10.5, 1.0))
	})

	Specify("RotationMat3", func() {
		vector := NewVec3(1.0, 0.0, 1.0)
		rotationMatrix := RotationMat3(Degrees(30.0))
		transformedVector := Mat3Vec3Prod(rotationMatrix, vector)
		Expect(transformedVector).To(HaveVec3Coords(0.866025403784, 0.5, 1.0))
	})

	Specify("OrthoMat3", func() {
		orthoMatrix := OrthoMat3(-1.1, 2.1, 1.5, -3.4)

		// test bottom left boundary vector projection
		bottomLeftCorner := NewVec3(-1.1, -3.4, 1.0)
		projectedBottomLeftCorner := Mat3Vec3Prod(orthoMatrix, bottomLeftCorner)
		projectedBottomLeftCorner = Vec3Quot(projectedBottomLeftCorner, projectedBottomLeftCorner.Z)
		Expect(projectedBottomLeftCorner).To(HaveVec3Coords(-1.0, -1.0, 1.0))

		// test top right boundary vector projection
		topRightCorner := NewVec3(2.1, 1.5, 1.0)
		projectedTopRightCorner := Mat3Vec3Prod(orthoMatrix, topRightCorner)
		projectedTopRightCorner = Vec3Quot(projectedTopRightCorner, projectedTopRightCorner.Z)
		Expect(projectedTopRightCorner).To(HaveVec3Coords(1.0, 1.0, 1.0))
	})

	Specify("FastInverseMat3", func() {
		matrix = IdentityMat3()
		matrix = Mat3Prod(matrix, TranslationMat3(1.5, 2.3))
		matrix = Mat3Prod(matrix, RotationMat3(Degrees(45.0)))

		inverseMatrix := FastInverseMat3(matrix)
		productMatrix := Mat3Prod(inverseMatrix, matrix)
		Expect(productMatrix).To(HaveMat3Elements(
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
		))
	})

	Specify("InverseMat3", func() {
		matrix := NewMat3(
			4.0, 3.0, 2.0,
			1.1, 4.1, 3.1,
			2.2, 3.2, 4.2,
		)
		inverseMatrix := InverseMat3(matrix)
		productMatrix := Mat3Prod(inverseMatrix, matrix)
		Expect(productMatrix).To(HaveMat3Elements(
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0,
		))
	})

	Specify("TransformationMat3", func() {
		matrix := TransformationMat3(
			NewVec2(1.0, 2.0),
			NewVec2(5.0, 6.0),
			NewVec2(9.0, 10.0),
		)
		Expect(matrix).To(HaveMat3Elements(
			1.0, 5.0, 9.0,
			2.0, 6.0, 10.0,
			0.0, 0.0, 1.0,
		))
	})

	Specify("RowMajorArrayToMat3", func() {
		matrix := RowMajorArrayToMat3([9]float64{
			1.0, 2.0, 3.0,
			5.0, 6.0, 7.0,
			9.0, 10.0, 11.0,
		})
		Expect(matrix).To(HaveMat3Elements(
			1.0, 2.0, 3.0,
			5.0, 6.0, 7.0,
			9.0, 10.0, 11.0,
		))
	})

	Specify("ColumnMajorArrayToMat3", func() {
		matrix := ColumnMajorArrayToMat3([9]float64{
			1.0, 5.0, 9.0,
			2.0, 6.0, 10.0,
			3.0, 7.0, 11.0,
		})
		Expect(matrix).To(HaveMat3Elements(
			1.0, 2.0, 3.0,
			5.0, 6.0, 7.0,
			9.0, 10.0, 11.0,
		))
	})

	Specify("Mat3Prod", func() {
		result := Mat3Prod(matrix, secondMatrix)
		Expect(result).To(HaveMat3Elements(
			0.9, 0.96, 1.02,
			2.16, 2.31, 2.46,
			3.42, 3.66, 3.9,
		))
	})

	Specify("Mat3MultiProd", func() {
		multiResult := Mat3MultiProd(
			matrix,
			secondMatrix,
			thirdMatrix,
		)
		manualResult := matrix
		manualResult = Mat3Prod(manualResult, secondMatrix)
		manualResult = Mat3Prod(manualResult, thirdMatrix)
		Expect(multiResult).To(HaveMat3Elements(
			manualResult.M11, manualResult.M12, manualResult.M13,
			manualResult.M21, manualResult.M22, manualResult.M23,
			manualResult.M31, manualResult.M32, manualResult.M33,
		))
	})

	Specify("Mat3Vec3Prod", func() {
		result := Mat3Vec3Prod(matrix, vector)
		Expect(result).To(HaveVec3Coords(1.25, 3.35, 5.45))
	})

	DescribeTable("#IsNaN",
		func(mat Mat3, expected bool) {
			Expect(mat.IsNaN()).To(Equal(expected))
		},

		Entry("standard floats", NewMat3(
			1.0, 2.0, 3.0,
			4.0, 5.0, 6.0,
			7.0, 8.0, 9.0,
		), false),

		Entry("M11 is +inf", Mat3{M11: math.Inf(1)}, false),
		Entry("M12 is +inf", Mat3{M12: math.Inf(1)}, false),
		Entry("M13 is +inf", Mat3{M13: math.Inf(1)}, false),
		Entry("M21 is +inf", Mat3{M21: math.Inf(1)}, false),
		Entry("M22 is +inf", Mat3{M22: math.Inf(1)}, false),
		Entry("M23 is +inf", Mat3{M23: math.Inf(1)}, false),
		Entry("M31 is +inf", Mat3{M31: math.Inf(1)}, false),
		Entry("M32 is +inf", Mat3{M32: math.Inf(1)}, false),
		Entry("M33 is +inf", Mat3{M33: math.Inf(1)}, false),

		Entry("M11 is -inf", Mat3{M11: math.Inf(-1)}, false),
		Entry("M12 is -inf", Mat3{M12: math.Inf(-1)}, false),
		Entry("M13 is -inf", Mat3{M13: math.Inf(-1)}, false),
		Entry("M21 is -inf", Mat3{M21: math.Inf(-1)}, false),
		Entry("M22 is -inf", Mat3{M22: math.Inf(-1)}, false),
		Entry("M23 is -inf", Mat3{M23: math.Inf(-1)}, false),
		Entry("M31 is -inf", Mat3{M31: math.Inf(-1)}, false),
		Entry("M32 is -inf", Mat3{M32: math.Inf(-1)}, false),
		Entry("M33 is -inf", Mat3{M33: math.Inf(-1)}, false),

		Entry("M11 is NaN", Mat3{M11: math.NaN()}, true),
		Entry("M12 is NaN", Mat3{M12: math.NaN()}, true),
		Entry("M13 is NaN", Mat3{M13: math.NaN()}, true),
		Entry("M21 is NaN", Mat3{M21: math.NaN()}, true),
		Entry("M22 is NaN", Mat3{M22: math.NaN()}, true),
		Entry("M23 is NaN", Mat3{M23: math.NaN()}, true),
		Entry("M31 is NaN", Mat3{M31: math.NaN()}, true),
		Entry("M32 is NaN", Mat3{M32: math.NaN()}, true),
		Entry("M33 is NaN", Mat3{M33: math.NaN()}, true),
	)

	DescribeTable("#IsInf",
		func(mat Mat3, expected bool) {
			Expect(mat.IsInf()).To(Equal(expected))
		},

		Entry("standard floats", NewMat3(
			1.0, 2.0, 3.0,
			4.0, 5.0, 6.0,
			7.0, 8.0, 9.0,
		), false),

		Entry("M11 is +inf", Mat3{M11: math.Inf(1)}, true),
		Entry("M12 is +inf", Mat3{M12: math.Inf(1)}, true),
		Entry("M13 is +inf", Mat3{M13: math.Inf(1)}, true),
		Entry("M21 is +inf", Mat3{M21: math.Inf(1)}, true),
		Entry("M22 is +inf", Mat3{M22: math.Inf(1)}, true),
		Entry("M23 is +inf", Mat3{M23: math.Inf(1)}, true),
		Entry("M31 is +inf", Mat3{M31: math.Inf(1)}, true),
		Entry("M32 is +inf", Mat3{M32: math.Inf(1)}, true),
		Entry("M33 is +inf", Mat3{M33: math.Inf(1)}, true),

		Entry("M11 is -inf", Mat3{M11: math.Inf(-1)}, true),
		Entry("M12 is -inf", Mat3{M12: math.Inf(-1)}, true),
		Entry("M13 is -inf", Mat3{M13: math.Inf(-1)}, true),
		Entry("M21 is -inf", Mat3{M21: math.Inf(-1)}, true),
		Entry("M22 is -inf", Mat3{M22: math.Inf(-1)}, true),
		Entry("M23 is -inf", Mat3{M23: math.Inf(-1)}, true),
		Entry("M31 is -inf", Mat3{M31: math.Inf(-1)}, true),
		Entry("M32 is -inf", Mat3{M32: math.Inf(-1)}, true),
		Entry("M33 is -inf", Mat3{M33: math.Inf(-1)}, true),

		Entry("M11 is NaN", Mat3{M11: math.NaN()}, false),
		Entry("M12 is NaN", Mat3{M12: math.NaN()}, false),
		Entry("M13 is NaN", Mat3{M13: math.NaN()}, false),
		Entry("M21 is NaN", Mat3{M21: math.NaN()}, false),
		Entry("M22 is NaN", Mat3{M22: math.NaN()}, false),
		Entry("M23 is NaN", Mat3{M23: math.NaN()}, false),
		Entry("M31 is NaN", Mat3{M31: math.NaN()}, false),
		Entry("M32 is NaN", Mat3{M32: math.NaN()}, false),
		Entry("M33 is NaN", Mat3{M33: math.NaN()}, false),
	)

	Specify("#Row1", func() {
		vector := matrix.Row1()
		Expect(vector).To(HaveVec3Coords(0.1, 0.2, 0.3))
	})

	Specify("#Row2", func() {
		vector := matrix.Row2()
		Expect(vector).To(HaveVec3Coords(0.4, 0.5, 0.6))
	})

	Specify("#Row3", func() {
		vector := matrix.Row3()
		Expect(vector).To(HaveVec3Coords(0.7, 0.8, 0.9))
	})

	Specify("#Column1", func() {
		vector := matrix.Column1()
		Expect(vector).To(HaveVec3Coords(0.1, 0.4, 0.7))
	})

	Specify("#Column2", func() {
		vector := matrix.Column2()
		Expect(vector).To(HaveVec3Coords(0.2, 0.5, 0.8))
	})

	Specify("#Column3", func() {
		vector := matrix.Column3()
		Expect(vector).To(HaveVec3Coords(0.3, 0.6, 0.9))
	})

	Specify("#OrientationX", func() {
		vector := matrix.OrientationX()
		Expect(vector).To(HaveVec2Coords(0.1, 0.4))
	})

	Specify("#OrientationY", func() {
		vector := matrix.OrientationY()
		Expect(vector).To(HaveVec2Coords(0.2, 0.5))
	})

	Specify("#Translation", func() {
		vector := matrix.Translation()
		Expect(vector).To(HaveVec2Coords(0.3, 0.6))
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
	})

	Specify("#ColumnMajorArray", func() {
		array := matrix.ColumnMajorArray()
		Expect(array[0]).To(EqualFloat64(0.1))
		Expect(array[1]).To(EqualFloat64(0.4))
		Expect(array[2]).To(EqualFloat64(0.7))
		Expect(array[3]).To(EqualFloat64(0.2))
		Expect(array[4]).To(EqualFloat64(0.5))
		Expect(array[5]).To(EqualFloat64(0.8))
		Expect(array[6]).To(EqualFloat64(0.3))
		Expect(array[7]).To(EqualFloat64(0.6))
		Expect(array[8]).To(EqualFloat64(0.9))
	})

	Specify("#GoString", func() {
		result := matrix.GoString()
		Expect(result).Should(Equal("(" +
			"(0.100000, 0.200000, 0.300000), " +
			"(0.400000, 0.500000, 0.600000), " +
			"(0.700000, 0.800000, 0.900000))",
		))
	})
})
