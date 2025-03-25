package sprec_test

import (
	"math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Util", func() {

	Specify("Abs", func() {
		Expect(Abs(float32(-0.1))).To(EqualFloat32(0.1))
		Expect(Abs(float32(-13.57))).To(EqualFloat32(13.57))
		Expect(Abs(float32(11.01))).To(EqualFloat32(11.01))
	})

	Specify("Max", func() {
		Expect(Max(float32(1.0), float32(2.0))).To(EqualFloat32(2.0))
		Expect(Max(float32(1.0), float32(-1.0))).To(EqualFloat32(1.0))
		Expect(Max(float32(5.0), float32(5.0))).To(EqualFloat32(5.0))
	})

	Specify("Min", func() {
		Expect(Min(float32(1.0), float32(2.0))).To(EqualFloat32(1.0))
		Expect(Min(float32(1.0), float32(-1.0))).To(EqualFloat32(-1.0))
		Expect(Min(float32(5.0), float32(5.0))).To(EqualFloat32(5.0))
	})

	Specify("Sum", func() {
		Expect(Sum(float32(1.0), float32(2.0))).To(EqualFloat32(3.0))
		Expect(Sum(float32(1.0), float32(-1.0))).To(EqualFloat32(0.0))
	})

	Specify("Sqr", func() {
		Expect(Sqr(float32(2.0))).To(EqualFloat32(4.0))
		Expect(Sqr(float32(-3.0))).To(EqualFloat32(9.0))
		Expect(Sqr(float32(0.0))).To(EqualFloat32(0.0))
	})

	Specify("Floor", func() {
		Expect(Floor(float32(1.1))).To(EqualFloat32(1.0))
		Expect(Floor(float32(5.9))).To(EqualFloat32(5.0))
	})

	Specify("Ceil", func() {
		Expect(Ceil(float32(1.1))).To(EqualFloat32(2.0))
		Expect(Ceil(float32(5.9))).To(EqualFloat32(6.0))
	})

	Specify("Clamp", func() {
		Expect(Clamp(float32(1.0), float32(2.0), float32(3.0))).To(EqualFloat32(2.0))
		Expect(Clamp(float32(2.5), float32(2.0), float32(3.0))).To(EqualFloat32(2.5))
		Expect(Clamp(float32(4.0), float32(2.0), float32(3.0))).To(EqualFloat32(3.0))
	})

	Specify("Mix", func() {
		Expect(Mix(float32(1.0), float32(2.0), float32(0.0))).To(EqualFloat32(1.0))
		Expect(Mix(float32(1.0), float32(2.0), float32(1.0))).To(EqualFloat32(2.0))
		Expect(Mix(float32(1.0), float32(2.0), float32(0.5))).To(EqualFloat32(1.5))
	})

	Specify("Eq", func() {
		Expect(Eq(0.000001, 0.000001)).To(BeTrue())
		Expect(Eq(0.000001, 0.000002)).To(BeFalse())
		Expect(Eq(0.0000003, 0.0000002)).To(BeTrue()) // outside precision
	})

	Specify("EqEps", func() {
		Expect(EqEps(0.01, 0.01, 0.01)).To(BeTrue())
		Expect(EqEps(0.01, 0.02, 0.01)).To(BeFalse())
		Expect(EqEps(0.003, 0.002, 0.01)).To(BeTrue()) // outside precision
	})

	Specify("Mod", func() {
		Expect(Mod(5.0, 3.0)).To(EqualFloat32(2.0))
	})

	Specify("Sqrt", func() {
		Expect(Sqrt(17.64)).To(EqualFloat32(4.2))
	})

	Specify("Pow", func() {
		Expect(Pow(2.0, 4.0)).To(EqualFloat32(16.0))
	})

	Specify("Cos", func() {
		Expect(Cos(Radians(0.0))).To(EqualFloat32(1.0))
		Expect(Cos(Radians(Pi / 3))).To(EqualFloat32(0.5))
		Expect(Cos(Radians(Pi / 2))).To(EqualFloat32(0.0))
	})

	Specify("Acos", func() {
		Expect(Acos(1.0).Degrees()).To(EqualFloat32(0))
		Expect(Acos(0.5).Degrees()).To(EqualFloat32(60))
		Expect(Acos(0.0).Degrees()).To(EqualFloat32(90))
	})

	Specify("Sin", func() {
		Expect(Sin(Radians(0.0))).To(EqualFloat32(0.0))
		Expect(Sin(Radians(Pi / 6))).To(EqualFloat32(0.5))
		Expect(Sin(Radians(Pi / 2))).To(EqualFloat32(1.0))
	})

	Specify("Asin", func() {
		Expect(Asin(1.0).Degrees()).To(EqualFloat32(90))
		Expect(Asin(0.5).Degrees()).To(EqualFloat32(30))
		Expect(Asin(0.0).Degrees()).To(EqualFloat32(0))
	})

	Specify("Atan2", func() {
		Expect(Atan2(1.0, 1.0).Degrees()).To(EqualFloat32(45))
		Expect(Atan2(1.0, 0.0).Degrees()).To(EqualFloat32(90))
		Expect(Atan2(0.0, 1.0).Degrees()).To(EqualFloat32(0))
	})

	Specify("Tan", func() {
		Expect(Tan(Radians(0.0))).To(EqualFloat32(0.0))
		Expect(Tan(Radians(Pi / 3))).To(EqualFloat32(float32(math.Sqrt(3.0))))
		Expect(Tan(Radians(Pi / 4))).To(EqualFloat32(1.0))
	})

	Specify("Sign", func() {
		Expect(Sign(0.1)).To(EqualFloat32(1.0))
		Expect(Sign(-0.1)).To(EqualFloat32(-1.0))
	})

	Specify("IsNegative", func() {
		Expect(IsNegative(float32(0.1))).To(BeFalse())
		Expect(IsNegative(float32(0.0))).To(BeFalse())
		Expect(IsNegative(float32(-0.1))).To(BeTrue())
	})

	Specify("IsValid", func() {
		Expect(IsValid(float32(0.0))).To(BeTrue())
		Expect(IsValid(float32(-15.0))).To(BeTrue())
		Expect(IsValid(float32(3.4))).To(BeTrue())
		Expect(IsValid(float32(math.NaN()))).To(BeFalse())
		Expect(IsValid(float32(math.Inf(1)))).To(BeFalse())
		Expect(IsValid(float32(math.Inf(-1)))).To(BeFalse())
	})

})
