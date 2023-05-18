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

func TestNewDelegate(privKey string, amount int64) (string, error) {
	testdata, err := NewDelegateArgs(amount)
	if err != nil {
		return "", err
	}

	amt := testdata.Amount

	res, err := StakeKeeper.Delegate(privKey, amt)
	if err != nil {
		zap.S().Errorf("Delegate error %v", err)
		return "", err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("Delegate TxResponse error %v", res.TxResponse.RawLog)
		return "", err
	}
	return privKey, nil
}
