package testcase

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
	DelegateData  DelegateArgs

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

//func init() {
//	StakeKeeper, Cancel = staking.NewKeeper()
//	ValidatorData, _ = NewValidatorArgs()
//	RegionData, _ = NewRegionArgs()
//	KycData, _ = NewKycArgs()
//	DelegateData, _ = NewDelegateArgs(1000000)
//
//	TestSuites = []Testsuite{
//		{Step: 1, Data: ValidatorData},
//		{Step: 2, Data: RegionData},
//	}
//}
