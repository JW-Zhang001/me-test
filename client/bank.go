package client

import (
	"context"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"

	"me-test/config"
)

func (c *CmClient) Balance(ctx context.Context, addr string) (*bankpb.QueryBalanceResponse, error) {
	req := &bankpb.QueryBalanceRequest{Address: addr, Denom: config.DefaultDenom}

	rpcRes, err := c.BankClient.Balance(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (c *CmClient) AllBalances(ctx context.Context, addr string) (*bankpb.QueryAllBalancesResponse, error) {
	req := &bankpb.QueryAllBalancesRequest{Address: addr}

	rpcRes, err := c.BankClient.AllBalances(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (c *CmClient) TotalSupply(ctx context.Context) (*bankpb.QueryTotalSupplyResponse, error) {
	req := &bankpb.QueryTotalSupplyRequest{}

	rpcRes, err := c.BankClient.TotalSupply(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
