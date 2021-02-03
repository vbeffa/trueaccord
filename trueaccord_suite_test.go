package trueaccord_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTrueaccord(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Trueaccord Suite")
}
