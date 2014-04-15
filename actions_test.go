package factory_gorl_test

import (
	. "github.com/gdavison/factory_gorl"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type AutoIncrementIdTest struct {
	Id   int
	Name string
}

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

var _ = Describe("Create", func() {
	var (
		result interface{}
		err    error
	)
	BeforeEach(func() { ResetConfiguration() })
	Context("when there is a registered factory", func() {
		JustBeforeEach(func() {
			factory, _ := NewFactory("AutoIncrementIdTest", AutoIncrementIdTest{}, nil)
			RegisterFactory(factory)
		})
		JustBeforeEach(func() {
			result, err = Create("AutoIncrementIdTest")
		})
		It("has a result value", func() {
			Expect(result).ToNot(BeNil())
		})
		It("creates correct type", func() {
			Expect(result).To(BeAssignableToTypeOf(&AutoIncrementIdTest{}))
		})
		It("assigns an Id", func() {
			Expect(result.(*AutoIncrementIdTest).Id).ToNot(Equal(0))
		})
		It("saves it to the database", func() {
			retrieved, _ := DbMap.Get(AutoIncrementIdTest{}, result.(*AutoIncrementIdTest).Id)
			Expect(retrieved.(*AutoIncrementIdTest).Name).To(Equal(result.(*AutoIncrementIdTest).Name))
		})
		It("does not return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("when there is not a registered factory", func() {
		JustBeforeEach(func() {
			result, err = Create("Nonesuch")
		})
		It("has no result value", func() {
			Expect(result).To(BeNil())
		})
		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})
