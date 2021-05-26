package app

import (
	"github.com/fdymylja/tmos/runtime"
	"github.com/fdymylja/tmos/runtime/authentication"
	"github.com/fdymylja/tmos/runtime/module"
	abci "github.com/tendermint/tendermint/abci/types"
)

var app = runtime.NewBuilder()

func RegisterModule(mod module.Module) {
	app.AddModule(mod)
}

func SetDecoder(dec authentication.TxDecoder) {
	app.SetDecoder(dec)
}

func Build() abci.Application {
	rt, err := app.Build()
	if err != nil {
		panic(err)
	}
	return runtime.NewABCIApplication(rt)
}
