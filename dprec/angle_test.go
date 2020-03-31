package dprec_test

import (
	. "github.com/onsi/ginkgo"
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
})
