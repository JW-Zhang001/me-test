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
	return KycArgs{config.SuperAdminPrivKey, "", "${RegionID}", &Dependence{extract}}, nil
}

func TestNewKyc(regionID, userPrivKey string) error {
	testdata, err := NewKycArgs()
	if err != nil {
		return err
	}

	fromPrivKey := testdata.PrivKey

	userAccAddr, err := tools.GetAccAddress(userPrivKey)
	if err != nil {
		return err
	}
	res, err := StakeKeeper.NewKyc(fromPrivKey, userAccAddr.String(), regionID)
	if err != nil {
		zap.S().Errorf("NewKyc error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewKyc TxResponse error %v", res.TxResponse.RawLog)
		return err
	}

	return nil
}
