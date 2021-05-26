package auth

import framework "github.com/tendermint/starport-framework/starport-framework"

// Auth is a standard module.
type Auth struct {
	framework.Module
}

func (a *Auth) Name() string {
	return "auth"
}
