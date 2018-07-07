package challenge_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/amanhigh/go-fun/learn/algos/hackerrank/challenge"
	"github.com/amanhigh/go-fun/util"
	"github.com/amanhigh/go-fun/util/helper"
)

var _ = Describe("MaxSubSquare", func() {
	var (
		input = `3
-1 -2 -4
-8 -2 5
-3 6 7`
	)
	It("should compute sum", func() {
		scanner := util.NewStringScanner(input)
		inputMatrix := helper.ReadMatrix(scanner, util.ReadInt(scanner))
		coordinates, sum := challenge.MaximumSumSubSquare(inputMatrix)
		/* Top,Left,Bottom,Right = 1,1,2,2 */
		Expect(coordinates).To(Equal([]int{1, 1, 2, 2}))
		Expect(sum).To(Equal(16))
	})
})
