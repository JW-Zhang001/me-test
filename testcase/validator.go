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

type EditValidatorArgs struct {
	PrivKey         string
	OperatorAddress string
	OwnerAddress    string
	Moniker         string
	*Dependence
}

func NewEditValidatorArgs(ownerAddr, moniker string) (EditValidatorArgs, error) {
	operatorAddr, err := TestNewValidator(moniker, config.ValidatorStakeAmount)
	if err != nil {
		zap.S().Errorf("Init NewEditValidatorArgs error %v", err)
		return EditValidatorArgs{}, err
	}
	return EditValidatorArgs{config.SuperAdminPrivKey, operatorAddr, ownerAddr, moniker,
		&Dependence{extract}}, nil
}

func TestEditValidator(ownerAddr, moniker string) (validatorID string, err error) {
	testdata, err := NewEditValidatorArgs(ownerAddr, moniker)
	if err != nil {
		return "", err
	}
	res, err := StakeKeeper.EditValidator(testdata.PrivKey, testdata.OperatorAddress, testdata.OwnerAddress, testdata.Moniker)
	if err != nil {
		zap.S().Errorf("EditValidator error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("EditValidator TxResponse error %v", res.TxResponse.RawLog)
		return "", fmt.Errorf("EditValidator TxResponse.Code error %v", res.TxResponse.Code)
	}
	// validatorID is operatorAddress
	return testdata.OperatorAddress, nil
}
