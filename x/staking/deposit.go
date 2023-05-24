package staking

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

func (k *Keeper) Deposit(PriKey string, amount sdk.Coin, term stakepb.FixedDepositTerm) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("Deposit/PriKey: ", PriKey)

	userAccAddr, _ := client.GetAccAddress(PriKey)
	zap.S().Info("Deposit/Addr: ", userAccAddr.String())
	zap.S().Info("Deposit/Amount: ", amount.String())
	zap.S().Info("Deposit/Term: ", term.String())

	msg := stakepb.NewMsgDoFixedDeposit(userAccAddr.String(), amount, term)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, PriKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *Keeper) DepositWithdraw(PriKey string, DepositID uint64) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("DepositWithdraw/PriKey: ", PriKey)

	userAccAddr, _ := client.GetAccAddress(PriKey)
	zap.S().Info("DepositWithdraw/Addr: ", userAccAddr.String())
	zap.S().Info("DepositWithdraw/DepositID: ", DepositID)

	msg := stakepb.NewMsgDoFixedWithdraw(userAccAddr.String(), DepositID)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, PriKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
