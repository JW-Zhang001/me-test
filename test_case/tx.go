package test_case

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"

	"me-test/client"
	"me-test/config"
)

func TransferTx() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	privKey := "f7ef778caa413f8fae5c881c1d1c1d8a3ee0365bbee52514b2fd55be252131a2"
	fromAddr, _ := client.GetAccAddress(privKey)
	toAddr, _ := client.GetAccAddress(client.GenAccPriKey())
	zap.S().Info("0", fromAddr.String())
	zap.S().Info("1", toAddr.String())

	i, err := c.GetAccountI(ctx, fromAddr.String())

	coins := sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 1))
	msg := bankpb.NewMsgSend(fromAddr, toAddr, coins)
	if msg.ValidateBasic() != nil {
		return
	}
	pk, _ := client.ConvertsAccPrivKey(privKey)

	tx, err := c.BuildTx(msg, pk, i.GetSequence(), i.GetAccountNumber())
	if err != nil {
		return
	}

	txBytes, err := c.Encoder(tx)
	if err != nil {
		return
	}
	res, err := c.BroadcastTx(txBytes)
	if err != nil {
		return
	}
	zap.S().Info(res)
}
