package factory_gorl

import (
	"github.com/coopernurse/gorp"
)

var (
	factories map[string]*Factory
	DbMap     *gorp.DbMap
)

func RegisterFactory(factory *Factory) {
	factories[factory.Name()] = factory
}

func FactoryByName(name string) (factory *Factory) {
	factory = factories[name]
	return
}

func ResetConfiguration() {
	factories = make(map[string]*Factory)
}

func init() {
	factories = make(map[string]*Factory)
}
