package testsuite

import (
	"me-test/check"
	"me-test/config"
	"me-test/testcase"
	"me-test/tools"
)

/*
TestNewValidatorRegion
@Description: Create a new validator and region
@param nodeID
@return bool
@return string regionID
*/
func TestNewValidatorRegion(nodeID string) (string, bool) {

	fn := check.CheckerNewValidator(testcase.TestNewValidator)
	validatorID, err := fn(nodeID, config.ValidatorStakeAmount)

	regionID, err := testcase.TestNewRegion(validatorID)
	if err != nil {
		return "", false
	}
	return regionID, true
}

/*
TestKycDelegate
@Description: KYC User delegate to validator
@param regionID
@return bool
@return map[string]string Wallet account
*/
func TestKycDelegate(regionID string) (map[string]string, bool) {
	// New user account
	walletAcc, err := tools.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	userPrivKey := walletAcc["PrivKey"]

	if regionID == "" {
		newRegionID, ok := TestNewValidatorRegion(tools.RandNodeID())
		if !ok {
			return walletAcc, false
		}
		regionID = newRegionID
	}

	if err = testcase.TestNewKyc(regionID, userPrivKey); err != nil {
		return walletAcc, false
	}
	if err = testcase.TestTx(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}
	if err = testcase.TestNewDelegate(userPrivKey, config.DelegateAmount); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

/*
TestDelegate
@Description: Not Kyc Delegate to validator
@return bool
@return map[string]string Wallet account
*/
func TestDelegate() (map[string]string, bool) {
	walletAcc, err := tools.GenWalletAcc()
	if err != nil {
		return walletAcc, false
	}
	if err = testcase.TestTx(config.SuperAdminPrivKey, walletAcc["Addr"], config.TxAmount); err != nil {
		return walletAcc, false
	}

	if err = testcase.TestNewDelegate(walletAcc["PrivKey"], config.DelegateAmount); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

/*
TestMixDelegate
@Description: Mix kyc and non-KYC delegates
@param regionID
@return map[string]string Wallet account
@return bool
*/
func TestMixDelegate(regionID string) (map[string]string, bool) {
	// New user account and not kyc delegate
	walletAcc, ok := TestDelegate()
	if !ok {
		return walletAcc, false
	}

	if err := testcase.TestNewKyc(regionID, walletAcc["PrivKey"]); err != nil {
		return walletAcc, false
	}

	if err := testcase.TestNewDelegate(walletAcc["PrivKey"], config.DelegateAmount); err != nil {
		return walletAcc, false
	}

	return walletAcc, true
}

/*
TestKycUnDelegate
@Description: KYC user un-delegate
@param regionID
@return map[string]string
@return bool
*/
func TestKycUnDelegate(regionID string) (map[string]string, bool) {
	walletAcc, ok := TestKycDelegate(regionID)
	if !ok {
		return walletAcc, false
	}

	if err := testcase.TestUnDelegate(walletAcc["PrivKey"], config.DelegateAmount, true); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

/*
TestUnDelegate
@Description: Not KYC user un-delegate
@return map[string]string
@return bool
*/
func TestUnDelegate() (map[string]string, bool) {
	walletAcc, ok := TestDelegate()
	if !ok {
		return walletAcc, false
	}

	if err := testcase.TestUnDelegate(walletAcc["PrivKey"], config.DelegateAmount, false); err != nil {
		return walletAcc, false
	}
	return walletAcc, true
}

func TestMixUnDelegate(regionID string) (map[string]string, bool) {
	walletAcc, ok := TestMixDelegate(regionID)
	if !ok {
		return walletAcc, false
	}

	if err := testcase.TestUnDelegate(walletAcc["PrivKey"], config.DelegateAmount, false); err != nil {
		return walletAcc, false
	}
	if err := testcase.TestUnDelegate(walletAcc["PrivKey"], config.DelegateAmount, true); err != nil {
		return walletAcc, false
	}

	return walletAcc, true
}
