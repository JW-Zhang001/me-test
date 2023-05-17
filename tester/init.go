package tester

import (
	"context"
	"me-test/x/staking"
)

var (
	StakeKeeper *staking.Keeper
	Cancel      context.CancelFunc

	ValidatorData ValidatorArgs
	RegionData    RegionArgs
	KycData       KycArgs

	TestSuites []Testsuite

	extract = make(map[string]string, 256)
)

type Testsuite struct {
	Step int
	Data interface{}
}

type Dependence struct {
	Extract map[string]string
}

func init() {
	StakeKeeper, Cancel = staking.NewKeeper()
	ValidatorData, _ = NewValidatorArgs()
	RegionData, _ = NewRegionArgs()
	KycData, _ = NewKycArgs()

	TestSuites = []Testsuite{
		{Step: 1, Data: ValidatorData},
		{Step: 2, Data: RegionData},
	}
}
