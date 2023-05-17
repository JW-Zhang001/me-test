package tester

import (
	"github.com/google/uuid"
	"math/rand"
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
		"${validator}", &Dependence{extract}}, nil
}
