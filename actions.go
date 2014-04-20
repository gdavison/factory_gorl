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
	product, err := Build(name)
	if err != nil {
		return nil, err
	}
	if err := DbMap.Insert(product); err != nil {
		return nil, err
	} else {
		return product, nil
	}
}
