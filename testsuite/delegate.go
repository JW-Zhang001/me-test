package testsuite

import (
	"me-test/config"
	"me-test/testcase"
)

/*
TestKycDelegate
@Description: KYC User delegate to validator
@return bool
*/
func TestKycDelegate() bool {
	nodeID := "node10"
	validatorID, err := testcase.TestNewValidator(nodeID, config.ValidatorStakeAmount)
	if err != nil {
		return false
	}

	regionID, err := testcase.TestNewRegion(validatorID)
	if err != nil {
		return false
	}

	userAccInfo, err := testcase.TestNewKyc(regionID)
	if err != nil {
		return false
	}
	userPrivKey := userAccInfo["PrivKey"]

	if err := testcase.TestTx("", userAccInfo["Addr"], config.TxAmount); err != nil {
		return false
	}

	if _, err = testcase.TestNewDelegate(userPrivKey, config.DelegateAmount); err != nil {
		return false
	}
	return true
}
