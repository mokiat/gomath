package dprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/dprec"
	. "github.com/mokiat/gomath/testing/dprectest"
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
		Expect(result).To(EqualFloat64(-11.0))
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
		result := ArrayToVec4([4]float64{1.1, 2.2, 3.3, 4.4})
		Expect(result).To(HaveVec4Coords(1.1, 2.2, 3.3, 4.4))
	})

	DescribeTable("#IsNaN",
		func(vec Vec4, expected bool) {
			Expect(vec.IsNaN()).To(Equal(expected))
		},
		Entry("standard floats", NewVec4(1.0, 2.0, 3.0, 4.0), false),
		Entry("X is +inf", NewVec4(math.Inf(1), 2.0, 3.0, 4.0), false),
		Entry("Y is +inf", NewVec4(1.0, math.Inf(1), 3.0, 4.0), false),
		Entry("Z is +inf", NewVec4(1.0, 2.0, math.Inf(1), 4.0), false),
		Entry("W is +inf", NewVec4(1.0, 2.0, 3.0, math.Inf(1)), false),
		Entry("X is -inf", NewVec4(math.Inf(-1), 2.0, 3.0, 4.0), false),
		Entry("Y is -inf", NewVec4(1.0, math.Inf(-1), 3.0, 4.0), false),
		Entry("Z is -inf", NewVec4(1.0, 2.0, math.Inf(-1), 4.0), false),
		Entry("W is -inf", NewVec4(1.0, 2.0, 3.0, math.Inf(-1)), false),
		Entry("X is NaN", NewVec4(math.NaN(), 2.0, 3.0, 4.0), true),
		Entry("Y is NaN", NewVec4(1.0, math.NaN(), 3.0, 4.0), true),
		Entry("Z is NaN", NewVec4(1.0, 2.0, math.NaN(), 4.0), true),
		Entry("W is NaN", NewVec4(1.0, 2.0, 3.0, math.NaN()), true),
	)

	DescribeTable("#IsInf",
		func(vec Vec4, expected bool) {
			Expect(vec.IsInf()).To(Equal(expected))
		},
		Entry("standard floats", NewVec4(1.0, 2.0, 3.0, 4.0), false),
		Entry("X is +inf", NewVec4(math.Inf(1), 2.0, 3.0, 4.0), true),
		Entry("Y is +inf", NewVec4(1.0, math.Inf(1), 3.0, 4.0), true),
		Entry("Z is +inf", NewVec4(1.0, 2.0, math.Inf(1), 4.0), true),
		Entry("W is +inf", NewVec4(1.0, 2.0, 3.0, math.Inf(1)), true),
		Entry("X is -inf", NewVec4(math.Inf(-1), 2.0, 3.0, 4.0), true),
		Entry("Y is -inf", NewVec4(1.0, math.Inf(-1), 3.0, 4.0), true),
		Entry("Z is -inf", NewVec4(1.0, 2.0, math.Inf(-1), 4.0), true),
		Entry("W is -inf", NewVec4(1.0, 2.0, 3.0, math.Inf(-1)), true),
		Entry("X is NaN", NewVec4(math.NaN(), 2.0, 3.0, 4.0), false),
		Entry("Y is NaN", NewVec4(1.0, math.NaN(), 3.0, 4.0), false),
		Entry("Z is NaN", NewVec4(1.0, 2.0, math.NaN(), 4.0), false),
		Entry("W is NaN", NewVec4(1.0, 2.0, 3.0, math.NaN()), false),
	)

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
		Expect(firstVector.Array()).To(Equal([4]float64{
			2.0, 3.0, 4.0, 5.0,
		}))
	})

	Specify("#GoString", func() {
		result := firstVector.GoString()
		Expect(result).To(Equal("(2.000000, 3.000000, 4.000000, 5.000000)"))
	})
})
