package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(v0 int, exp int) {
	var ans = MaxBall(v0)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		testequal(37, 10)
		testequal(45, 13)
		testequal(99, 28)
		testequal(85, 24)
		testequal(136, 39)
	})
})
