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

var _ = Describe("Failures and safety", func() {
	It("should fail if stepping into primitive type node and ignore rest of them", func() {
		ann, err := NewAnnie("test_complex_config.yml")
		gomega.Expect(err).Should(gomega.BeNil())
		gomega.Expect(ann).ShouldNot(gomega.BeNil())

		ann.StepInto("options").
			CannotBeEmpty("options")

		errs := ann.Errors()

		gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
		gomega.Expect(errs).Should(gomega.HaveLen(1))
		gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

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
})

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
