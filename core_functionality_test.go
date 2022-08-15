package annie

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Base actions", func() {
	Context("on any keys", func() {

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

var _ = Describe("Minimal assertions", func() {
	Context("on any keys", func() {
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

		It("should evaluate multiple configuration value", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.
				CannotBeEmpty("configuration").
				StepInto("configuration").
				CannotBeEmpty("array").
				CannotBeEmpty("simpleString").
				IsString("simpleString").
				IsArray("arrayList").
				CannotBeEmpty("complex").
				StepInto("complex").
				CannotBeEmpty("entryString").
				IsString("entryString").
				StepInto("entry").
				CannotBeEmpty("arrayList").
				IsArray("arrayList").
				StepOut().
				StepOut().
				IsNumeric("simpleNumber").
				StepOut().
				CannotBeEmpty("options").
				IsArray("options")

			errs := ann.Errors()

			gomega.Expect(errs).Should(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(0))

			ann.Close()
		})
	})
})

var _ = Describe("Real world", func() {
	Context("on docker-compose", func() {
		It("should validate", func() {
			ann, err := NewAnnie("docker-compose-config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty("version").
				IsString("version").
				IsMap("services")

			ann.Close()
		})
	})
})
