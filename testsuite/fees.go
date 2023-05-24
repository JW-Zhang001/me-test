package testsuite

import (
	"go.uber.org/zap"
	"me-test/check"
	"me-test/client"
	"me-test/config"
	"me-test/testcase"
)

func TestNotKycFees() (map[string]string, bool) {
	walletAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}

	fn := check.CheckerNotKycFees(testcase.TestTx)
	if err := fn(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

func TestKycFeesValidatorOwnerIsPM(regionID string) (map[string]string, bool) {
	walletAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	toAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}

	if err = testcase.TestNewKyc(regionID, walletAcc["PrivKey"]); err != nil {
		return walletAcc, false
	}
	if err = testcase.TestTx(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}

	fn := check.CheckerKycFeesValidatorIsPM(testcase.TestTx)
	if err := fn(walletAcc["PrivKey"], toAcc["Addr"], config.TxAmount/10); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

func TestKycFeesValidatorOwnerIsUser(nodeID string) (map[string]string, bool) {
	walletAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	toAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}

	ownerAcc, err := client.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}

	validatorID, ok := TestEditValidator(ownerAcc["Addr"], nodeID)
	if !ok {
		zap.S().Error("TestEditValidator error")
		return walletAcc, false
	}

	regionID, err := testcase.TestNewRegion(validatorID)
	if err != nil {
		zap.S().Error("TestNewRegion error")
		return walletAcc, false
	}

	if err = testcase.TestNewKyc(regionID, walletAcc["PrivKey"]); err != nil {
		return walletAcc, false
	}
	if err = testcase.TestTx(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}

	fn := check.CheckerKycFeesValidatorIsUser(testcase.TestTx)
	if err := fn(walletAcc["PrivKey"], toAcc["Addr"], config.TxAmount/10); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}
