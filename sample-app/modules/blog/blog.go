package blog

import framework "github.com/tendermint/starport-framework/starport-framework"

// Blog is a custom module.
type Blog struct {
	framework.Module
}

func (b *Blog) Name() string {
	return "blog"
}
