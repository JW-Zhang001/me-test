package staking

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"

	"me-test/client"
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

	i, err := k.Cil.GetAccountI(k.Ctx, creator)
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
	zap.S().Info("NewKyc res: ", res)
	return res, nil
}
