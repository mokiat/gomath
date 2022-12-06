package sprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("QuatTest", func() {
	var quat Quat

	BeforeEach(func() {
		quat = Quat{
			W: 5.1,
			X: -4.1,
			Y: 3.1,
			Z: -2.1,
		}
	})

	Specify("NewQuat", func() {
		quat := NewQuat(2.5, 9.8, 2.3, 1.5)
		Expect(quat).To(HaveQuatCoords(2.5, 9.8, 2.3, 1.5))
	})

	Specify("IdentityQuat", func() {
		quat := IdentityQuat()
		Expect(quat).To(HaveQuatCoords(1.0, 0.0, 0.0, 0.0))
	})

	Specify("NegativeQuat", func() {
		quat := NegativeQuat(NewQuat(1.0, 2.0, 3.0, 4.0))
		Expect(quat).To(HaveQuatCoords(-1.0, -2.0, -3.0, -4.0))
	})

	Specify("RotationQuat", func() {
		quat := RotationQuat(Degrees(60.0), NewVec3(2.0, 5.0, 3.0))
		Expect(quat).To(HaveQuatCoords(0.866025403784, 0.162221421130763, 0.405553552826907, 0.243332131696144))
	})

	Specify("ConjugateQuat", func() {
		conjugate := ConjugateQuat(quat)
		Expect(conjugate).To(HaveQuatCoords(5.1, 4.1, -3.1, 2.1))
	})

	Specify("QuatScalarProd", func() {
		result := QuatScalarProd(quat, 2.0)
		Expect(result).To(HaveQuatCoords(10.2, -8.2, 6.2, -4.2))
	})

	Specify("QuatScalarQuot", func() {
		result := QuatScalarQuot(quat, 2.0)
		Expect(result).To(HaveQuatCoords(2.55, -2.05, 1.55, -1.05))
	})

	Specify("QuatProd", func() {
		other := NewQuat(1.2, 1.3, 1.4, 1.5)
		result := QuatProd(quat, other)
		Expect(result).To(HaveQuatCoords(10.26, 9.3, 14.28, -4.64))
	})

	Specify("QuatDot", func() {
		first := NewQuat(1.0, 2.0, 3.0, 4.0)
		second := NewQuat(3.0, 4.0, 5.0, 6.0)
		dot := QuatDot(first, second)
		Expect(dot).To(EqualFloat32(50.0))
	})

	Specify("QuatLerp", func() {
		first := RotationQuat(Degrees(25), NewVec3(1.0, 2.0, 3.0))
		second := RotationQuat(Degrees(45), NewVec3(1.0, 2.0, 3.0))

		lerp := QuatLerp(first, second, 0.5)
		expectedQuat := RotationQuat(Degrees(35), NewVec3(1.0, 2.0, 3.0))
		Expect(lerp).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))

		lerp = QuatLerp(first, second, 0.0)
		Expect(lerp).To(HaveQuatCoords(first.W, first.X, first.Y, first.Z))

		lerp = QuatLerp(first, second, 1.0)
		Expect(lerp).To(HaveQuatCoords(second.W, second.X, second.Y, second.Z))

		// Linear interpolation does not handle such fractions
		lerp = QuatLerp(first, second, 0.25)
		idealQuat := RotationQuat(Degrees(30), NewVec3(1.0, 2.0, 3.0))
		Expect(lerp).ToNot(HaveQuatCoords(idealQuat.W, idealQuat.X, idealQuat.Y, idealQuat.Z))

		lerp = QuatLerp(first, second, 0.75)
		idealQuat = RotationQuat(Degrees(40), NewVec3(1.0, 2.0, 3.0))
		Expect(lerp).ToNot(HaveQuatCoords(idealQuat.W, idealQuat.X, idealQuat.Y, idealQuat.Z))
	})

	Specify("QuatDiff", func() {
		first := RotationQuat(Degrees(25), NewVec3(1.0, 2.0, 3.0))
		second := RotationQuat(Degrees(45), NewVec3(1.0, 2.0, 3.0))
		expectedQuat := RotationQuat(Degrees(20), NewVec3(1.0, 2.0, 3.0))
		relative := QuatDiff(second, first, true)
		Expect(relative).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))
	})

	Specify("QuatPow", func() {
		q := RotationQuat(Degrees(20), NewVec3(1.0, 2.0, 3.0))
		pow := QuatPow(q, 2.5)
		expectedQuat := RotationQuat(Degrees(50), NewVec3(1.0, 2.0, 3.0))
		Expect(pow).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))
	})

	Specify("QuatSlerp", func() {
		first := RotationQuat(Degrees(25), NewVec3(1.0, 2.0, 3.0))
		second := RotationQuat(Degrees(45), NewVec3(1.0, 2.0, 3.0))

		slerp := QuatSlerp(first, second, 0.5)
		expectedQuat := RotationQuat(Degrees(35), NewVec3(1.0, 2.0, 3.0))
		Expect(slerp).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))

		slerp = QuatSlerp(first, second, 0.0)
		Expect(slerp).To(HaveQuatCoords(first.W, first.X, first.Y, first.Z))

		slerp = QuatSlerp(first, second, 1.0)
		Expect(slerp).To(HaveQuatCoords(second.W, second.X, second.Y, second.Z))

		// Spherical interpolation handles such fractions (unlike linear)
		slerp = QuatSlerp(first, second, 0.25)
		expectedQuat = RotationQuat(Degrees(30), NewVec3(1.0, 2.0, 3.0))
		Expect(slerp).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))

		slerp = QuatSlerp(first, second, 0.75)
		expectedQuat = RotationQuat(Degrees(40), NewVec3(1.0, 2.0, 3.0))
		Expect(slerp).To(HaveQuatCoords(expectedQuat.W, expectedQuat.X, expectedQuat.Y, expectedQuat.Z))
	})

	Specify("#QuatVec3Rotation", func() {
		quat := RotationQuat(Degrees(180), NewVec3(1.0, 1.0, 1.0))
		rotatedVector := QuatVec3Rotation(quat, NewVec3(1.0, 0.0, 0.0))
		Expect(rotatedVector).To(HaveVec3Coords(-0.333333333333, 0.666666666666, 0.666666666666))
	})

	Specify("UnitVec3", func() {
		result := UnitQuat(quat)
		Expect(result).To(HaveQuatCoords(0.676461589073, -0.543822061804, 0.411182534534, -0.278543007265))
	})

	Specify("InverseQuat", func() {
		inverse := InverseQuat(quat)
		Expect(QuatProd(quat, inverse)).To(HaveQuatCoords(1.0, 0.0, 0.0, 0.0))
	})

	DescribeTable("#IsNaN",
		func(quat Quat, expected bool) {
			Expect(quat.IsNaN()).To(Equal(expected))
		},
		Entry("standard floats", NewQuat(1.0, 2.0, 3.0, 4.0), false),
		Entry("X is +inf", NewQuat(float32(math.Inf(1)), 2.0, 3.0, 4.0), false),
		Entry("Y is +inf", NewQuat(1.0, float32(math.Inf(1)), 3.0, 4.0), false),
		Entry("Z is +inf", NewQuat(1.0, 2.0, float32(math.Inf(1)), 4.0), false),
		Entry("W is +inf", NewQuat(1.0, 2.0, 3.0, float32(math.Inf(1))), false),
		Entry("X is -inf", NewQuat(float32(math.Inf(-1)), 2.0, 3.0, 4.0), false),
		Entry("Y is -inf", NewQuat(1.0, float32(math.Inf(-1)), 3.0, 4.0), false),
		Entry("Z is -inf", NewQuat(1.0, 2.0, float32(math.Inf(-1)), 4.0), false),
		Entry("W is -inf", NewQuat(1.0, 2.0, 3.0, float32(math.Inf(-1))), false),
		Entry("X is NaN", NewQuat(float32(math.NaN()), 2.0, 3.0, 4.0), true),
		Entry("Y is NaN", NewQuat(1.0, float32(math.NaN()), 3.0, 4.0), true),
		Entry("Z is NaN", NewQuat(1.0, 2.0, float32(math.NaN()), 4.0), true),
		Entry("W is NaN", NewQuat(1.0, 2.0, 3.0, float32(math.NaN())), true),
	)

	DescribeTable("#IsInf",
		func(quat Quat, expected bool) {
			Expect(quat.IsInf()).To(Equal(expected))
		},
		Entry("standard floats", NewQuat(1.0, 2.0, 3.0, 4.0), false),
		Entry("X is +inf", NewQuat(float32(math.Inf(1)), 2.0, 3.0, 4.0), true),
		Entry("Y is +inf", NewQuat(1.0, float32(math.Inf(1)), 3.0, 4.0), true),
		Entry("Z is +inf", NewQuat(1.0, 2.0, float32(math.Inf(1)), 4.0), true),
		Entry("W is +inf", NewQuat(1.0, 2.0, 3.0, float32(math.Inf(1))), true),
		Entry("X is -inf", NewQuat(float32(math.Inf(-1)), 2.0, 3.0, 4.0), true),
		Entry("Y is -inf", NewQuat(1.0, float32(math.Inf(-1)), 3.0, 4.0), true),
		Entry("Z is -inf", NewQuat(1.0, 2.0, float32(math.Inf(-1)), 4.0), true),
		Entry("W is -inf", NewQuat(1.0, 2.0, 3.0, float32(math.Inf(-1))), true),
		Entry("X is NaN", NewQuat(float32(math.NaN()), 2.0, 3.0, 4.0), false),
		Entry("Y is NaN", NewQuat(1.0, float32(math.NaN()), 3.0, 4.0), false),
		Entry("Z is NaN", NewQuat(1.0, 2.0, float32(math.NaN()), 4.0), false),
		Entry("W is NaN", NewQuat(1.0, 2.0, 3.0, float32(math.NaN())), false),
	)

	Specify("#IsIdentity", func() {
		quat := QuatProd(
			RotationQuat(Degrees(180), BasisXVec3()),
			RotationQuat(-Degrees(180), BasisXVec3()),
		)
		Expect(quat.IsIdentity()).To(BeTrue())
	})

	Specify("#SqrNorm", func() {
		norm := quat.SqrNorm()
		Expect(norm).To(EqualFloat32(56.84))
	})

	Specify("#Norm", func() {
		norm := quat.Norm()
		Expect(norm).To(EqualFloat32(7.539230729988))
	})

	Specify("#OrientationX", func() {
		quat := RotationQuat(Degrees(180), NewVec3(1.0, 1.0, 1.0))
		Expect(quat.OrientationX()).To(HaveVec3Coords(-0.333333333333, 0.666666666666, 0.666666666666))

		quat = RotationQuat(Degrees(90), NewVec3(0.0, 0.0, 1.0))
		Expect(quat.OrientationX()).To(HaveVec3Coords(0.0, 1.0, 0.0))
	})

	Specify("#OrientationY", func() {
		quat := RotationQuat(Degrees(180), NewVec3(1.0, 1.0, 1.0))
		Expect(quat.OrientationY()).To(HaveVec3Coords(0.666666666666, -0.333333333333, 0.666666666666))

		quat = RotationQuat(Degrees(90), NewVec3(1.0, 0.0, 0.0))
		Expect(quat.OrientationY()).To(HaveVec3Coords(0.0, 0.0, 1.0))
	})

	Specify("#OrientationZ", func() {
		quat := RotationQuat(Degrees(180), NewVec3(1.0, 1.0, 1.0))
		Expect(quat.OrientationZ()).To(HaveVec3Coords(0.666666666666, 0.666666666666, -0.333333333333))

		quat = RotationQuat(Degrees(90), NewVec3(0.0, 1.0, 0.0))
		Expect(quat.OrientationZ()).To(HaveVec3Coords(1.0, 0.0, 0.0))
	})

	Specify("#GoString", func() {
		result := quat.GoString()
		Expect(result).To(Equal("(5.100000, -4.100000, 3.100000, -2.100000)"))
	})
})
