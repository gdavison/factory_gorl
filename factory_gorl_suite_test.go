package factory_gorl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFactory_gorl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factory_gorl Suite")
}
