package testsuite

import (
	"go.uber.org/zap"
	"me-test/testcase"
)

func NewValidatorRegion() (string, error) {
	for _, v := range testcase.TestSuites {
		if v.Step == 1 {
			valData := v.Data.(testcase.ValidatorArgs)

			privKey := valData.PrivKey
			tmPubKeyStr := valData.TmPubKeyStr
			coinStr := valData.CoinStr
			moniker := valData.Moniker

			res, err := testcase.StakeKeeper.NewValidator(privKey, tmPubKeyStr, coinStr, moniker)
			if err != nil {
				zap.S().Errorf("NewValidator error %v", err)
				return "", err
			}
			if res.TxResponse.Code != 0 {
				zap.S().Errorf("NewValidator TxResponse error %v", res.TxResponse.RawLog)
				return "", err
			}
			validatorID, err := testcase.StakeKeeper.GetValidatorID(res)
			if err != nil {
				zap.S().Errorf("GetValidatorID error %v", err)
				return "", err
			}
			zap.S().Info("NewValidator/validatorID: ", validatorID)
			valData.Extract["validatorID"] = validatorID
		} else if v.Step == 2 {
			regionData := v.Data.(testcase.RegionArgs)

			privKey := regionData.PrivKey
			regionId := regionData.RegionId
			name := regionData.Name
			validator := regionData.Extract["validatorID"]

			res, err := testcase.StakeKeeper.NewRegion(privKey, regionId, name, validator)
			if err != nil {
				zap.S().Errorf("NewRegion error %v", err)
				return "", err
			}
			if res.TxResponse.Code != 0 {
				zap.S().Errorf("NewRegion TxResponse error %v", res.TxResponse.RawLog)
				return "", err
			}
			return regionId, nil
		}

	}
	return "", nil
}

func TestKycDelegate() bool {
	validatorID, err := testcase.TestNewValidator()
	if err != nil {
		return false
	}

	regionID, err := testcase.TestNewRegion(validatorID)
	if err != nil {
		return false
	}

	userPrivKey, err := testcase.TestNewKyc(regionID)
	if err != nil {
		return false
	}

	if _, err = testcase.TestNewDelegate(userPrivKey); err != nil {
		return false
	}
	return true
}
