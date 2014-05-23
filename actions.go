package factory_gorl

import (
	"fmt"
)

func Build(name string, override Builder) (interface{}, error) {
	factory := FactoryByName(name)
	if factory == nil {
		return nil, fmt.Errorf("factory_gorl: could not find Factory '%s'", name)
	}
	return factory.Run(override), nil
}

func Create(name string, override Builder) (interface{}, error) {
	product, err := Build(name, override)
	if err != nil {
		return nil, err
	}
	if err := DbMap.Insert(product); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}
