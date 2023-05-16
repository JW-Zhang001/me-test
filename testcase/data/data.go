package data

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"me-test/config"
	"me-test/initialize"
	"me-test/x/staking"
)

var (
	StakeKeeper *staking.Keeper
	Cancel      context.CancelFunc

	ValidatorData ValidatorArgs
	RegionData    RegionArgs

	TestSuites []Testsuite
)

type Testsuite struct {
	Step int
	Data interface{}
}

func init() {
	StakeKeeper, Cancel = staking.NewKeeper()
	ValidatorData, _ = NewValidatorArgs()
	RegionData, _ = NewRegionArgs()

	TestSuites = []Testsuite{
		{Step: 1, Data: ValidatorData},
		{Step: 2, Data: RegionData},
	}
}

type ValidatorArgs struct {
	PrivKey     string
	TmPubKeyStr string
	CoinStr     string
	Moniker     string
}

func NewValidatorArgs() (ValidatorArgs, error) {
	nodeID := "node4"
	tmPubK, err := initialize.GetValidatorPubKey(nodeID)
	if err != nil {
		return ValidatorArgs{}, fmt.Errorf("GetValidatorPubKey error %v", err)
	}
	return ValidatorArgs{config.SuperAdminPrivKey, tmPubK, "100mec", nodeID}, nil
}

type RegionArgs struct {
	privKey   string
	regionId  string
	name      string
	validator string
}

func RandRegionID() string {
	u := uuid.New()
	return u.String()
}

func NewRegionArgs() (RegionArgs, error) {
	regionKey := StakeKeeper.RandRegionKey()
	randomIndex := rand.Intn(len(regionKey))

	return RegionArgs{config.SuperAdminPrivKey, RandRegionID(),
		regionKey[randomIndex], "${validator}"}, nil
}
