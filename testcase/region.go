package testcase

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"me-test/config"
)

type RegionArgs struct {
	PrivKey   string
	RegionId  string
	Name      string
	Validator string
	*Dependence
}

func RandRegionID() string {
	u := uuid.New()
	return u.String()
}

func NewRegionArgs() (RegionArgs, error) {
	regionKey := StakeKeeper.RandRegionKey()
	randomIndex := rand.Intn(len(regionKey))

	return RegionArgs{config.SuperAdminPrivKey, RandRegionID(), regionKey[randomIndex],
		"${Validator}", &Dependence{extract}}, nil
}

func TestNewRegion(validatorID string) (regionID string, err error) {
	testdata, err := NewRegionArgs()
	if err != nil {
		return "", err
	}

	privKey := testdata.PrivKey
	regionID = testdata.RegionId
	name := testdata.Name
	validator := validatorID

	res, err := StakeKeeper.NewRegion(privKey, regionID, name, validator)
	if err != nil {
		zap.S().Errorf("NewRegion error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("NewRegion TxResponse error %v", res.TxResponse.RawLog)
		return "", fmt.Errorf("NewRegion TxResponse.Code error %v", res.TxResponse.Code)
	}
	return regionID, nil
}
