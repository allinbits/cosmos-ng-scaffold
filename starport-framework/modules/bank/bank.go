package bank

import (
	"github.com/spf13/cobra"
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

// AddCommands adds commands to the root command. It is optional to implement AddCommands.
func (b *Bank) AddCommands(root *cobra.Command) {}
