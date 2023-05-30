package query

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"

	"me-test/client"
	"me-test/config"
)

type BankQueryClient struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewBankQuery() (*BankQueryClient, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)
	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &BankQueryClient{Cli: c, Ctx: ctx}, cancel
}

func (q *BankQueryClient) Balance(addr string) (*bankpb.QueryBalanceResponse, error) {
	req := &bankpb.QueryBalanceRequest{Address: addr, Denom: sdk.BaseMEDenom}
	rpcRes, err := q.Cli.BankClient.Balance(BankQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *BankQueryClient) AllBalances(addr string) (*bankpb.QueryAllBalancesResponse, error) {
	req := &bankpb.QueryAllBalancesRequest{Address: addr}
	rpcRes, err := q.Cli.BankClient.AllBalances(BankQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *BankQueryClient) TotalSupply() (*bankpb.QueryTotalSupplyResponse, error) {
	req := &bankpb.QueryTotalSupplyRequest{}
	rpcRes, err := q.Cli.BankClient.TotalSupply(BankQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
