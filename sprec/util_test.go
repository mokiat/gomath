package sprec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Util", func() {

	Specify("Abs", func() {
		Expect(Abs(-0.1)).To(EqualFloat32(0.1))
		Expect(Abs(-13.57)).To(EqualFloat32(13.57))
		Expect(Abs(11.01)).To(EqualFloat32(11.01))
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

	Specify("Sqrt", func() {
		Expect(Sqrt(17.64)).To(EqualFloat32(4.2))
	})

})
