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

func (k *Keeper) Delegate(delPriKey string, amount sdk.Coin) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("Delegate/delPriKey: ", delPriKey)

	delAccAddr, _ := client.GetAccAddress(delPriKey)
	zap.S().Info("Delegate/delAddr: ", delAccAddr.String())
	zap.S().Info("Delegate/amount: ", amount.String())

	// valAddr and valStr Pass the null value to determine the kyc status of the user based on delAccAddr
	msg := stakepb.NewMsgDelegate(delAccAddr, sdk.ValAddress{}, amount, "")
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, delPriKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *Keeper) UnDelegate(delPriKey string, amount sdk.Coin, kyc bool) (*txpb.BroadcastTxResponse, error) {
	zap.S().Info("Undelegate/delPriKey: ", delPriKey)

	delAccAddr, _ := client.GetAccAddress(delPriKey)
	zap.S().Info("Undelegate/delAccAddr: ", delAccAddr.String())
	zap.S().Info("Undelegate/amount: ", amount.String())

	// valAddr Pass the null value
	msg := stakepb.NewMsgUndelegate(delAccAddr, sdk.ValAddress{}, amount, kyc)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, delPriKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
