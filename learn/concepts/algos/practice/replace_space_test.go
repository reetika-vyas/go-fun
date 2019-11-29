package practice_test

import (
	"github.com/amanhigh/go-fun/learn/concepts/algos/practice"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReplaceSpace", func() {
	It("should encode to %20", func() {
		result := practice.ReplaceSpace("Aman Preet Singh")
		Expect(result).To(Not(BeNil()))
		Expect(result).To(Equal("Aman%20Preet%20Singh"))
	})
})
