package testsuite

import (
	"me-test/check"
	"me-test/config"
	"me-test/testcase"
)

func TestNewValidator(nodeID string) (string, bool) {
	fn := check.CheckerNewValidator(testcase.TestNewValidator)
	validatorID, err := fn(nodeID, config.ValidatorStakeAmount)
	if err != nil {
		return "", false
	}
	return validatorID, true
}

func TestEditValidator(ownerAddr, nodeID string) (string, bool) {
	validatorID, err := testcase.TestEditValidator(ownerAddr, nodeID)
	if err != nil {
		return "", false
	}
	return validatorID, true
}
