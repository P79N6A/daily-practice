package Hashtable

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testequal(t1 uint64, exp uint64) {
	Expect(t1).To(Equal(exp))
}

var _ = Describe("Test Example", func() {

	It("should equal", func() {
		d := NewDict(16)
		d.Insert(NewThing("Jack", 21).Hashcode())
		d.Insert(NewThing("Tom", 12).Hashcode())
		d.Insert(NewThing("Chuck", 23).Hashcode())
		d.Insert(NewThing("Tim", 18).Hashcode())
		d.Insert(NewThing("Peter", 46).Hashcode())
		d.Insert(NewThing("Allen", 34).Hashcode())
		d.Insert(NewThing("Bob", 32).Hashcode())
	})

})
