package staking

import (
	"context"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"

	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func NewKyc() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	privKey := config.SuperAdminPrivKey
	fromAddr, _ := client.GetAccAddress(privKey)
	toAddr, _ := client.GetAccAddress(client.GenAccPriKey())
	zap.S().Info("0", fromAddr.String())
	zap.S().Info("1", toAddr.String())

	i, err := c.GetAccountI(ctx, fromAddr.String())

	creator := fromAddr.String()
	account := toAddr.String()
	regionId := "1"
	nftId := "go-test-nft-1"

	msg := stakepb.NewMsgNewKyc(creator, account, regionId, nftId)
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
