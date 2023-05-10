package main

import (
	"context"
	"fmt"
	"me-test/client"
)

func main() {

	c, _ := client.NewCmClient("")
	ctx := context.Background()

	//// 查询节点状态
	//nodeInfo, err := c.TmClient.GetNodeInfo(ctx, &tmservice.GetNodeInfoRequest{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(nodeInfo)
	//
	//sync, err := c.TmClient.GetSyncing(ctx, &tmservice.GetSyncingRequest{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(sync)

	balance, err := c.Balance(ctx, "cosmos1gmpxkchcdgfq995zye5efwzfw86zfa4vt4ke8g")
	if err != nil {
		return
	}
	fmt.Println(balance)
}

//func main1() {
//
//	c, _ := client.NewCmClient("")
//	ctx := context.Background()
//
//	// 创建发送方的账户信息
//	senderAddr := sdk.AccAddress("cosmos1ldzu259gdeta5eypglecz0rgpynj3vepqruqne")
//	senderPrivKey := "3ed28b8a51fbd6e4644f66d869e95e908aa949855eea7ff3ffa9fa8ff787a9b9"
//	acc := authTypes.NewBaseAccountWithAddress(senderAddr)
//
//	// 创建转账交易
//	recipientAddr := sdk.AccAddress("cosmos1ytcr6f0guyx3cg45wx8fxuh2etxzf2h7hclmza")
//	amountStr := "100"
//	amount, err := sdk.ParseCoinsNormalized(amountStr + config.DefaultDenom)
//	if err != nil {
//		panic(err)
//	}
//	msg := bankTypes.NewMsgSend(senderAddr, recipientAddr, amount)
//
//}
