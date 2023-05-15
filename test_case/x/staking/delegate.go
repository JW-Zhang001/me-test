package staking

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

func Delegate() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	delPriKey := client.GenAccPriKey()
	zap.S().Info("delPriKey: ", delPriKey)

	delAccAddr, _ := client.GetAccAddress(delPriKey)
	sendToUser(delAccAddr)
	zap.S().Info("delAccAddr: ", delAccAddr.String())

	amount := sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000)

	i, err := c.GetAccountI(ctx, delAccAddr.String())

	// valAddr and valStr Pass the null value to determine the kyc status of the user based on delAccAddr
	msg := stakepb.NewMsgDelegate(delAccAddr, sdk.ValAddress{}, amount, "")
	if msg.ValidateBasic() != nil {
		return
	}
	pk, _ := client.ConvertsAccPrivKey(delPriKey)

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
	fmt.Println(res)
}

func sendToUser(toAddr sdk.AccAddress) {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	privKey := "62b8fc967f3081a2ae95d69a183e7a6ac710b7e03b9e1195aa614ca4f3408092"
	fromAddr, _ := client.GetAccAddress(privKey)
	//toAddr, _ := client.GetAccAddress(client.GenAccPriKey())
	zap.S().Info("0", fromAddr.String())
	//zap.S().Info("1", toAddr.String())

	i, err := c.GetAccountI(ctx, fromAddr.String())

	coins := sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 10000000))
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

func initStakeClient() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	delAddr := "cosmos1arr3ekq844te4mp3l2lt4lddcrvntgwncr0jdz"
	valAddr := ""

	req := &stakepb.QueryDelegationRequest{DelegatorAddr: delAddr, ValidatorAddr: valAddr}

	rpcRes, err := c.StakeClient.Delegation(ctx, req)
	if err != nil {
		return
	}
	fmt.Println(rpcRes)
}

func Undelegate() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	delPriKey := "d3dbd304664049602931abd7b1931dc33dd38087d7f051f18d02682375006905"
	delAccAddr, _ := client.GetAccAddress(delPriKey)

	zap.S().Info("delAccAddr: ", delAccAddr.String())

	amount := sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000)

	i, err := c.GetAccountI(ctx, delAccAddr.String())

	// valAddr Pass the null value
	msg := stakepb.NewMsgUndelegate(delAccAddr, sdk.ValAddress{}, amount, false)
	if msg.ValidateBasic() != nil {
		return
	}
	pk, _ := client.ConvertsAccPrivKey(delPriKey)

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
	fmt.Println(res)
}
