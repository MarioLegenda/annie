package annie

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = Describe("Argument validation", func() {
	Context("should fail", func() {
		It("when using CannotBeEmpty (empty argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using CannotBeEmpty (crowded argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.CannotBeEmpty("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsArray (empty argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsArray()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsArray (crowded argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsArray("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsMap (empty argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsMap()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsMap (crowded argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsMap("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsString (empty argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsString()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsString (crowded argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsString("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsNumeric (empty argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsNumeric()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsNumeric (crowded argument list)", func() {
			ann, err := NewAnnie("test_base_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.IsNumeric("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsNumeric in node (empty argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsNumeric()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsNumeric (crowded argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsNumeric("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsString in node (empty argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsString()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsString (crowded argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsString("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsArray in node (empty argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsArray()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsArray (crowded argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsArray("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsMap in node (empty argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsMap()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using IsMap (crowded argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").IsMap("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using CannotBeEmpty in node (empty argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").CannotBeEmpty()

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})

		It("when using CannotBeEmpty (crowded argument list)", func() {
			ann, err := NewAnnie("test_complex_config.yml")
			gomega.Expect(err).Should(gomega.BeNil())
			gomega.Expect(ann).ShouldNot(gomega.BeNil())

			ann.StepInto("configuration").CannotBeEmpty("", "", "")

			errs := ann.Errors()

			gomega.Expect(errs).ShouldNot(gomega.BeEmpty())
			gomega.Expect(errs).Should(gomega.HaveLen(1))
			gomega.Expect(errs[0]).ShouldNot(gomega.BeNil())

			ann.Close()
		})
	})
})
