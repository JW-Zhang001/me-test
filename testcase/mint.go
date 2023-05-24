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
		return yearTotalBlocks, false
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

func GetWhichYearBlockReward(year uint64) uint64 {
	oneBlockReward, _ := TestInitBlockReward()
	oneBlockReward = oneBlockReward / 1000000
	if year == 1 {
		return uint64(oneBlockReward)
	} else {
		return uint64(math.Ceil(float64(oneBlockReward) / math.Pow(2, float64(year-1))))
	}
}

func BlockRewardEqualZero() {
	totalCoin := mintpb.TotalMintCoinsAmount
	var yearReward int
	for i := 1; i <= 15; i++ {
		blockReward := GetWhichYearBlockReward(uint64(i))
		zap.S().Infof("year %v, block reward = %v", i, blockReward)
		yearReward += int(blockReward * mintpb.OneYearTotalBlocks)
		zap.S().Infof("totalYearReward = %v", yearReward)
		if yearReward >= totalCoin {
			zap.S().Infof("year %v, block reward equal zero", i)
			break
		}
	}
}
