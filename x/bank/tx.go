package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

func (k *Keeper) TransferTx(privKey, toAddr string, amount sdk.Coin) (*txpb.BroadcastTxResponse, error) {

	fromAccAddr, _ := client.GetAccAddress(privKey)
	toAccAddr, err := sdk.AccAddressFromBech32(toAddr)
	if err != nil {
		return nil, err
	}
	zap.S().Info("TransferTx/fromAccAddr: ", fromAccAddr.String())
	zap.S().Info("TransferTx/toAddr: ", toAccAddr.String())

	msg := bankpb.NewMsgSend(fromAccAddr, toAccAddr, sdk.NewCoins(amount))
	if msg.ValidateBasic() != nil {
		return nil, err
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *Keeper) TxToAdmin(privKey string, amount sdk.Coin) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, err := client.GetAccAddress(privKey)
	if err != nil {
		return nil, err
	}
	zap.S().Info("TxToAdmin/fromAccAddr: ", fromAccAddr.String())

	msg := bankpb.NewMsgSendToAdmin(fromAccAddr, sdk.NewCoins(amount))
	if msg.ValidateBasic() != nil {
		return nil, err
	}
	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, 0)
	if err != nil {
		return nil, err
	}
	return res, nil
}
