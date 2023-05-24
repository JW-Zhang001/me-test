package check

import (
	"fmt"

	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"

	"me-test/client"
	q "me-test/client/query"
	"me-test/config"
)

func QueryTreasuryPool() (*bankpb.QueryBalanceResponse, error) {
	treasuryPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["treasury_pool"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("treasuryPool Balance: %v", treasuryPool)
	return treasuryPool, nil
}

func QueryProjectManage() (*bankpb.QueryBalanceResponse, error) {
	PMBalance, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.BaseAccountList["PM"])
	if err != nil {
		return nil, err
	}
	zap.S().Infof("ProjectManage Balance: %v", PMBalance)
	return PMBalance, nil
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

func CheckerKycFeesValidatorIsPM(fn func(privKey, toAddr string, amount int64) error) func(privKey, toAddr string, amount int64) error {
	return func(privKey, toAddr string, amount int64) error {
		treasuryPool, err := QueryTreasuryPool()
		if err != nil {
			zap.S().Errorf("Before get balance error: %v", err)
			return err
		}
		zap.S().Info("TreasuryPool Before: ", treasuryPool)

		PMBalance, err := QueryProjectManage()
		if err != nil {
			zap.S().Errorf("Before get PMBalance error: %v", err)
			return err
		}

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

		laterPMBalance, err := QueryProjectManage()
		if err != nil {
			zap.S().Errorf("Later get PMBalance error: %v", err)
			return err
		}

		actual1 := laterPMBalance.Balance.Amount.Int64() - PMBalance.Balance.Amount.Int64()
		expect1 := config.DefaultFees * 0.2
		zap.S().Info("PM allot fees:", expect1)
		if float64(actual1) != expect1 {
			return fmt.Errorf("PM balance fail actual: %v, expect: %v", actual1, expect1)
		}

		treasuryReward, err := CalculateTreasuryReward()
		if err != nil {
			return err
		}

		actual := laterTreasuryPool.Balance.Amount.Int64() - treasuryPool.Balance.Amount.Int64()
		expect := treasuryReward + (config.DefaultFees * 0.8)
		zap.S().Info("treasury allot fees:", expect)
		if uint64(actual) != expect {
			return fmt.Errorf("actual: %v, expect: %v", actual, expect)
		}
		return nil
	}
}

func CheckerKycFeesValidatorIsUser(fn func(privKey, toAddr string, amount int64) error) func(privKey, toAddr string, amount int64) error {
	return func(privKey, toAddr string, amount int64) error {
		treasuryPool, err := QueryTreasuryPool()
		if err != nil {
			zap.S().Errorf("Before get balance error: %v", err)
			return err
		}
		zap.S().Info("TreasuryPool Before: ", treasuryPool)

		PMBalance, err := QueryProjectManage()
		if err != nil {
			zap.S().Errorf("Before get PMBalance error: %v", err)
			return err
		}
		fromAddr, _ := client.GetAccAddrStr(privKey)

		ownerAddr, err := GetValidatorOwner(fromAddr)
		if err != nil {
			zap.S().Errorf("GetValidatorOwner error: %v", err)
			return err
		}
		zap.S().Info("ownerAddr: ", ownerAddr)
		ownerAddrBalance, _ := q.BankQuery.Balance(q.BankQuery.Ctx, ownerAddr)
		zap.S().Info("ownerAddrBalance: ", ownerAddrBalance)

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

		laterPMBalance, err := QueryProjectManage()
		if err != nil {
			zap.S().Errorf("Later get PMBalance error: %v", err)
			return err
		}

		laterOwnerAddrBalance, _ := q.BankQuery.Balance(q.BankQuery.Ctx, ownerAddr)
		zap.S().Info("laterOwnerAddrBalance: ", laterOwnerAddrBalance)

		actual2 := laterOwnerAddrBalance.Balance.Amount.Int64() - ownerAddrBalance.Balance.Amount.Int64()
		expect2 := config.DefaultFees * 0.1
		zap.S().Info("validatorOwner allot fees:", expect2)
		if float64(actual2) != expect2 {
			return fmt.Errorf("validatorOwner balance fail actual: %v, expect: %v", actual2, expect2)
		}

		actual1 := laterPMBalance.Balance.Amount.Int64() - PMBalance.Balance.Amount.Int64()
		expect1 := config.DefaultFees * 0.1
		zap.S().Info("PM allot fees:", expect1)
		if float64(actual1) != expect1 {
			return fmt.Errorf("PM balance fail actual: %v, expect: %v", actual1, expect1)
		}

		treasuryReward, err := CalculateTreasuryReward()
		if err != nil {
			return err
		}

		actual := laterTreasuryPool.Balance.Amount.Int64() - treasuryPool.Balance.Amount.Int64()
		expect := treasuryReward + (config.DefaultFees * 0.8)
		zap.S().Info("treasury allot fees:", expect)
		if uint64(actual) != expect {
			return fmt.Errorf("actual: %v, expect: %v", actual, expect)
		}
		return nil
	}
}

func GetValidatorOwner(kycUserAddr string) (string, error) {
	kycInfo, err := q.StakeQuery.ShowKyc(q.StakeQuery.Ctx, kycUserAddr)
	if err != nil {
		return "", err
	}

	regionInfo, err := q.StakeQuery.ShowRegion(q.StakeQuery.Ctx, kycInfo.Kyc.RegionId)
	if err != nil {
		return "", err
	}

	validatorInfo, err := q.StakeQuery.ShowValidator(q.StakeQuery.Ctx, regionInfo.Region.OperatorAddress)
	if err != nil {
		return "", err
	}
	return validatorInfo.Validator.OwnerAddress, nil
}
