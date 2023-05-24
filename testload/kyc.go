package testload

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"me-test/client"
	"me-test/config"
)

func NewKyc(privKey, toAddr, regionId string, sequence uint64) (*txpb.BroadcastTxResponse, error) {
	fromAddr, _ := client.GetAccAddrStr(privKey)
	msg := stakepb.NewMsgNewKyc(fromAddr, toAddr, regionId)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("ValidateBasic error")
	}

	res, err := k.Cli.SendBroadcastCheckTx(k.Ctx, privKey, msg, sequence, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
