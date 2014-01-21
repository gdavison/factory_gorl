package factory_gorl

import (
	"fmt"
)

func Build(name string) (interface{}, error) {
	factory:= FactoryByName(name)
	if factory== nil {
		return nil, fmt.Errorf("factory_gorl: could not find Factory '%s'", name)
	}
	return factory.Run(), nil
}
