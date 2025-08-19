package dprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/dprec"
	. "github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Angle", func() {
	var flat Angle
	var vertical Angle
	var inclined Angle

	Describe("Degrees", func() {
		BeforeEach(func() {
			flat = Degrees(0.0)
			inclined = Degrees(45.0)
			vertical = Degrees(90.0)
		})

		Specify("#Degrees", func() {
			Expect(flat.Degrees()).To(EqualFloat64(0.0))
			Expect(inclined.Degrees()).To(EqualFloat64(45.0))
			Expect(vertical.Degrees()).To(EqualFloat64(90.0))
		})

		Specify("#Radians", func() {
			Expect(flat.Radians()).To(EqualFloat64(0.0))
			Expect(inclined.Radians()).To(EqualFloat64(Pi / 4.0))
			Expect(vertical.Radians()).To(EqualFloat64(Pi / 2.0))
		})
	})

	Describe("Radians", func() {
		BeforeEach(func() {
			flat = Radians(0.0)
			inclined = Radians(Pi / 4.0)
			vertical = Radians(Pi / 2.0)
		})

		Specify("#Degrees", func() {
			Expect(flat.Degrees()).To(EqualFloat64(0.0))
			Expect(inclined.Degrees()).To(EqualFloat64(45.0))
			Expect(vertical.Degrees()).To(EqualFloat64(90.0))
		})

		Specify("#Radians", func() {
			Expect(flat.Radians()).To(EqualFloat64(0.0))
			Expect(inclined.Radians()).To(EqualFloat64(Pi / 4.0))
			Expect(vertical.Radians()).To(EqualFloat64(Pi / 2.0))
		})
	})

	DescribeTable("NormalizeAngle",
		func(angle, expected Angle) {
			Expect(NormalizeAngle(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Degrees(0.0), Degrees(0.0)),
		Entry("-90", Degrees(-90.0), Degrees(-90.0)),
		Entry("90", Degrees(90.0), Degrees(90.0)),
		Entry("-180", Degrees(-180.0), Degrees(-180.0)),
		Entry("180", Degrees(180.0), Degrees(180.0)),
		Entry("-270", Degrees(-270.0), Degrees(90.0)),
		Entry("270", Degrees(270.0), Degrees(-90.0)),
		Entry("-360", Degrees(-360.0), Degrees(0.0)),
		Entry("360", Degrees(360.0), Degrees(0.0)),
		Entry("-540", Degrees(-540.0), Degrees(-180.0)),
		Entry("540", Degrees(540.0), Degrees(180.0)),
	)

	DescribeTable("NormalizeAnglePos",
		func(angle, expected Angle) {
			Expect(NormalizeAnglePos(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Degrees(0.0), Degrees(0.0)),
		Entry("-90", Degrees(-90.0), Degrees(270.0)),
		Entry("90", Degrees(90.0), Degrees(90.0)),
		Entry("-180", Degrees(-180.0), Degrees(180.0)),
		Entry("180", Degrees(180.0), Degrees(180.0)),
		Entry("-270", Degrees(-270.0), Degrees(90.0)),
		Entry("270", Degrees(270.0), Degrees(270.0)),
		Entry("-360", Degrees(-360.0), Degrees(0.0)),
		Entry("360", Degrees(360.0), Degrees(0.0)),
		Entry("-540", Degrees(-540.0), Degrees(180.0)),
		Entry("540", Degrees(540.0), Degrees(180.0)),
	)

	DescribeTable("NormalizeAngleNeg",
		func(angle, expected Angle) {
			Expect(NormalizeAngleNeg(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Degrees(0.0), Degrees(0.0)),
		Entry("-90", Degrees(-90.0), Degrees(-90.0)),
		Entry("90", Degrees(90.0), Degrees(-270.0)),
		Entry("-180", Degrees(-180.0), Degrees(-180.0)),
		Entry("180", Degrees(180.0), Degrees(-180.0)),
		Entry("-270", Degrees(-270.0), Degrees(-270.0)),
		Entry("270", Degrees(270.0), Degrees(-90.0)),
		Entry("-360", Degrees(-360.0), Degrees(0.0)),
		Entry("360", Degrees(360.0), Degrees(0.0)),
		Entry("-540", Degrees(-540.0), Degrees(-180.0)),
		Entry("540", Degrees(540.0), Degrees(-180.0)),
	)

	DescribeTable("#IsNaN",
		func(angle Angle, expected bool) {
			Expect(angle.IsNaN()).To(Equal(expected))
		},
		Entry("standard float", Angle(0.0), false),
		Entry("+inf", Angle(math.Inf(1)), false),
		Entry("-inf", Angle(math.Inf(-1)), false),
		Entry("NaN", Angle(math.NaN()), true),
	)

	DescribeTable("#IsInf",
		func(angle Angle, expected bool) {
			Expect(angle.IsInf()).To(Equal(expected))
		},
		Entry("standard float", Angle(0.0), false),
		Entry("+inf", Angle(math.Inf(1)), true),
		Entry("-inf", Angle(math.Inf(-1)), true),
		Entry("NaN", Angle(math.NaN()), false),
	)
})
