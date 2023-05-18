package testcase

import (
	"context"
	"me-test/x/bank"
	"me-test/x/staking"
)

var (
	StakeKeeper *staking.Keeper
	BankKeeper  *bank.Keeper
	Cancel      context.CancelFunc

	extract = make(map[string]string, 256)
)

type Dependence struct {
	Extract map[string]string
}

func init() {
	StakeKeeper, Cancel = staking.NewKeeper()
	BankKeeper, _ = bank.NewKeeper()
}
