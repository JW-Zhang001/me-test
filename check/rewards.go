package check

import (
	"fmt"
	"math"
	q "me-test/client/query"
	"me-test/config"
	"me-test/testcase"

	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"
)

const (
	TokenTotal uint64 = 20000000000000000
)

func QueryBondedPool() (*bankpb.QueryBalanceResponse, error) {
	bondedPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["bonded_tokens_pool"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("bondedPool Balance: %v", bondedPool)
	return bondedPool, nil
}

func QueryNotBondedPool() (*bankpb.QueryBalanceResponse, error) {
	notBondedPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["not_bonded_tokens_pool"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("notBondedPool Balance: %v", notBondedPool)
	return notBondedPool, nil
}

func QueryDistribution() (*bankpb.QueryBalanceResponse, error) {
	distribution, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["distribution"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("distribution Balance: %v", distribution)
	return distribution, nil
}

func CalculateTreasuryReward() (uint64, error) {
	BlockReward, ok := testcase.TestInitBlockReward()
	if !ok {
		zap.S().Infof("Get TestInitBlockReward error")
		return 0, fmt.Errorf("Get TestInitBlockReward error %v", ok)
	}
	allUserReward, err := CalculateAllUserReward()
	if err != nil {
		return 0, err
	}
	treasuryReward := BlockReward - allUserReward
	zap.S().Infof("treasuryReward: %v", treasuryReward)
	return treasuryReward, nil
}

func CalculateAllUserReward() (uint64, error) {
	BlockReward, ok := testcase.TestInitBlockReward()
	if !ok {
		zap.S().Infof("Get TestInitBlockReward error")
		return 0, fmt.Errorf("Get TestInitBlockReward error %v", ok)
	}

	bondedPool, err := QueryBondedPool()
	if err != nil {
		return 0, err
	}
	notBondedPool, err := QueryNotBondedPool()
	if err != nil {
		return 0, err
	}
	allUserDelegateAmount := bondedPool.Balance.Amount.Uint64() + notBondedPool.Balance.Amount.Uint64()
	zap.S().Infof("allUserDelegateAmount: %v", allUserDelegateAmount)

	kycList, err := q.StakeQuery.KycList(q.StakeQuery.Ctx)
	if err != nil {
		zap.S().Infof("Get kycList error: %v", err)
		return 0, err
	}
	kycNumber := len(kycList.Kyc)
	zap.S().Infof("kycNumber: %v", kycNumber)

	allDelegateAmount := allUserDelegateAmount + uint64(kycNumber*1000000)
	zap.S().Infof("allDelegateAmount: %v", allDelegateAmount)

	userReward := math.Ceil(float64(BlockReward) * (float64(allDelegateAmount) / float64(TokenTotal)))
	zap.S().Info("userReward: ", userReward)
	return uint64(userReward), nil
}
