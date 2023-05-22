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

func NewValidatorArgs(nodeID, coinStr string) (ValidatorArgs, error) {
	tmPubK, err := initialize.GetValidatorPubKey(nodeID)
	if err != nil {
		return ValidatorArgs{}, fmt.Errorf("GetValidatorPubKey error %v", err)
	}

	return ValidatorArgs{config.SuperAdminPrivKey, tmPubK, coinStr, nodeID,
		&Dependence{extract}}, nil
}

func TestNewValidator(nodeID, coinStr string) (validatorID string, err error) {
	testdata, err := NewValidatorArgs(nodeID, coinStr)
	if err != nil {
		return "", err
	}

	privKey := testdata.PrivKey
	tmPubKeyStr := testdata.TmPubKeyStr
	coin := testdata.CoinStr
	moniker := testdata.Moniker

	res, err := StakeKeeper.NewValidator(privKey, tmPubKeyStr, coin, moniker)
	if err != nil {
		zap.S().Errorf("NewValidator error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
		return "", fmt.Errorf("NewValidator TxResponse.Code error %v", res.TxResponse.Code)
	}
	validatorID, err = StakeKeeper.GetValidatorID(res)
	if err != nil {
		zap.S().Errorf("GetValidatorID error %v", err)
		return "", err
	}
	zap.S().Info("NewValidator/validatorID: ", validatorID)
	return validatorID, nil
}
