package testcase

import (
	"fmt"
	"me-test/testcase/data"
)

func NewValidatorRegion() {
	for i, v := range data.TestSuites {
		if v.Step == 1 {
			fmt.Println("------------", i)

			v1 := v.Data.(data.ValidatorArgs)

			privKey := v1.PrivKey
			tmPubKeyStr := v1.TmPubKeyStr
			coinStr := v1.CoinStr
			moniker := v1.Moniker

			res, err := data.StakeKeeper.NewValidator(privKey, tmPubKeyStr, coinStr, moniker)
			fmt.Println(res, err)
		}
	}
}
