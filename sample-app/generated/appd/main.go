package main

import (
	"context"

	_ "github.com/tendermint/starport-framework/sample-app/generated"

	framework "github.com/tendermint/starport-framework/starport-framework"
)

func main() {
	framework.Run(context.TODO())
}
