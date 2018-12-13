package kata

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(t1 *Node, t2 *Node, exp bool) {
	Expect(t1.IsSimilarTo(t2)).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should equal", func() {
		t1 := NewTree(0, NewTree(1, NewTree(2, nil, nil), nil), nil)
		t2 := NewTree(2, NewTree(3, NewTree(4, nil, nil), nil), nil)

		testequal(t1, t2, true)
	})

	It("should not equal", func() {
		t1 := NewTree(0, NewTree(1, NewTree(2, nil, nil), nil), nil)
		t3 := NewTree(2, nil, NewTree(4, nil, nil))
		testequal(t1, t3, false)
	})
})
