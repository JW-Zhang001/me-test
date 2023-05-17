package testcase

import (
	"fmt"
	"go.uber.org/zap"
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
	nodeID := "node9"
	tmPubK, err := initialize.GetValidatorPubKey(nodeID)
	if err != nil {
		return ValidatorArgs{}, fmt.Errorf("GetValidatorPubKey error %v", err)
	}

	return ValidatorArgs{config.SuperAdminPrivKey, tmPubK, "100mec", nodeID,
		&Dependence{extract}}, nil
}

func TestNewValidator() (validatorID string, err error) {
	testdata, err := NewValidatorArgs()
	if err != nil {
		return "", err
	}

	privKey := testdata.PrivKey
	tmPubKeyStr := testdata.TmPubKeyStr
	coinStr := testdata.CoinStr
	moniker := testdata.Moniker

	res, err := StakeKeeper.NewValidator(privKey, tmPubKeyStr, coinStr, moniker)
	if err != nil {
		zap.S().Errorf("NewValidator error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
		return "", err
	}
	validatorID, err = StakeKeeper.GetValidatorID(res)
	if err != nil {
		zap.S().Errorf("GetValidatorID error %v", err)
		return "", err
	}
	zap.S().Info("NewValidator/validatorID: ", validatorID)
	return validatorID, nil
}
