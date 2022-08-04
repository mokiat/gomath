package sprec_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Vec4", func() {
	var nullVector Vec4
	var firstVector Vec4
	var secondVector Vec4

	BeforeEach(func() {
		nullVector = Vec4{}
		firstVector = Vec4{
			X: 2.0,
			Y: 3.0,
			Z: 4.0,
			W: 5.0,
		}
		secondVector = Vec4{
			X: -1.0,
			Y: 2.0,
			Z: 5.0,
			W: -7.0,
		}
	})

	Specify("NewVec4", func() {
		vector := NewVec4(9.8, 2.3, 1.5, 5.7)
		Expect(vector).To(HaveVec4Coords(9.8, 2.3, 1.5, 5.7))
	})

	Specify("ZeroVec4", func() {
		vector := ZeroVec4()
		Expect(vector).To(HaveVec4Coords(0.0, 0.0, 0.0, 0.0))
	})

	Specify("Vec4Sum", func() {
		result := Vec4Sum(firstVector, secondVector)
		Expect(result).To(HaveVec4Coords(1.0, 5.0, 9.0, -2.0))
	})

	Specify("Vec4Diff", func() {
		result := Vec4Diff(firstVector, secondVector)
		Expect(result).To(HaveVec4Coords(3.0, 1.0, -1.0, 12.0))
	})

	Specify("Vec4Prod", func() {
		result := Vec4Prod(firstVector, 2.0)
		Expect(result).To(HaveVec4Coords(4.0, 6.0, 8.0, 10.0))
	})

	Specify("Vec4Quot", func() {
		result := Vec4Quot(firstVector, 2.0)
		Expect(result).To(HaveVec4Coords(1.0, 1.5, 2.0, 2.5))
	})

	Specify("Vec4Dot", func() {
		result := Vec4Dot(firstVector, secondVector)
		Expect(result).To(EqualFloat32(-11.0))
	})

	Specify("Vec4Lerp", func() {
		first := NewVec4(1.0, 2.0, 3.0, 4.0)
		second := NewVec4(5.0, 4.0, 3.0, 2.0)
		result := Vec4Lerp(first, second, 0.25)
		Expect(result).To(HaveVec4Coords(2.0, 2.5, 3.0, 3.5))
	})

	Specify("InverseVec4", func() {
		result := InverseVec4(firstVector)
		Expect(result).To(HaveVec4Coords(-2.0, -3.0, -4.0, -5.0))
	})

	Specify("ArrayToVec4", func() {
		result := ArrayToVec4([4]float32{1.1, 2.2, 3.3, 4.4})
		Expect(result).To(HaveVec4Coords(1.1, 2.2, 3.3, 4.4))
	})

	Specify("#IsZero", func() {
		Expect(nullVector.IsZero()).To(BeTrue())
		Expect(firstVector.IsZero()).To(BeFalse())
		Expect(NewVec4(Epsilon, 0.0, 0.0, 0.0).IsZero()).To(BeFalse())
		Expect(NewVec4(0.0, Epsilon, 0.0, 0.0).IsZero()).To(BeFalse())
		Expect(NewVec4(0.0, 0.0, Epsilon, 0.0).IsZero()).To(BeFalse())
		Expect(NewVec4(0.0, 0.0, 0.0, Epsilon).IsZero()).To(BeFalse())
	})

	Specify("#VecXYZ", func() {
		Expect(firstVector.VecXYZ()).To(HaveVec3Coords(2.0, 3.0, 4.0))
	})

	Specify("#Array", func() {
		Expect(firstVector.Array()).To(Equal([4]float32{
			2.0, 3.0, 4.0, 5.0,
		}))
	})

	Specify("#GoString", func() {
		result := firstVector.GoString()
		Expect(result).To(Equal("(2.000000, 3.000000, 4.000000, 5.000000)"))
	})
})
