package staking

import (
	"fmt"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	"me-test/config"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"go.uber.org/zap"

	"me-test/tools"
)

func (k *Keeper) NewValidator(privKey, tmPubKeyStr, coinStr, moniker string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := tools.GetAccAddress(privKey)
	valAddr := sdk.ValAddress(fromAccAddr)
	zap.S().Info("NewValidator/fromAccAddr: ", fromAccAddr.String())
	zap.S().Info("NewValidator/valAddr: ", valAddr.String())
	zap.S().Info("NewValidator/coinStr: ", coinStr)
	zap.S().Info("NewValidator/moniker: ", moniker)

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

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
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

func (k *Keeper) EditValidator(privKey, operatorAddress, ownerAddress, moniker string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := tools.GetAccAddress(privKey)
	valAddr := sdk.ValAddress(fromAccAddr)

	zap.S().Info("EditValidator/fromAddr: ", fromAccAddr.String())
	zap.S().Info("EditValidator/valAddr: ", valAddr.String())
	zap.S().Info("EditValidator/operatorAddress: ", operatorAddress)
	zap.S().Info("EditValidator/ownerAddress: ", ownerAddress)

	description := stakepb.NewDescription(moniker, "", "", "", "")
	msg := stakepb.NewMsgEditValidator(valAddr, description, nil, ownerAddress, operatorAddress)
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic() error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
