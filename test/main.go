package main

import (
	"context"
	"fmt"
	"me-test/client"
	"me-test/config"
)

func main() {
	//fmt.Println(api.Api, len(api.Api))
	//cases.BlockReward()
	//handle.GetBlockHeight()
	cmClient, err := client.NewCmClient(config.DefaultAPI, config.DefaultRPC)
	if err != nil {
		return
	}
	block, err := cmClient.CliHTTP.Block(context.Background(), nil)
	if err != nil {
		return
	}
	fmt.Println(block)
}
