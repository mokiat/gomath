package sprec_test

import (
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
})
