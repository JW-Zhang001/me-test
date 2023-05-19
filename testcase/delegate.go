package testcase

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/zap"
)

type DelegateArgs struct {
	PrivKey string
	Amount  sdk.Coin
	*Dependence
}

func NewDelegateArgs(amount int64) (DelegateArgs, error) {
	amt := sdk.NewInt64Coin(sdk.DefaultBondDenom, amount)

	return DelegateArgs{"${PrivKey}", amt, &Dependence{extract}}, nil
}

func TestNewDelegate(privKey string, amount int64) error {
	testdata, err := NewDelegateArgs(amount)
	if err != nil {
		return err
	}

	res, err := StakeKeeper.Delegate(privKey, testdata.Amount)
	if err != nil {
		zap.S().Errorf("Delegate error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("Delegate TxResponse error %v", res.TxResponse.RawLog)
		return err
	}
	return nil
}

type UnDelegateArgs struct {
	PrivKey string
	Amount  sdk.Coin
	Kyc     bool
	*Dependence
}

func TestUnDelegate(privKey string, amount int64, kyc bool) error {
	testdata := UnDelegateArgs{privKey, sdk.NewInt64Coin(sdk.DefaultBondDenom, amount),
		kyc, &Dependence{extract}}

	res, err := StakeKeeper.UnDelegate(privKey, testdata.Amount, testdata.Kyc)
	if err != nil {
		zap.S().Errorf("UnDelegate error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("UnDelegate TxResponse error %v", res.TxResponse.RawLog)
		return err
	}
	return nil
}
