package main

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"me-test/client"
	"me-test/config"
	"me-test/tools"
	"time"
)

func main() {

	test2()
}

func test2() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	privKey := "f7ef778caa413f8fae5c881c1d1c1d8a3ee0365bbee52514b2fd55be252131a2"
	fromAddr, _ := tools.GetAccAddr(privKey)
	toAddr, _ := tools.GetAccAddr(tools.GenPriKey())
	fmt.Println("0", fromAddr.String())
	fmt.Println("1", toAddr.String())

	i, err := c.GetAccountI(ctx, fromAddr.String())

	coins := sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 1))
	msg := bankpb.NewMsgSend(fromAddr, toAddr, coins)
	if msg.ValidateBasic() != nil {
		return
	}
	pk, _ := tools.GetPrivKey(privKey)

	tx, err := c.BuildTx(msg, pk, i.GetSequence(), i.GetAccountNumber())
	if err != nil {
		return
	}

	txBytes, err := c.Encoder(tx)
	if err != nil {
		return
	}
	_ = c.BroadcastTx(txBytes)
}
