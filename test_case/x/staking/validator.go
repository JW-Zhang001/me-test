package staking

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmcryptopb "github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
)

func NewValidator() {
	c, _ := client.NewCmClient("")
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()

	privKey := config.SuperAdminPrivKey
	valAdd, _ := client.GetAccAddress(privKey)
	valAddr := sdk.ValAddress(valAdd)

	fromAddr, _ := client.GetAccAddress(privKey)
	zap.S().Info("0", fromAddr.String())

	i, err := c.GetAccountI(ctx, fromAddr.String())

	var tmPubKey tmcryptopb.PubKey
	err = tmjson.Unmarshal([]byte("{\"type\": \"tendermint/PubKeyEd25519\",\"value\": \"LvHZp4ALtrhPxn55Q9nIwZ69zLKeCBY2pmBB26WPHTM=\"}"), &tmPubKey)
	if err != nil {
		return
	}

	pubKey, err := codec.FromTmPubKeyInterface(tmPubKey)
	if err != nil {
		fmt.Println("public key to app public key error: ", err)
	}

	amountStr := "100mec"
	amount, err := sdk.ParseCoinNormalized(amountStr)

	description := stakepb.NewDescription("node5", "", "", "", "")

	rate, err := sdk.NewDecFromStr("0.1")
	if err != nil {
		return
	}
	maxRate, err := sdk.NewDecFromStr("0.2")
	if err != nil {
		return
	}
	maxChangeRate, err := sdk.NewDecFromStr("0.01")
	if err != nil {
		return
	}

	commissionRates := stakepb.NewCommissionRates(rate, maxRate, maxChangeRate)

	minSelfStake, ok := sdk.NewIntFromString("1")
	if !ok {
		return
	}

	msg, err := stakepb.NewMsgCreateValidator(valAddr, pubKey, amount, description, commissionRates, minSelfStake)
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
