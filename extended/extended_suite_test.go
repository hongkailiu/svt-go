package extended_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExtended(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Extended Suite")
}
