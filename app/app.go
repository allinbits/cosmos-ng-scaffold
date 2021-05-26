package app

import (
	"github.com/allinbits/cosmos-ng-scaffold/x/blog"
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/x/authn"
	"github.com/fdymylja/tmos/x/bank"
	abci "github.com/tendermint/tendermint/abci/types"
)

func New() abci.Application {
	rtb := runtime.NewBuilder()
	auth := authn.NewModule()
	rtb.AddModule(auth)
	// NOTE(dshulyak) GetTxDecoded needs a reference to runtime/module.Client, but apart from that
	// it doesn't need to be a part of the module.
	rtb.SetDecoder(auth.GetTxDecoder())
	rtb.AddModule(bank.NewModule())
	// NOTE(dshulyak) just an observation that NewModule is redundant in all of the current examples.
	rtb.AddModule(&blog.Module{})
	rt, err := rtb.Build()
	if err != nil {
		panic(err)
	}
	return runtime.NewABCIApplication(rt)
}
