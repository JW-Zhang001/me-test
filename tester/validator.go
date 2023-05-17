package tester

import (
	"fmt"
	"me-test/config"
	"me-test/initialize"
)

type ValidatorArgs struct {
	PrivKey     string
	TmPubKeyStr string
	CoinStr     string
	Moniker     string
	*Dependence
}

func NewValidatorArgs() (ValidatorArgs, error) {
	nodeID := "node8"
	tmPubK, err := initialize.GetValidatorPubKey(nodeID)
	if err != nil {
		return ValidatorArgs{}, fmt.Errorf("GetValidatorPubKey error %v", err)
	}

	return ValidatorArgs{config.SuperAdminPrivKey, tmPubK, "100mec", nodeID,
		&Dependence{extract}}, nil
}
