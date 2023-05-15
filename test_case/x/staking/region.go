package staking

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"

	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func NewRegion() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	privKey := config.SuperAdminPrivKey

	fromAddr, _ := client.GetAccAddress(privKey)

	i, err := c.GetAccountI(ctx, fromAddr.String())

	creator := fromAddr.String()
	regionId := "regionId-01"
	name := "regionName-01"
	validator := "cosmosvaloper1auhr8a3sq3u0mvx94pszqmvy8azwvn9y2kcyvt"
	nftClassId := "nft_class_id-01"

	msg := stakepb.NewMsgNewRegion(creator, regionId, name, validator, nftClassId)
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
	fmt.Println(res)
	zap.S().Info(res)
}
