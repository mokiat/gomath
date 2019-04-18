package sprec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mokiat/gomath/sprec"
	. "github.com/mokiat/gomath/testing/sprectest"
)

var _ = Describe("Vec3", func() {
	Specify("NewVec3", func() {
		vector := NewVec3(9.8, 2.3, 1.5)
		Expect(vector).To(HaveVec3Coords(9.8, 2.3, 1.5))
	})
})
