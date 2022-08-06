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

var _ = Describe("Creating annie", func() {
	It("should succeed", func() {
		ann, err := NewAnnie("test_base_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.Close()
	})
})

var _ = Describe("Minimal assertions", func() {
	Context("on base keys", func() {
		It("should CannotBeEmpty succeed", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty("configuration")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("should not be able to step into a node if the node is empty", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(2))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
			gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("should chain multiple assertions and fail with CannotBeEmpty", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty("configuration").
				CannotBeEmpty("options")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(2))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
			gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("should chain multiple assertions and succeed with complex yml file", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty("configuration").
				CannotBeEmpty("options")

			errs := ann.Errors()

			gomega.Expect(errs).Should(gomega.BeEmpty())

			ann.Close()
		})
	})

	Context("on base keys", func() {
		It("should step into a node and evaluate it", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").
				IsString("simpleString").
				IsNumeric("simpleNumber").
				StepInto("complex").
				StepInto("entry").
				IsArray("arrayList").
				StepOut().
				StepOut().
				IsArray("arrayList")

			errs := ann.Errors()

			gomega.Expect(errs).Should(gomega.BeEmpty())

			ann.Close()
		})
	})
})
