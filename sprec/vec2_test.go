package sprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Vec2", func() {
	var nullVector Vec2
	var firstVector Vec2
	var secondVector Vec2

	BeforeEach(func() {
		nullVector = Vec2{}
		firstVector = Vec2{
			X: 2.0,
			Y: 3.0,
		}
		secondVector = Vec2{
			X: -1.0,
			Y: 2.0,
		}
	})

	Specify("NewVec2", func() {
		vector := NewVec2(9.8, 2.3)
		Expect(vector).To(HaveVec2Coords(9.8, 2.3))
	})

	Specify("ZeroVec2", func() {
		vector := ZeroVec2()
		Expect(vector).To(HaveVec2Coords(0.0, 0.0))
	})

	Specify("BasisXVec2", func() {
		vector := BasisXVec2()
		Expect(vector).To(HaveVec2Coords(1.0, 0.0))
	})

	Specify("BasisYVec2", func() {
		vector := BasisYVec2()
		Expect(vector).To(HaveVec2Coords(0.0, 1.0))
	})

	Specify("Vec2Sum", func() {
		result := Vec2Sum(firstVector, secondVector)
		Expect(result).To(HaveVec2Coords(1.0, 5.0))
	})

	Specify("Vec2MultiSum", func() {
		result := Vec2MultiSum(
			NewVec2(1.0, 2.0),
			NewVec2(3.0, 4.0),
			NewVec2(5.0, 6.0),
		)
		Expect(result).To(HaveVec2Coords(9.0, 12.0))
	})

	Specify("Vec2Diff", func() {
		result := Vec2Diff(firstVector, secondVector)
		Expect(result).To(HaveVec2Coords(3.0, 1.0))
	})

	Specify("Vec2MultiDiff", func() {
		result := Vec2MultiDiff(
			NewVec2(1.0, 2.0),
			NewVec2(3.0, 4.0),
			NewVec2(5.0, 6.0),
		)
		Expect(result).To(HaveVec2Coords(-7.0, -8.0))
	})

	Specify("Vec2Prod", func() {
		result := Vec2Prod(firstVector, 2.0)
		Expect(result).To(HaveVec2Coords(4.0, 6.0))
	})

	Specify("Vec2Quot", func() {
		result := Vec2Quot(firstVector, 2.0)
		Expect(result).To(HaveVec2Coords(1.0, 1.5))
	})

	Specify("Vec2Dot", func() {
		result := Vec2Dot(firstVector, secondVector)
		Expect(result).To(EqualFloat32(4.0))
	})

	Specify("Vec2Lerp", func() {
		first := NewVec2(1.0, 2.0)
		second := NewVec2(2.0, 1.0)
		result := Vec2Lerp(first, second, 0.25)
		Expect(result).To(HaveVec2Coords(1.25, 1.75))
	})

	Specify("UnitVec2", func() {
		result := UnitVec2(firstVector)
		Expect(result).To(HaveVec2Coords(0.554700196225, 0.832050294337))
	})

	Specify("ResizedVec2", func() {
		result := ResizedVec2(firstVector, 7.211102550927)
		Expect(result).To(HaveVec2Coords(4.0, 6.0))
	})

	Specify("InverseVec2", func() {
		result := InverseVec2(firstVector)
		Expect(result).To(HaveVec2Coords(-2.0, -3.0))
	})

	Specify("ArrayToVec2", func() {
		result := ArrayToVec2([2]float32{1.1, 2.2})
		Expect(result).To(HaveVec2Coords(1.1, 2.2))
	})

	DescribeTable("#IsNaN",
		func(vec Vec2, expected bool) {
			Expect(vec.IsNaN()).To(Equal(expected))
		},
		Entry("standard floats", NewVec2(1.0, 2.0), false),
		Entry("X is +inf", NewVec2(float32(math.Inf(1)), 2.0), false),
		Entry("Y is +inf", NewVec2(1.0, float32(math.Inf(1))), false),
		Entry("X is -inf", NewVec2(float32(math.Inf(-1)), 2.0), false),
		Entry("Y is -inf", NewVec2(1.0, float32(math.Inf(-1))), false),
		Entry("X is NaN", NewVec2(float32(math.NaN()), 2.0), true),
		Entry("Y is NaN", NewVec2(1.0, float32(math.NaN())), true),
	)

	DescribeTable("#IsInf",
		func(vec Vec2, expected bool) {
			Expect(vec.IsInf()).To(Equal(expected))
		},
		Entry("standard floats", NewVec2(1.0, 2.0), false),
		Entry("X is +inf", NewVec2(float32(math.Inf(1)), 2.0), true),
		Entry("Y is +inf", NewVec2(1.0, float32(math.Inf(1))), true),
		Entry("X is -inf", NewVec2(float32(math.Inf(-1)), 2.0), true),
		Entry("Y is -inf", NewVec2(1.0, float32(math.Inf(-1))), true),
		Entry("X is NaN", NewVec2(float32(math.NaN()), 2.0), false),
		Entry("Y is NaN", NewVec2(1.0, float32(math.NaN())), false),
	)

	Specify("#IsZero", func() {
		Expect(nullVector.IsZero()).To(BeTrue())
		Expect(firstVector.IsZero()).To(BeFalse())
		Expect(NewVec2(Epsilon, Epsilon).IsZero()).To(BeFalse())
	})

	Specify("#SqrLength", func() {
		lng := firstVector.SqrLength()
		Expect(lng).To(EqualFloat32(13.0))
	})

	Specify("#Length", func() {
		lng := firstVector.Length()
		Expect(lng).To(EqualFloat32(3.605551275463))
	})

	Specify("#GoString", func() {
		result := firstVector.GoString()
		Expect(result).To(Equal("(2.000000, 3.000000)"))
	})
})
