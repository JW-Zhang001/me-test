package testcase

import (
	"go.uber.org/zap"
	"me-test/config"
	"me-test/tools"
)

type KycArgs struct {
	PrivKey  string
	ToAddr   string
	RegionID string
	*Dependence
}

func NewKycArgs() (KycArgs, error) {
	PrivKey := tools.GenAccPriKey()

	return KycArgs{config.SuperAdminPrivKey, PrivKey, "${RegionID}", &Dependence{extract}}, nil
}

func TestNewKyc(regionID string) (userAccInfo map[string]string, err error) {
	testdata, err := NewKycArgs()
	if err != nil {
		return userAccInfo, err
	}
	userAccInfo = make(map[string]string)

	privKey := testdata.PrivKey
	userPrivKey := testdata.ToAddr
	userAccAddr, err := tools.GetAccAddress(userPrivKey)
	if err != nil {
		return userAccInfo, err
	}
	res, err := StakeKeeper.NewKyc(privKey, userAccAddr.String(), regionID)
	if err != nil {
		zap.S().Errorf("NewKyc error %v", err)
		return userAccInfo, err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewKyc TxResponse error %v", res.TxResponse.RawLog)
		return userAccInfo, err
	}
	userAccInfo["PrivKey"] = userPrivKey
	userAccInfo["Addr"] = userAccAddr.String()
	return userAccInfo, nil
}
