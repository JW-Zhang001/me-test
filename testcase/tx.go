package testcase

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"go.uber.org/zap"
	"me-test/config"
	"me-test/tools"
)

type TxArgs struct {
	privKey string
	toAddr  string
	amount  sdk.Coin
	*Dependence
}

func NewTxArgs(privKey, toAddr string, amount int64) (TxArgs, error) {
	if privKey == "" {
		privKey = config.SuperAdminPrivKey
	}
	if toAddr == "" {
		toPriKey := tools.GenAccPriKey()
		toAccAddr, err := tools.GetAccAddress(toPriKey)
		if err != nil {
			return TxArgs{}, err
		}
		toAddr = toAccAddr.String()
	}
	amt := sdk.NewInt64Coin(sdk.BaseMEDenom, amount)

	return TxArgs{privKey, toAddr, amt, &Dependence{extract}}, nil
}

func TestTx(privKey, toAddr string, amount int64) error {
	testdata, err := NewTxArgs(privKey, toAddr, amount)
	if err != nil {
		return err
	}
	privKey = testdata.privKey
	toAddr = testdata.toAddr
	amt := testdata.amount

	res, err := BankKeeper.TransferTx(privKey, toAddr, amt)
	if err != nil {
		zap.S().Errorf("TestTx error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("TestTx TxResponse error %v", res.TxResponse.RawLog)
		return err
	}
	return nil
}
