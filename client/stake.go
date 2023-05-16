package client

import (
	"context"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (c *CmClient) Delegation(ctx context.Context, delAddr string) (*stakepb.QueryDelegationResponse, error) {
	req := &stakepb.QueryDelegationRequest{DelegatorAddr: delAddr, ValidatorAddr: ""}
	rpcRes, err := c.StakeClient.Delegation(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
