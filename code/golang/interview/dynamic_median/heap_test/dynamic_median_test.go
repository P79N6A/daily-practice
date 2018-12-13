package heap_test

import (
	. "../heap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("validate correctness of dynamic-median algorithm", func() {
	It("case #1", func() {
		array := []int{12, 100, 16, 4, 8, 70, 2, 36, 22, 5, 46, 56, 112, 233, 666}
		dm := Init(&array)

		expected := 36
		actual := dm.Median()

		Expect(expected).To(Equal(actual))
	})

	It("case #2", func() {
		array := []int{12, 100, 16, 4, 8, 70, 2, 36, 22, 5, 46, 56, 112}
		dm := Init(&array)

		expected := 22
		actual := dm.Median()

		Expect(expected).To(Equal(actual))
	})
})
