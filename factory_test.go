package factory_gorl_test

import (
	. "github.com/gdavison/factory_gorl"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type Test struct {
	foo string
	bar int
}

type OtherType struct {
	foo int64
}

var _ = Describe("NewFactory", func() {
	var (
		factory *Factory
		err     error
	)
	BeforeEach(func() {
		factory, err = NewFactory("Test", Test{}, nil)
	})
	It("has a name", func() {
		Expect(factory.Name()).To(Equal("Test"))
	})
	It("succeeds", func() {
		Expect(err).NotTo(HaveOccurred())
	})
})

var _ = Describe("NewFactoryWithParent", func() {
	var (
		factory *Factory
		err     error
	)
	Context("with same product type", func() {
		BeforeEach(func() {
			parentFactory, _ := NewFactory("Parent", Test{}, nil)
			factory, err = NewFactoryWithParent("Child", Test{}, parentFactory, nil)
		})
		It("has a name", func() {
			Expect(factory.Name()).To(Equal("Child"))
		})
		It("succeeds", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})
	Context("with a mismatched product type", func() {
		BeforeEach(func() {
			parentFactory, _ := NewFactory("Parent", OtherType{}, nil)
			factory, err = NewFactoryWithParent("Child", Test{}, parentFactory, nil)
		})
		It("is not created", func() {
			Expect(factory).To(BeNil())
		})
		It("returns an error", func() {
			Expect(err.Error()).To(Equal("factory_gorl: cannot create a factory for 'factory_gorl_test.Test' with parent 'factory_gorl_test.OtherType'"))
		})
	})
})

var _ = Describe("Type checking", func() {
	It("cannot build simple types", func() {
		aString := ""
		factory, err := NewFactory("TooSimple", aString, nil)
		Expect(factory).To(BeNil())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("factory_gorl: cannot create this type: 'string'"))
	})
})

var _ = Describe("Running Factory", func() {
	Context("Factory with no parent", func() {
		var (
			builder Builder
			factory *Factory
		)
		JustBeforeEach(func() {
			factory, _ = NewFactory("Test", Test{}, builder)
		})
		It("creates correct type", func() {
			Expect(factory.Run()).To(BeAssignableToTypeOf(&Test{}))
		})

		Context("with a nil builder func", func() {
			var (
				result *Test
			)
			BeforeEach(func() {
				builder = nil
			})
			JustBeforeEach(func() {
				result = factory.Run().(*Test)
			})
			It("doesn't set any fields", func() {
				Expect(result.foo).To(BeZero(), "foo should be Zero")
				Expect(result.bar).To(BeZero(), "bar should be Zero")
			})
		})
		Context("with a builder func", func() {
			var (
				result *Test
			)
			BeforeEach(func() {
				builder = func(i interface{}) {
					t := i.(*Test)
					t.foo = "forty-two"
					t.bar = 42
				}
			})
			JustBeforeEach(func() {
				result = factory.Run().(*Test)
			})
			It("sets the fields", func() {
				Expect(result.foo).To(Equal("forty-two"))
				Expect(result.bar).To(Equal(42))
			})
		})
		Context("with a partial builder func", func() {
			var (
				result *Test
			)
			BeforeEach(func() {
				builder = func(i interface{}) {
					t := i.(*Test)
					t.bar = 69
				}
			})
			JustBeforeEach(func() {
				result = factory.Run().(*Test)
			})
			It("sets the fields in the builder", func() {
				Expect(result.bar).To(Equal(69))
			})
			It("doesn't set the fields not in the builder", func() {
				Expect(result.foo).To(BeZero())
			})
		})
	})

	Context("Factory with a parent", func() {
		var (
			parentBuilder Builder
			childBuilder  Builder
			result        *Test
		)
		JustBeforeEach(func() {
			parent, _ := NewFactory("Parent", Test{}, parentBuilder)
			factory, _ := NewFactoryWithParent("Child", Test{}, parent, childBuilder)
			result = factory.Run().(*Test)
		})
		Context("with a nil parent builder func", func() {
			BeforeEach(func() {
				parentBuilder = nil
			})
			Context("with a nil child builder func", func() {
				BeforeEach(func() {
					childBuilder = nil
				})
				It("doesn't set any fields", func() {
					Expect(result.foo).To(BeZero(), "foo should be Zero")
					Expect(result.bar).To(BeZero(), "bar should be Zero")
				})
			})
			Context("with a child builder func", func() {
				BeforeEach(func() {
					childBuilder = func(i interface{}) {
						t := i.(*Test)
						t.foo = "forty-two"
						t.bar = 42
					}
				})
				It("sets the fields", func() {
					Expect(result.foo).To(Equal("forty-two"))
					Expect(result.bar).To(Equal(42))
				})
			})
			Context("with a partial builder func", func() {
				BeforeEach(func() {
					childBuilder = func(i interface{}) {
						t := i.(*Test)
						t.bar = 69
					}
				})
				It("sets the fields in the builder", func() {
					Expect(result.bar).To(Equal(69))
				})
				It("doesn't set the fields not in the builder", func() {
					Expect(result.foo).To(BeZero())
				})
			})
		})
		Context("with a parent builder func", func() {
			BeforeEach(func() {
				parentBuilder = func(i interface{}) {
					t := i.(*Test)
					t.foo = "forty-two"
					t.bar = 42
				}
			})
			Context("with a nil child builder func", func() {
				BeforeEach(func() {
					childBuilder = nil
				})
				It("doesn't override any fields", func() {
					Expect(result.foo).To(Equal("forty-two"))
					Expect(result.bar).To(Equal(42))
				})
			})
			Context("with a child builder func", func() {
				BeforeEach(func() {
					childBuilder = func(i interface{}) {
						t := i.(*Test)
						t.foo = "sixty-nine"
						t.bar = 69
					}
				})
				It("sets the fields", func() {
					Expect(result.foo).To(Equal("sixty-nine"))
					Expect(result.bar).To(Equal(69))
				})
			})
			Context("with a partial builder func", func() {
				BeforeEach(func() {
					childBuilder = func(i interface{}) {
						t := i.(*Test)
						t.bar = 69
					}
				})
				It("sets the fields in the builder", func() {
					Expect(result.bar).To(Equal(69))
				})
				It("doesn't set the fields not in the builder", func() {
					Expect(result.foo).To(Equal("forty-two"))
				})
			})
		})
	})
})
