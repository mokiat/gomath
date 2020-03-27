package dprec_test

import (
	. "github.com/onsi/ginkgo"
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
		rotationMatrix := RotationMat4(120.0, 1.0, 1.0, 1.0)
		transformedVector := Mat4Vec4Prod(rotationMatrix, vector)
		Expect(transformedVector).To(HaveVec4Coords(0.0, 1.0, 0.0, 1.0))
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
		matrix = Mat4Prod(matrix, RotationMat4(45.0, 0.5, 0.3, 0.2))

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

	Specify("Mat4Prod", func() {
		result := Mat4Prod(matrix, otherMatrix)
		Expect(result).To(HaveMat4Elements(
			0.36, 0.49, 0.55, 0.36,
			0.92, 1.25, 1.39, 0.88,
			1.48, 2.01, 2.23, 1.40,
			2.04, 2.77, 3.07, 1.92,
		))
	})

	Specify("Mat4Vec4Prod", func() {
		result := Mat4Vec4Prod(matrix, vector)
		Expect(result).To(HaveVec4Coords(1.85, 5.05, 8.25, 11.45))
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
