package app

import "github.com/allinbits/cosmos-ng-scaffold/x/blog"

func init() {
	RegisterModule(&blog.Module{})
}
