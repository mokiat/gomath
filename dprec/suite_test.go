package dprec_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSmath(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Double Precision Math Suite")
}
