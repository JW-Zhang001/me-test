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
