package sprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Vec3", func() {
	var nullVector Vec3
	var firstVector Vec3
	var secondVector Vec3

	BeforeEach(func() {
		nullVector = Vec3{}
		firstVector = Vec3{
			X: 2.0,
			Y: 3.0,
			Z: 4.0,
		}
		secondVector = Vec3{
			X: -1.0,
			Y: 2.0,
			Z: 5.0,
		}
	})

	Specify("NewVec3", func() {
		vector := NewVec3(9.8, 2.3, 1.5)
		Expect(vector).To(HaveVec3Coords(9.8, 2.3, 1.5))
	})

	Specify("ZeroVec3", func() {
		vector := ZeroVec3()
		Expect(vector).To(HaveVec3Coords(0.0, 0.0, 0.0))
	})

	Specify("BasisXVec3", func() {
		vector := BasisXVec3()
		Expect(vector).To(HaveVec3Coords(1.0, 0.0, 0.0))
	})

	Specify("BasisYVec3", func() {
		vector := BasisYVec3()
		Expect(vector).To(HaveVec3Coords(0.0, 1.0, 0.0))
	})

	Specify("BasisZVec3", func() {
		vector := BasisZVec3()
		Expect(vector).To(HaveVec3Coords(0.0, 0.0, 1.0))
	})

	Specify("Vec3Sum", func() {
		result := Vec3Sum(firstVector, secondVector)
		Expect(result).To(HaveVec3Coords(1.0, 5.0, 9.0))
	})

	Specify("Vec3Diff", func() {
		result := Vec3Diff(firstVector, secondVector)
		Expect(result).To(HaveVec3Coords(3.0, 1.0, -1.0))
	})

	Specify("Vec3Prod", func() {
		result := Vec3Prod(firstVector, 2.0)
		Expect(result).To(HaveVec3Coords(4.0, 6.0, 8.0))
	})

	Specify("Vec3Quot", func() {
		result := Vec3Quot(firstVector, 2.0)
		Expect(result).To(HaveVec3Coords(1.0, 1.5, 2.0))
	})

	Specify("Vec3Dot", func() {
		result := Vec3Dot(firstVector, secondVector)
		Expect(result).To(EqualFloat32(24.0))
	})

	Specify("Vec3Cross", func() {
		result := Vec3Cross(firstVector, secondVector)
		Expect(result).To(HaveVec3Coords(7.0, -14.0, 7.0))
	})

	Specify("Vec3Lerp", func() {
		first := NewVec3(1.0, 2.0, 3.0)
		second := NewVec3(3.0, 2.0, 1.0)
		result := Vec3Lerp(first, second, 0.25)
		Expect(result).To(HaveVec3Coords(1.5, 2.0, 2.5))
	})

	Specify("UnitVec3", func() {
		result := UnitVec3(firstVector)
		Expect(result).To(HaveVec3Coords(0.371390676354, 0.557086014531, 0.742781352708))
	})

	Specify("ResizedVec3", func() {
		result := ResizedVec3(firstVector, 10.770329614269)
		Expect(result).To(HaveVec3Coords(4.0, 6.0, 8.0))
	})

	Specify("InverseVec3", func() {
		result := InverseVec3(firstVector)
		Expect(result).To(HaveVec3Coords(-2.0, -3.0, -4.0))
	})

	Specify("ArrayToVec3", func() {
		result := ArrayToVec3([3]float32{1.1, 2.2, 3.3})
		Expect(result).To(HaveVec3Coords(1.1, 2.2, 3.3))
	})

	DescribeTable("#IsNaN",
		func(vec Vec3, expected bool) {
			Expect(vec.IsNaN()).To(Equal(expected))
		},
		Entry("standard floats", NewVec3(1.0, 2.0, 3.0), false),
		Entry("X is +inf", NewVec3(float32(math.Inf(1)), 2.0, 3.0), false),
		Entry("Y is +inf", NewVec3(1.0, float32(math.Inf(1)), 3.0), false),
		Entry("Z is +inf", NewVec3(1.0, 2.0, float32(math.Inf(1))), false),
		Entry("X is -inf", NewVec3(float32(math.Inf(-1)), 2.0, 3.0), false),
		Entry("Y is -inf", NewVec3(1.0, float32(math.Inf(-1)), 3.0), false),
		Entry("Z is -inf", NewVec3(1.0, 2.0, float32(math.Inf(-1))), false),
		Entry("X is NaN", NewVec3(float32(math.NaN()), 2.0, 3.0), true),
		Entry("Y is NaN", NewVec3(1.0, float32(math.NaN()), 3.0), true),
		Entry("Z is NaN", NewVec3(1.0, 2.0, float32(math.NaN())), true),
	)

	DescribeTable("#IsInf",
		func(vec Vec3, expected bool) {
			Expect(vec.IsInf()).To(Equal(expected))
		},
		Entry("standard floats", NewVec3(1.0, 2.0, 3.0), false),
		Entry("X is +inf", NewVec3(float32(math.Inf(1)), 2.0, 3.0), true),
		Entry("Y is +inf", NewVec3(1.0, float32(math.Inf(1)), 3.0), true),
		Entry("Z is +inf", NewVec3(1.0, 2.0, float32(math.Inf(1))), true),
		Entry("X is -inf", NewVec3(float32(math.Inf(-1)), 2.0, 3.0), true),
		Entry("Y is -inf", NewVec3(1.0, float32(math.Inf(-1)), 3.0), true),
		Entry("Z is -inf", NewVec3(1.0, 2.0, float32(math.Inf(-1))), true),
		Entry("X is NaN", NewVec3(float32(math.NaN()), 2.0, 3.0), false),
		Entry("Y is NaN", NewVec3(1.0, float32(math.NaN()), 3.0), false),
		Entry("Z is NaN", NewVec3(1.0, 2.0, float32(math.NaN())), false),
	)

	Specify("#IsZero", func() {
		Expect(nullVector.IsZero()).To(BeTrue())
		Expect(firstVector.IsZero()).To(BeFalse())
		Expect(NewVec3(Epsilon, Epsilon, Epsilon).IsZero()).To(BeFalse())
	})

	Specify("#SqrLength", func() {
		lng := firstVector.SqrLength()
		Expect(lng).To(EqualFloat32(29))
	})

	Specify("#Length", func() {
		lng := firstVector.Length()
		Expect(lng).To(EqualFloat32(5.385164807134))
	})

	Specify("#Array", func() {
		Expect(firstVector.Array()).To(Equal([3]float32{
			2.0, 3.0, 4.0,
		}))
	})

	Specify("#GoString", func() {
		result := firstVector.GoString()
		Expect(result).To(Equal("(2.000000, 3.000000, 4.000000)"))
	})
})
