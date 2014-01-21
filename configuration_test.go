package factory_gorl_test

import (
	. "github.com/gdavison/factory_gorl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configuration", func() {
	It("finds a registered factory", func() {
		factory, _ := NewFactory("Test", Test{}, nil)
		RegisterFactory(factory)
		Expect(FactoryByName(factory.Name())).To(Equal(factory))
	})
})
