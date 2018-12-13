package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(height []int, exp int) {
	var ans = trap(height)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should equal", func() {
		testequal([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
		testequal([]int{0, 1, 2, 1, 0, 2, 2, 0, 3, 2}, 5)
		testequal([]int{0, 1, 0, 1, 0, 3, 0, 3, 0, 1}, 6)
	})
})
