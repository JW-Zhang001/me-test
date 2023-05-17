package testcase

import (
	"go.uber.org/zap"
	"me-test/client"
	"me-test/tester"
)

func NewValidatorRegion() (string, error) {
	for _, v := range tester.TestSuites {
		if v.Step == 1 {
			valData := v.Data.(tester.ValidatorArgs)

			privKey := valData.PrivKey
			tmPubKeyStr := valData.TmPubKeyStr
			coinStr := valData.CoinStr
			moniker := valData.Moniker

			res, err := tester.StakeKeeper.NewValidator(privKey, tmPubKeyStr, coinStr, moniker)
			if err != nil {
				zap.S().Errorf("NewValidator error %v", err)
				return "", err
			}
			if res.TxResponse.Code != 0 {
				zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
				return "", err
			}
			validatorID, err := tester.StakeKeeper.GetValidatorID(res)
			if err != nil {
				zap.S().Errorf("GetValidatorID error %v", err)
				return "", err
			}
			zap.S().Info("NewValidator/validatorID: ", validatorID)
			valData.Extract["validatorID"] = validatorID
		} else if v.Step == 2 {
			regionData := v.Data.(tester.RegionArgs)

			privKey := regionData.PrivKey
			regionId := regionData.RegionId
			name := regionData.Name
			validator := regionData.Extract["validatorID"]

			res, err := tester.StakeKeeper.NewRegion(privKey, regionId, name, validator)
			if err != nil {
				zap.S().Errorf("NewRegion error %v", err)
				return "", err
			}
			if res.TxResponse.Code != 0 {
				zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
				return "", err
			}
			return regionId, nil
		}

	}
	return "", nil
}

func NewKyc() (string, error) {
	regionID, err := NewValidatorRegion()
	if err != nil {
		return "", err
	}

	privKey := tester.KycData.PrivKey
	userPrivKey := tester.KycData.ToAddr
	userAccAddr, err := client.GetAccAddress(userPrivKey)
	if err != nil {
		return "", err
	}
	res, err := tester.StakeKeeper.NewKyc(privKey, userAccAddr.String(), regionID)
	if err != nil {
		zap.S().Errorf("NewRegion error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
		return "", err
	}
	return userPrivKey, nil
}
