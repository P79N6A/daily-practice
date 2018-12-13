package Hashtable

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKataTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KataTest Suite")
}
