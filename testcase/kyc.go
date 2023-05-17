package testcase

import (
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

type KycArgs struct {
	PrivKey  string
	ToAddr   string
	RegionId string
	*Dependence
}

func NewKycArgs() (KycArgs, error) {
	PrivKey := client.GenAccPriKey()

	return KycArgs{config.SuperAdminPrivKey, PrivKey, "", &Dependence{extract}}, nil
}

func TestNewKyc(regionID string) (userPrivKey string, err error) {
	testdata, err := NewKycArgs()
	if err != nil {
		return "", err
	}

	privKey := testdata.PrivKey
	userPrivKey = testdata.ToAddr
	userAccAddr, err := client.GetAccAddress(userPrivKey)
	if err != nil {
		return "", err
	}
	res, err := StakeKeeper.NewKyc(privKey, userAccAddr.String(), regionID)
	if err != nil {
		zap.S().Errorf("NewKyc error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewKyc TxResponse error %v", res.TxResponse.RawLog)
		return "", err
	}
	return userPrivKey, nil
}
