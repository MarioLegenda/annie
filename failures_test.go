package annie

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Failures and safety", func() {
	It("should fail if stepping into primitive type node and ignore rest of them from annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("options").
			CannotBeEmpty("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(2))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should fail if stepping into primitive type node and ignore rest of them from internal node", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").
			StepInto("simpleString").
			CannotBeEmpty("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(2))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should fail if stepping into primitive type from complex configuration", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").
			StepInto("complex").
			StepInto("entryString")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should fail if stepping into primitive type from complex configuration but can continue after stepping out", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").
			StepInto("complex").
			StepInto("entryString").
			StepOut().
			StepOut().
			CannotBeEmpty("configuration")

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

	It("should fail if node is not a map", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.IsMap("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

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

	It("should chain multiple assertions and fail with CannotBeEmpty on node, not annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").CannotBeEmpty("empty")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsString", func() {
		ann, err := NewAnnie("test_base_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.IsString("configuration").
			IsString("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(2))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsString on node, not annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").IsString("empty")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsNumeric", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.IsNumeric("configuration").
			IsNumeric("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(2))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsNumeric on node, not annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").IsNumeric("array")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsArray", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.IsArray("configuration").
			IsArray("string_base_value").
			IsArray("numeric_base_value")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(3))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[2]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsArray on node, not annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").IsNumeric("simpleString")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsMap", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.IsMap("string_base_value").
			IsMap("numeric_base_value")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(2))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())
		gomega.Expect(errs[1]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should chain multiple assertions and fail with IsMap on node, not annie", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("configuration").IsMap("simpleString")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})

	It("should fail if 'If' does not provide any condition function", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.If("configuration")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

		ann.Close()
	})
})
