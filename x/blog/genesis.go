package blog

import (
	"encoding/json"

	"github.com/allinbits/cosmos-ng-scaffold/x/blog/client"
	"github.com/fdymylja/tmos/runtime/module"
)

var _ module.GenesisHandler = (*genesis)(nil)

type genesis struct {
	client *client.Client
}

func (g *genesis) Default() error {
	return nil
}

// Import/Export doesn't need to be json.RawMessage just []byte

func (g *genesis) Import(json.RawMessage) error {
	return nil
}

func (g *genesis) Export() (json.RawMessage, error) {
	return nil, nil
}
