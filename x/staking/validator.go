package staking

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"go.uber.org/zap"

	"me-test/client"
)

func (k *Keeper) NewValidator(privKey, tmPubKeyStr, coinStr, moniker string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := client.GetAccAddress(privKey)
	valAddr := sdk.ValAddress(fromAccAddr)
	zap.S().Info("NewValidator/fromAccAddr: ", fromAccAddr.String())
	zap.S().Info("NewValidator/valAddr: ", valAddr.String())

	var tmPubKey tmcrypto.PubKey
	err := tmjson.Unmarshal([]byte(tmPubKeyStr), &tmPubKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := codec.FromTmPubKeyInterface(tmPubKey)
	if err != nil {
		return nil, fmt.Errorf("public key to app public key error: %v", err)
	}
	coinAmt, err := sdk.ParseCoinNormalized(coinStr)
	if err != nil {
		return nil, err
	}

	description := stakepb.NewDescription(moniker, "", "", "", "")

	rate, _ := sdk.NewDecFromStr("0.1")
	maxRate, _ := sdk.NewDecFromStr("0.2")
	maxChangeRate, _ := sdk.NewDecFromStr("0.01")
	commissionRates := stakepb.NewCommissionRates(rate, maxRate, maxChangeRate)

	minSelfStake, ok := sdk.NewIntFromString("1")
	if !ok {
		return nil, fmt.Errorf("NewIntFromString error")
	}

	msg, err := stakepb.NewMsgCreateValidator(valAddr, pubKey, coinAmt, description, commissionRates, minSelfStake)
	if msg.ValidateBasic() != nil {
		return nil, err
	}
	i, err := k.Cil.GetAccountI(k.Ctx, fromAccAddr.String())
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
	zap.S().Info("NewValidator res: ", res)
	return res, nil
}

func (k *Keeper) GetValidatorID(res *txpb.BroadcastTxResponse) (validatorID string, err error) {
	events := res.TxResponse.Logs[0].Events
	for _, v := range events {
		if v.Type == "create_validator" {
			for _, vv := range v.Attributes {
				if vv.Key == "validator" {
					zap.S().Info("vv.Value: ", vv.Value)
					validatorID = vv.Value
				}
			}
		}
	}
	if validatorID == "" {
		return "", fmt.Errorf("validatorID is empty")
	}
	return validatorID, nil
}
