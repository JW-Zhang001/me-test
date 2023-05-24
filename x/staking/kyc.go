package staking

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

func (k *Keeper) NewKyc(privKey, toAddr, regionId string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := client.GetAccAddress(privKey)
	creator := fromAccAddr.String()
	zap.S().Info("NewKyc/fromAccAddr: ", creator)
	zap.S().Info("NewKyc/toAddr: ", toAddr)

	msg := stakepb.NewMsgNewKyc(creator, toAddr, regionId)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("ValidateBasic error")
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
