package testcase

import (
	mintpb "github.com/cosmos/cosmos-sdk/x/mint/types"
	"go.uber.org/zap"
	"math"
)

const (
	BlockInterval  = 5
	InitYearReward = 5000000000
)

func TestOneYearTotalBlocks() (uint64, bool) {
	yearTotalBlocks := uint64(365 * 24 * 60 * 60 / BlockInterval)
	if yearTotalBlocks != mintpb.OneYearTotalBlocks {
		zap.S().Infof("yearTotalBlocks = %v, not eq mintpb.OneYearTotalBlocks = %v: ", yearTotalBlocks, mintpb.OneYearTotalBlocks)
		return uint64(yearTotalBlocks), false
	}

	return yearTotalBlocks, true
}

func TestInitBlockReward() (uint64, bool) {
	initBlockReward := math.Ceil(float64(InitYearReward) / float64(mintpb.OneYearTotalBlocks))
	if initBlockReward != mintpb.InitialMintAmount+1 {
		zap.S().Infof("initBlockReward = %v, not eq mintpb.InitialMintAmount = %v: ", initBlockReward, mintpb.InitialMintAmount+1)
		return uint64(initBlockReward), false
	}

	return uint64(initBlockReward * 1000000), true
}
