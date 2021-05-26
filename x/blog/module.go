package blog

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog/client"
	"github.com/fdymylja/tmos/runtime/module"
)

type Client = client.Client

type initializer func(client module.Client) func(*module.DescriptorBuilder)

var generatedInitializers []initializer

var _ module.Module = (*Module)(nil)

type Module struct{}

func (m *Module) Initialize(cl module.Client) module.Descriptor {
	builder := module.NewDescriptorBuilder().
		Named("blog").
		WithGenesis(&genesis{client: &Client{Client: cl}})

	for _, gen := range generatedInitializers {
		gen(cl)(builder)
	}
	return builder.Build()
}
