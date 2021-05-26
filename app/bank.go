package app

import "github.com/fdymylja/tmos/x/bank"

func init() {
	RegisterModule(bank.NewModule())
}
