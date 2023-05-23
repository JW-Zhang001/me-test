package check

import (
	"fmt"
	q "me-test/client/query"
	"me-test/config"

	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"
)

func QueryTreasuryPool() (*bankpb.QueryBalanceResponse, error) {
	treasuryPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["treasury_pool"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("treasuryPool Balance: %v", treasuryPool)
	return treasuryPool, nil
}

func QueryUserBalance(addr string) (*bankpb.QueryBalanceResponse, error) {
	userBalance, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["treasury_pool"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("userBalance: %v", userBalance)
	return userBalance, nil
}

func CheckerNotKycFees(fn func(privKey, toAddr string, amount int64) error) func(privKey, toAddr string, amount int64) error {
	return func(privKey, toAddr string, amount int64) error {
		treasuryPool, err := QueryTreasuryPool()
		if err != nil {
			zap.S().Errorf("Before get balance error: %v", err)
			return err
		}
		zap.S().Info("TreasuryPool before: ", treasuryPool)

		if err := fn(privKey, toAddr, amount); err != nil {
			zap.S().Errorf("fn error: %v", err)
			return err
		}

		laterTreasuryPool, err := QueryTreasuryPool()
		if err != nil {
			zap.S().Errorf("Later get balance error: %v", err)
			return err
		}
		zap.S().Info("TreasuryPool Later: ", laterTreasuryPool)

		treasuryReward, err := CalculateTreasuryReward()
		if err != nil {
			return err
		}
		zap.S().Info("DefaultFees: ", config.DefaultFees)
		actual := laterTreasuryPool.Balance.Amount.Int64() - treasuryPool.Balance.Amount.Int64()
		expect := treasuryReward + config.DefaultFees
		if uint64(actual) != expect {
			return fmt.Errorf("actual: %v, expect: %v", actual, expect)
		}

		return nil
	}
}
