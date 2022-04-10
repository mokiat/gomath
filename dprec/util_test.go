package dprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/dprec"
	. "github.com/mokiat/gomath/testing/dprectest"
)

var _ = Describe("Util", func() {

	Specify("Abs", func() {
		Expect(Abs(-0.1)).To(EqualFloat64(0.1))
		Expect(Abs(-13.57)).To(EqualFloat64(13.57))
		Expect(Abs(11.01)).To(EqualFloat64(11.01))
	})

	Specify("Max", func() {
		Expect(Max(1.0, 2.0)).To(EqualFloat64(2.0))
		Expect(Max(1.0, -1.0)).To(EqualFloat64(1.0))
		Expect(Max(5.0, 5.0)).To(EqualFloat64(5.0))
	})

	Specify("Min", func() {
		Expect(Min(1.0, 2.0)).To(EqualFloat64(1.0))
		Expect(Min(1.0, -1.0)).To(EqualFloat64(-1.0))
		Expect(Min(5.0, 5.0)).To(EqualFloat64(5.0))
	})

	Specify("Clamp", func() {
		Expect(Clamp(1.0, 2.0, 3.0)).To(EqualFloat64(2.0))
		Expect(Clamp(2.5, 2.0, 3.0)).To(EqualFloat64(2.5))
		Expect(Clamp(4.0, 2.0, 3.0)).To(EqualFloat64(3.0))
	})

	Specify("Eq", func() {
		Expect(Eq(0.000000000001, 0.000000000001)).To(BeTrue())
		Expect(Eq(0.000000000001, 0.000000000002)).To(BeFalse())
		Expect(Eq(0.0000000000003, 0.0000000000002)).To(BeTrue()) // outside precision
	})

	Specify("EqEps", func() {
		Expect(EqEps(0.01, 0.01, 0.01)).To(BeTrue())
		Expect(EqEps(0.01, 0.02, 0.01)).To(BeFalse())
		Expect(EqEps(0.003, 0.002, 0.01)).To(BeTrue()) // outside precision
	})

	Specify("Sqrt", func() {
		Expect(Sqrt(17.64)).To(EqualFloat64(4.2))
	})

	Specify("Cos", func() {
		Expect(Cos(Radians(0.0))).To(EqualFloat64(1.0))
		Expect(Cos(Radians(Pi / 3))).To(EqualFloat64(0.5))
		Expect(Cos(Radians(Pi / 2))).To(EqualFloat64(0.0))
	})

	Specify("Sin", func() {
		Expect(Sin(Radians(0.0))).To(EqualFloat64(0.0))
		Expect(Sin(Radians(Pi / 6))).To(EqualFloat64(0.5))
		Expect(Sin(Radians(Pi / 2))).To(EqualFloat64(1.0))
	})

	Specify("Tan", func() {
		Expect(Tan(Radians(0.0))).To(EqualFloat64(0.0))
		Expect(Tan(Radians(Pi / 3))).To(EqualFloat64(math.Sqrt(3.0)))
		Expect(Tan(Radians(Pi / 4))).To(EqualFloat64(1.0))
	})

	Specify("Sign", func() {
		Expect(Sign(0.1)).To(EqualFloat64(1.0))
		Expect(Sign(-0.1)).To(EqualFloat64(-1.0))
	})

})
