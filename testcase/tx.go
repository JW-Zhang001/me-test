package testcase

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/zap"
)

type TxArgs struct {
	privKey string
	toAddr  string
	amount  sdk.Coin
	*Dependence
}

func NewTxArgs(privKey, toAddr string, amount int64) (TxArgs, error) {
	return TxArgs{privKey, toAddr, sdk.NewInt64Coin(sdk.BaseMEDenom, amount), &Dependence{extract}}, nil
}

func TestTx(privKey, toAddr string, amount int64) error {
	testdata, err := NewTxArgs(privKey, toAddr, amount)
	if err != nil {
		return err
	}
	privKey = testdata.privKey
	toAddr = testdata.toAddr

	res, err := BankKeeper.TransferTx(privKey, toAddr, testdata.amount)
	if err != nil {
		zap.S().Errorf("TestTx error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("TestTx TxResponse error %v", res.TxResponse.RawLog)
		return fmt.Errorf("TestTx TxResponse.Code error %v", res.TxResponse.Code)
	}
	return nil
}
