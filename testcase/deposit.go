package testcase

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
)

type DepositArgs struct {
	PrivKey string
	Amount  sdk.Coin
	Term    stakepb.FixedDepositTerm
	*Dependence
}

func NewDepositArgs(amount int64, term stakepb.FixedDepositTerm) (DepositArgs, error) {
	return DepositArgs{"${PrivKey}", sdk.NewInt64Coin(sdk.DefaultBondDenom, amount), term, &Dependence{extract}}, nil
}

func TestDeposit(privKey string, amount int64, term stakepb.FixedDepositTerm) error {
	testdata, err := NewDepositArgs(amount, term)
	if err != nil {
		return err
	}

	res, err := StakeKeeper.Deposit(privKey, testdata.Amount, testdata.Term)
	if err != nil {
		zap.S().Errorf("Deposit error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("Deposit TxResponse error %v", res.TxResponse.RawLog)
		return fmt.Errorf("deposit TxResponse.Code error %v", res.TxResponse.Code)
	}
	return nil
}

type DepositWithdrawArgs struct {
	PrivKey string
	Id      uint64
	*Dependence
}

func TestDepositWithdraw(privKey string, depositID uint64) error {
	testdata := DepositWithdrawArgs{privKey, depositID, &Dependence{extract}}

	res, err := StakeKeeper.DepositWithdraw(testdata.PrivKey, testdata.Id)
	if err != nil {
		zap.S().Errorf("DepositWithdraw error %v", err)
		return err
	}
	if res.TxResponse.Code != 0 {
		zap.S().Errorf("DepositWithdraw TxResponse error %v", res.TxResponse.RawLog)
		return fmt.Errorf("DepositWithdraw TxResponse.Code error %v", res.TxResponse.Code)
	}
	return nil
}
