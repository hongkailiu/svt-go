package extended_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
	"github.com/onsi/ginkgo/reporters"
	"fmt"
	"github.com/onsi/ginkgo/config"
)

func TestExtended(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("junit_%d.xml", config.GinkgoConfig.ParallelNode))
	//RunSpecs(t, "Extended Suite")
	RunSpecsWithDefaultAndCustomReporters(t, "Extended Suite", []Reporter{junitReporter})
}
