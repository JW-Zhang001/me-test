package testsuite

import (
	"time"

	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"

	"me-test/client"
	"me-test/client/query"
	"me-test/config"
	"me-test/testcase"
)

func TestKycDeposit(regionID string) (map[string]string, bool) {
	walletAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}

	if err = testcase.TestNewKyc(regionID, walletAcc["PrivKey"]); err != nil {
		return walletAcc, false
	}

	if err = testcase.TestTx(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}

	if err = testcase.TestDeposit(walletAcc["PrivKey"], config.DelegateAmount, stakepb.Term1Months); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

func TestExpireDepositWithdraw(regionID string) (map[string]string, bool) {
	walletAcc, ok := TestKycDeposit(regionID)
	if !ok {
		zap.S().Infof("TestKycDeposit error")
		return walletAcc, false
	}
	// 60s = Term1Months
	time.Sleep(time.Second * 70)

	FixedDeposit, err := query.StakeQuery.DepositByAcc(walletAcc["Addr"], stakepb.FixedDepositState_Expired)
	if err != nil {
		zap.S().Infof("Query DepositByAcc error %v", err)
		return walletAcc, false
	}
	zap.S().Infof("FixedDeposit: %v", FixedDeposit)

	if err = testcase.TestDepositWithdraw(walletAcc["PrivKey"], FixedDeposit.FixedDeposit[0].Id); err != nil {
		return walletAcc, false
	}

	return walletAcc, true
}

func TestNotExpireDepositWithdraw(regionID string) (map[string]string, bool) {
	walletAcc, ok := TestKycDeposit(regionID)
	if !ok {
		zap.S().Infof("TestKycDeposit error")
		return walletAcc, false
	}

	FixedDeposit, err := query.StakeQuery.DepositByAcc(walletAcc["Addr"], stakepb.FixedDepositState_NotExpired)
	if err != nil {
		zap.S().Infof("Query DepositByAcc error %v", err)
		return walletAcc, false
	}
	zap.S().Infof("FixedDeposit: %v", FixedDeposit)

	if err = testcase.TestDepositWithdraw(walletAcc["PrivKey"], FixedDeposit.FixedDeposit[0].Id); err != nil {
		return walletAcc, false
	}

	return walletAcc, true
}
