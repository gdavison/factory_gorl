package factory_gorl

import (
	"fmt"
)

func Build(name string) (interface{}, error) {
	factory := FactoryByName(name)
	if factory == nil {
		return nil, fmt.Errorf("factory_gorl: could not find Factory '%s'", name)
	}
	return factory.Run(), nil
}

func Create(name string) (interface{}, error) {
	factory := FactoryByName(name)
	if factory == nil {
		return nil, fmt.Errorf("factory_gorl: could not find Factory '%s'", name)
	}
	product := factory.Run()
	if err := DbMap.Insert(product); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}
