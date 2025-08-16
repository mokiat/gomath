package sprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
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
			Expect(flat.Degrees()).To(EqualFloat32(0.0))
			Expect(inclined.Degrees()).To(EqualFloat32(45.0))
			Expect(vertical.Degrees()).To(EqualFloat32(90.0))
		})

		Specify("#Radians", func() {
			Expect(flat.Radians()).To(EqualFloat32(0.0))
			Expect(inclined.Radians()).To(EqualFloat32(Pi / 4.0))
			Expect(vertical.Radians()).To(EqualFloat32(Pi / 2.0))
		})
	})

	Describe("Radians", func() {
		BeforeEach(func() {
			flat = Radians(0.0)
			inclined = Radians(Pi / 4.0)
			vertical = Radians(Pi / 2.0)
		})

		Specify("#Degrees", func() {
			Expect(flat.Degrees()).To(EqualFloat32(0.0))
			Expect(inclined.Degrees()).To(EqualFloat32(45.0))
			Expect(vertical.Degrees()).To(EqualFloat32(90.0))
		})

		Specify("#Radians", func() {
			Expect(flat.Radians()).To(EqualFloat32(0.0))
			Expect(inclined.Radians()).To(EqualFloat32(Pi / 4.0))
			Expect(vertical.Radians()).To(EqualFloat32(Pi / 2.0))
		})
	})

	DescribeTable("NormalizeAngle",
		func(angle, expected Angle) {
			Expect(NormalizeAngle(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Radians(0.0), Radians(0.0)),
		Entry("-pi", Radians(-Pi), Radians(-Pi)),
		Entry("pi", Radians(Pi), Radians(Pi)),
		Entry("-tau", Radians(-Tau), Radians(0.0)),
		Entry("tau", Radians(Tau), Radians(0.0)),
		Entry("-tau - pi", Radians(-Tau-Pi), Radians(-Pi)),
		Entry("tau + pi", Radians(Tau+Pi), Radians(Pi)),
	)

	DescribeTable("NormalizeAnglePos",
		func(angle, expected Angle) {
			Expect(NormalizeAnglePos(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Radians(0.0), Radians(0.0)),
		Entry("-pi", Radians(-Pi), Radians(Pi)),
		Entry("pi", Radians(Pi), Radians(Pi)),
		Entry("-tau", Radians(-Tau), Radians(0.0)),
		Entry("tau", Radians(Tau), Radians(0.0)),
		Entry("-tau - pi", Radians(-Tau-Pi), Radians(Pi)),
		Entry("tau + pi", Radians(Tau+Pi), Radians(Pi)),
	)

	DescribeTable("NormalizeAngleNeg",
		func(angle, expected Angle) {
			Expect(NormalizeAngleNeg(angle)).To(BeNumerically("~", expected, Epsilon))
		},
		Entry("zero", Radians(0.0), Radians(0.0)),
		Entry("-pi", Radians(-Pi), Radians(-Pi)),
		Entry("pi", Radians(Pi), Radians(-Pi)),
		Entry("-tau", Radians(-Tau), Radians(0.0)),
		Entry("tau", Radians(Tau), Radians(0.0)),
		Entry("-tau - pi", Radians(-Tau-Pi), Radians(-Pi)),
		Entry("tau + pi", Radians(Tau+Pi), Radians(-Pi)),
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
