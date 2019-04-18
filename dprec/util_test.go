package dprec_test

import (
	. "github.com/onsi/ginkgo"
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

})
