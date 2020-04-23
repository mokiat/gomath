package sprec_test

import (
	. "github.com/onsi/ginkgo"
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
