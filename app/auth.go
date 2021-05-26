package app

import (
	"github.com/fdymylja/tmos/x/authn"
)

func init() {
	auth := authn.NewModule()
	RegisterModule(auth)
	SetDecoder(auth.GetTxDecoder())
}
