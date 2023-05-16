package staking

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/client"
)

func (k *Keeper) Delegate(delPriKey string, amount sdk.Coin) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("Delegate/delPriKey: ", delPriKey)

	delAccAddr, _ := client.GetAccAddress(delPriKey)
	zap.S().Info("Delegate/delAccAddr: ", delAccAddr.String())
	zap.S().Info("Delegate/amount: ", amount.String())

	i, err := k.Cil.GetAccountI(k.Ctx, delAccAddr.String())

	// valAddr and valStr Pass the null value to determine the kyc status of the user based on delAccAddr
	msg := stakepb.NewMsgDelegate(delAccAddr, sdk.ValAddress{}, amount, "")
	if msg.ValidateBasic() != nil {
		return nil, err
	}
	pk, _ := client.ConvertsAccPrivKey(delPriKey)

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
	zap.S().Info("Delegate res: ", res)
	return res, nil
}

func (k *Keeper) Undelegate(delPriKey string, amount sdk.Coin, kyc bool) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("Undelegate/delPriKey: ", delPriKey)

	delAccAddr, _ := client.GetAccAddress(delPriKey)
	zap.S().Info("Undelegate/delAccAddr: ", delAccAddr.String())
	zap.S().Info("Undelegate/amount: ", amount.String())

	i, err := k.Cil.GetAccountI(k.Ctx, delAccAddr.String())

	// valAddr Pass the null value
	msg := stakepb.NewMsgUndelegate(delAccAddr, sdk.ValAddress{}, amount, kyc)
	if msg.ValidateBasic() != nil {
		return nil, err
	}
	pk, _ := client.ConvertsAccPrivKey(delPriKey)

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
	zap.S().Info("Undelegate res: ", res)
	return res, nil
}
