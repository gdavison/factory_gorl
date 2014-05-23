package factory_gorl

import (
	"fmt"
	"reflect"
)

type Builder func(i interface{})

type Factory struct {
	name      string
	buildType reflect.Type
	parent    *Factory
	builder   Builder
}

func NewFactory(name string, i interface{}, builder Builder) (*Factory, error) {
	f := new(Factory)
	f.name = name
	var err error
	f.buildType, err = toType(i)
	if err != nil {
		return nil, err
	}
	f.builder = builder
	return f, nil
}

func NewFactoryWithParent(name string, i interface{}, parent *Factory, builder Builder) (*Factory, error) {
	buildType, err := toType(i)
	if err != nil {
		return nil, err
	}
	if buildType != parent.buildType {
		return nil, fmt.Errorf("factory_gorl: cannot create a factory for '%v' with parent '%v'", buildType, parent.buildType)
	}
	f := new(Factory)
	f.name = name
	f.buildType = buildType
	f.parent = parent
	f.builder = builder
	return f, nil
}

func toType(i interface{}) (reflect.Type, error) {
	t := reflect.TypeOf(i)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, fmt.Errorf("factory_gorl: cannot create this type: '%v'", reflect.TypeOf(i))
	}
	return t, nil
}

func (factory *Factory) Name() string {
	return factory.name
}

func (factory *Factory) Run(override Builder) interface{} {
	result := reflect.New(factory.buildType)
	factory.build(result.Interface())
	if override != nil {
		override(result.Interface())
	}
	return result.Interface()
}

func (factory *Factory) build(i interface{}) {
	if factory.parent != nil {
		factory.parent.build(i)
	}
	if factory.builder != nil {
		factory.builder(i)
	}
}
