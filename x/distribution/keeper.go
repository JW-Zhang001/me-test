package distribution

import (
	"context"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	dispb "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/config"
)

type Keeper struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewKeeper() (*Keeper, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &Keeper{Cli: c, Ctx: ctx}, cancel
}

func (k *Keeper) WithdrawRewards(privKey string) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := client.GetAccAddress(privKey)
	zap.S().Info("WithdrawRewards/fromAccAddr: ", fromAccAddr.String())

	msg := dispb.NewMsgWithdrawDelegatorReward(fromAccAddr, sdk.ValAddress{})
	if msg.ValidateBasic() != nil {
		return nil, errors.New("invalid msg")
	}

	res, err := k.Cli.SendBroadcastTx(k.Ctx, privKey, msg, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}
