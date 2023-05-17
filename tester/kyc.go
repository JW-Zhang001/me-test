package tester

import (
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
