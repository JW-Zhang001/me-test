package query

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"me-test/client"
	"me-test/config"
)

type Query struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewBankQuery() (*Query, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &Query{Cli: c, Ctx: ctx}, cancel
}

func (k *Query) Balance(ctx context.Context, addr string) (*bankpb.QueryBalanceResponse, error) {
	req := &bankpb.QueryBalanceRequest{Address: addr, Denom: sdk.BaseMEDenom}

	rpcRes, err := k.Cli.BankClient.Balance(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (k *Query) AllBalances(ctx context.Context, addr string) (*bankpb.QueryAllBalancesResponse, error) {
	req := &bankpb.QueryAllBalancesRequest{Address: addr}

	rpcRes, err := k.Cli.BankClient.AllBalances(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (k *Query) TotalSupply(ctx context.Context) (*bankpb.QueryTotalSupplyResponse, error) {
	req := &bankpb.QueryTotalSupplyRequest{}

	rpcRes, err := k.Cli.BankClient.TotalSupply(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
