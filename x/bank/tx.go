package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"
	"me-test/client"
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

	i, err := k.Cil.GetAccountI(k.Ctx, fromAccAddr.String())
	if err != nil {
		return nil, err
	}

	pk, _ := client.ConvertsAccPrivKey(privKey)

	tx, err := k.Cil.BuildTx(msg, pk, i.GetSequence(), i.GetAccountNumber())
	if err != nil {
		return nil, err
	}

	txBytes, err := k.Cil.Encoder(tx)
	if err != nil {
		return nil, err
	}
	res, err := k.Cil.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}
	zap.S().Info("TransferTx res: ", res)
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

	i, err := k.Cil.GetAccountI(k.Ctx, fromAccAddr.String())
	if err != nil {
		return nil, err
	}

	pk, _ := client.ConvertsAccPrivKey(privKey)

	tx, err := k.Cil.BuildTx(msg, pk, i.GetSequence(), i.GetAccountNumber())
	if err != nil {
		return nil, err
	}

	txBytes, err := k.Cil.Encoder(tx)
	if err != nil {
		return nil, err
	}
	res, err := k.Cil.BroadcastTx(txBytes)
	if err != nil {
		return nil, err
	}
	zap.S().Info("TransferTx res: ", res)
	return res, nil
}
