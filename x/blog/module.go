package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/client"
	"github.com/fdymylja/tmos/runtime/module"
)

type Client = client.Client

type initializer func(module.Client, *module.DescriptorBuilder)

var (
	mod          = &Module{}
	initializers []initializer
)

var _ module.Module = (*Module)(nil)

type Module struct{}

func (m *Module) Initialize(cl module.Client) module.Descriptor {
	builder := module.NewDescriptorBuilder().
		Named("blog")
	for _, gen := range initializers {
		gen(cl, builder)
	}
	return builder.Build()
}
