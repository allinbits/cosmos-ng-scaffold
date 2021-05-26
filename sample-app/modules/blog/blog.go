package blog

import (
	framework "github.com/tendermint/starport-framework/starport-framework"
)

// Blog is a custom module.
type Blog struct {
	framework.Module
}

// Name returns name of the module.
func (b *Blog) Name() string {
	return "blog"
}
