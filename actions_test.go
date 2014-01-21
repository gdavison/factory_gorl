package factory_gorl_test

import (
	. "github.com/gdavison/factory_gorl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Build", func() {
	var (
		result interface{}
		err    error
	)
	BeforeEach(func() { ResetConfiguration() })
	Context("when there is a registered factory", func() {
		JustBeforeEach(func() {
			factory, _ := NewFactory("Test", Test{}, nil)
			RegisterFactory(factory)
		})
		JustBeforeEach(func() {
			result, err = Build("Test")
		})
		It("has a result value", func() {
			Expect(result).ToNot(BeNil())
		})
		It("creates correct type", func() {
			Expect(result).To(BeAssignableToTypeOf(&Test{}))
		})
		It("does not return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when there is not a registered factory", func() {
		JustBeforeEach(func() {
			result, err = Build("Nonesuch")
		})
		It("has no result value", func() {
			Expect(result).To(BeNil())
		})
		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})
