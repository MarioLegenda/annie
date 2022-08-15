package annie

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Reading file", func() {
	Context("that are invalid", func() {
		It("should fail because of not existent file path", func() {
			_, err := NewAnnie("invalid path")
			gomega.Expect(err).ShouldNot(gomega.BeNil())
		})
	})
})

var _ = Describe("Creating annie and closing", func() {
	It("should succeed", func() {
		ann, err := NewAnnie("test_base_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should fail if config is empty file", func() {
		ann, err := NewAnnie("empty_config.yml")
		gomega.Expect(err).ShouldNot(gomega.BeNil())
		gomega.Expect(ann).Should(gomega.BeNil())

		gomega.Expect(err).ShouldNot(gomega.BeNil())
	})
})
