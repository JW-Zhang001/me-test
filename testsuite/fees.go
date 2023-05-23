package testsuite

import (
	"me-test/check"
	"me-test/config"
	"me-test/testcase"
	"me-test/tools"
)

func TestNotKycFees() (map[string]string, bool) {
	walletAcc, err := tools.GenWalletAcc()
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
	walletAcc, err := tools.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	toAcc, err := tools.GenWalletAcc()
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

func TestKycFeesValidatorOwnerIsUser(regionID string) (map[string]string, bool) {
	walletAcc, err := tools.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	toAcc, err := tools.GenWalletAcc()
	if err != nil {
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
