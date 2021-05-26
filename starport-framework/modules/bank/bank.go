package bank

import (
	framework "github.com/tendermint/starport-framework/starport-framework"
)

// Bank is a standard module.
type Bank struct {
	framework.Module
}

// Name of the module.
func (b *Bank) Name() string {
	return "bank"
}
